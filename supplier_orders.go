package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// SupplierOrdersService handles communication with the Inbound Order
// (Supplier Order) endpoints of the Linker Cloud API.
type SupplierOrdersService struct {
	client *Client
}

// ListSupplierOrdersOptions specifies the optional parameters for
// [SupplierOrdersService.List].
type ListSupplierOrdersOptions struct {
	// SortDir is the sort direction: "ASC" or "DESC" (default: "DESC").
	SortDir string
	// SortCol is the column to sort by.
	SortCol string
	// Offset is the zero-based pagination offset.
	Offset int
	// Limit is the maximum number of records to return (default: 100).
	Limit int
	// Filters are arbitrary key/value pairs passed as filters[key]=value.
	Filters map[string]string
}

// List returns a paginated collection of supplier (inbound) orders.
func (s *SupplierOrdersService) List(ctx context.Context, opts ListSupplierOrdersOptions) (*Page[models.SupplierOrder], error) {
	params := listParams(opts.SortDir, opts.SortCol, opts.Offset, opts.Limit, opts.Filters)
	var result Page[models.SupplierOrder]
	if err := s.client.doRequest(ctx, http.MethodGet, "/public-api/v1/supplierorders", params, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns the supplier order identified by id.
func (s *SupplierOrdersService) Get(ctx context.Context, id string) (*models.SupplierOrder, error) {
	var result models.SupplierOrder
	path := fmt.Sprintf("/public-api/v1/supplierorders/%s", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new inbound (supplier) order.
func (s *SupplierOrdersService) Create(ctx context.Context, req models.SupplierOrderType) (*models.SupplierOrderType, error) {
	var result models.SupplierOrderType
	if err := s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/supplierorders", nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update replaces an existing supplier order identified by id.
func (s *SupplierOrdersService) Update(ctx context.Context, id string, req models.SupplierOrderType) (*models.SupplierOrderType, error) {
	var result models.SupplierOrderType
	path := fmt.Sprintf("/public-api/v1/supplierorders/%s", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodPut, path, nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Confirm creates a new inbound confirmation document for the given supplier
// order.
func (s *SupplierOrdersService) Confirm(ctx context.Context, req models.SupplierOrderType) (*models.SupplierOrderType, error) {
	var result models.SupplierOrderType
	if err := s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/supplierorders/confirms", nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
