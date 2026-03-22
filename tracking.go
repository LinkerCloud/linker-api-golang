package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// TrackingService handles communication with the Tracking Number endpoints of
// the Linker Cloud API.
type TrackingService struct {
	client *Client
}

// SetTrackingNumber sets a single tracking number for the order identified by
// orderID. The order is automatically transitioned to closed status.
func (s *TrackingService) SetTrackingNumber(ctx context.Context, orderID string, req models.TrackingNumberRequest) (*models.Order2, error) {
	var result models.Order2
	path := fmt.Sprintf("/public-api/v1/orders/%s/trackingnumber", url.PathEscape(orderID))
	if err := s.client.doRequest(ctx, http.MethodPut, path, nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetTrackingNumbers sets multiple tracking numbers for the order identified
// by orderID. The order is automatically transitioned to closed status.
func (s *TrackingService) SetTrackingNumbers(ctx context.Context, orderID string, req models.TrackingNumbersRequest) (*models.Order2, error) {
	var result models.Order2
	path := fmt.Sprintf("/public-api/v1/orders/%s/trackingnumbers", url.PathEscape(orderID))
	if err := s.client.doRequest(ctx, http.MethodPut, path, nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
