package linkercloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// LabelFormat specifies the format for shipment labels.
type LabelFormat string

const (
	LabelFormatPNG LabelFormat = "PNG"
	LabelFormatPDF LabelFormat = "PDF"
	LabelFormatZPL LabelFormat = "ZPL"
)

// GetLabelsOptions specifies the parameters for [ShipmentsService.GetLabels].
type GetLabelsOptions struct {
	// OrderID is the order's MongoDB ID (required).
	OrderID string
	// PackageID is the package ID (required).
	PackageID string
	// Format is the label format (required): LabelFormatPNG, LabelFormatPDF, or LabelFormatZPL.
	Format LabelFormat
	// ParcelID optionally restricts to a single parcel within the package.
	ParcelID string
}

// ShipmentsService handles communication with the Shipment/Delivery endpoints
// of the Linker Cloud API.
type ShipmentsService struct {
	client *Client
}

// Create creates shipments for the orders specified in req, returning the
// created delivery IDs.
func (s *ShipmentsService) Create(ctx context.Context, req models.CreateShipmentsRequest) (*models.CreateShipmentsResponse, error) {
	var result models.CreateShipmentsResponse
	if err := s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/deliveries", nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePackages creates shipments for a selected order by order number,
// returning the created parcels.
func (s *ShipmentsService) CreatePackages(ctx context.Context, req models.ShipmentType) ([]models.ParcelType, error) {
	var result []models.ParcelType
	if err := s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/deliveries/packages", nil, req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateStatus updates the delivery status of a single package.
func (s *ShipmentsService) UpdateStatus(ctx context.Context, req models.DeliveryStatus) error {
	return s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/deliveries/statuses", nil, req, nil)
}

// CancelPackages cancels packages for the order identified by orderID.
func (s *ShipmentsService) CancelPackages(ctx context.Context, orderID string, req models.CancelPackagesRequest) error {
	path := fmt.Sprintf("/public-api/v1/deliveries/%s", url.PathEscape(orderID))
	return s.client.doRequest(ctx, http.MethodPatch, path, nil, req, nil)
}

// GetLabels returns decoded shipment labels for the specified package.
// Each entry in the returned slice is the raw label bytes (PNG, PDF, or ZPL
// depending on opts.Format).
func (s *ShipmentsService) GetLabels(ctx context.Context, opts GetLabelsOptions) ([][]byte, error) {
	// The label endpoint has no resource prefix — the API route is literally
	// /{orderId}/{packageId}/{format}/{parcelId} under /public-api/v1.
	// The server returns a JSON array of base64-encoded strings.
	path := fmt.Sprintf("/public-api/v1/%s/%s/%s/%s",
		url.PathEscape(opts.OrderID), url.PathEscape(opts.PackageID), url.PathEscape(string(opts.Format)), url.PathEscape(opts.ParcelID))

	var encoded []string
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &encoded); err != nil {
		return nil, err
	}

	labels := make([][]byte, len(encoded))
	for i, s := range encoded {
		b, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return nil, fmt.Errorf("linkercloud: decode label %d: %w", i, err)
		}
		labels[i] = b
	}
	return labels, nil
}
