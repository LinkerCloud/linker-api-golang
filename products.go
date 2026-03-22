package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// ProductsService handles communication with the Product endpoints of the
// Linker Cloud API.
type ProductsService struct {
	client *Client
}

// ListProductsOptions specifies the optional parameters for [ProductsService.List].
type ListProductsOptions struct {
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

// List returns a paginated collection of products.
func (s *ProductsService) List(ctx context.Context, opts ListProductsOptions) (*Page[models.Product], error) {
	params := listParams(opts.SortDir, opts.SortCol, opts.Offset, opts.Limit, opts.Filters)
	var result Page[models.Product]
	if err := s.client.doRequest(ctx, http.MethodGet, "/public-api/v1/products", params, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new product.
func (s *ProductsService) Create(ctx context.Context, req models.ProductType) (*models.Product, error) {
	var result models.Product
	if err := s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/products", nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update replaces an existing product identified by id.
func (s *ProductsService) Update(ctx context.Context, id string, req models.ProductType) (*models.Product, error) {
	var result models.Product
	path := fmt.Sprintf("/public-api/v1/products/%s", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodPut, path, nil, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
