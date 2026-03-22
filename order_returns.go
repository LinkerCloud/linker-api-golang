package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// OrderReturnsService handles communication with the OrderReturns endpoints.
type OrderReturnsService struct {
	client *Client
}

// ListOrderReturnsOptions specifies the optional parameters for
// [OrderReturnsService.List].
type ListOrderReturnsOptions struct {
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

// List returns a paginated collection of order returns.
func (s *OrderReturnsService) List(ctx context.Context, opts ListOrderReturnsOptions) (*Page[models.OrderReturn], error) {
	params := listParams(opts.SortDir, opts.SortCol, opts.Offset, opts.Limit, opts.Filters)
	var result Page[models.OrderReturn]
	if err := s.client.doRequest(ctx, http.MethodGet, "/public-api/v1/orderreturns", params, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns the order return identified by id.
func (s *OrderReturnsService) Get(ctx context.Context, id string) (*models.OrderReturn, error) {
	var result models.OrderReturn
	path := fmt.Sprintf("/public-api/v1/orderreturns/%s", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new order return.
func (s *OrderReturnsService) Create(ctx context.Context, req models.OrderReturnType) (*models.OrderReturn, error) {
	var result models.OrderReturn
	if err := s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/orderreturns", nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Accept confirms an order return identified by id.
func (s *OrderReturnsService) Accept(ctx context.Context, id string) error {
	path := fmt.Sprintf("/public-api/v1/orderreturns/%s/accept", url.PathEscape(id))
	return s.client.doRequest(ctx, http.MethodPost, path, nil, nil, nil)
}
