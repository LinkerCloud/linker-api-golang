package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/linkercloud/linker-api-golang/models"
)

// OrdersService handles communication with the Order endpoints of the
// Linker Cloud API.
type OrdersService struct {
	client *Client
}

// ListOrdersOptions specifies the optional parameters for [OrdersService.List].
// Filters uses the API's filters[key]=value syntax and are appended to the
// query string verbatim.
type ListOrdersOptions struct {
	// SortDir is the sort direction: "ASC" or "DESC" (default: "DESC").
	SortDir string
	// SortCol is the column to sort by (default: "order_date").
	SortCol string
	// Offset is the zero-based pagination offset (default: 0).
	Offset int
	// Limit is the maximum number of records to return (default: 100).
	Limit int
	// Filters are arbitrary key/value pairs passed as filters[key]=value.
	// Example: {"order_date": "15.11.2021", "order_status": "Y"}
	Filters map[string]string
}

// List returns a paginated collection of orders.
// Use opts to filter, sort, and paginate results.
func (s *OrdersService) List(ctx context.Context, opts ListOrdersOptions) (*Page[models.Order], error) {
	params := listParams(opts.SortDir, opts.SortCol, opts.Offset, opts.Limit, opts.Filters)
	var result Page[models.Order]
	if err := s.client.doRequest(ctx, http.MethodGet, "/public-api/v1/orders", params, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns the order identified by id.
func (s *OrdersService) Get(ctx context.Context, id string) (*models.Order, error) {
	var result models.Order
	path := fmt.Sprintf("/public-api/v1/orders/%s", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new order and returns the created resource.
func (s *OrdersService) Create(ctx context.Context, req models.OrderType) (*models.Order2, error) {
	var result models.Order2
	if err := s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/orders", nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update replaces an existing order (full update).
func (s *OrdersService) Update(ctx context.Context, id string, req models.OrderType) (*models.Order2, error) {
	var result models.Order2
	path := fmt.Sprintf("/public-api/v1/orders/%s", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodPut, path, nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Patch partially updates an existing order.
func (s *OrdersService) Patch(ctx context.Context, id string, req models.OrderType) (*models.Order2, error) {
	var result models.Order2
	path := fmt.Sprintf("/public-api/v1/orders/%s", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodPatch, path, nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Cancel cancels the order identified by id.
func (s *OrdersService) Cancel(ctx context.Context, id string) error {
	path := fmt.Sprintf("/public-api/v1/orders/%s/cancel", url.PathEscape(id))
	return s.client.doRequest(ctx, http.MethodPut, path, nil, nil, nil)
}

// GetOriginalItems returns the original items for the order identified by id.
func (s *OrdersService) GetOriginalItems(ctx context.Context, id string) ([]models.OrderItem, error) {
	var result []models.OrderItem
	path := fmt.Sprintf("/public-api/v1/orders/%s/original/items", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ConfirmPicking confirms the picking of an order.
func (s *OrdersService) ConfirmPicking(ctx context.Context, req models.PickingConfirmation) error {
	return s.client.doRequest(ctx, http.MethodPut, "/public-api/v1/order/pickingconfirmation", nil, req, nil)
}

// MarkPicked marks an order as picked.
func (s *OrdersService) MarkPicked(ctx context.Context, req models.MarkPickedRequest) (*models.Order2, error) {
	var result models.Order2
	if err := s.client.doRequest(ctx, http.MethodPut, "/public-api/v1/order/picked", nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePaymentStatus updates the payment status for one or more orders.
func (s *OrdersService) UpdatePaymentStatus(ctx context.Context, req models.PaymentRequest) error {
	return s.client.doRequest(ctx, http.MethodPut, "/public-api/v1/payment-status", nil, req, nil)
}

// listParams builds a [url.Values] from common list-endpoint parameters.
func listParams(sortDir, sortCol string, offset, limit int, filters map[string]string) url.Values {
	v := make(url.Values)
	if sortDir != "" {
		v.Set("sortDir", sortDir)
	}
	if sortCol != "" {
		v.Set("sortCol", sortCol)
	}
	if offset > 0 {
		v.Set("offset", strconv.Itoa(offset))
	}
	if limit > 0 {
		v.Set("limit", strconv.Itoa(limit))
	}
	for k, val := range filters {
		v.Set("filters["+k+"]", val)
	}
	if len(v) == 0 {
		return nil
	}
	return v
}
