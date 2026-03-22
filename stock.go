package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/linkercloud/linker-api-golang/models"
)

// StockService handles communication with the Stock endpoints of the Linker
// Cloud API.
type StockService struct {
	client *Client
}

// ListStocksOptions specifies the optional parameters for [StockService.List].
type ListStocksOptions struct {
	// Offset is the zero-based pagination offset.
	Offset int
	// Limit is the maximum number of records to return.
	Limit int
	// WmsID filters results to the given WMS ID.
	WmsID string
	// HideIgnoredInWMS excludes products that are ignored in WMS when true.
	HideIgnoredInWMS bool
	// Filters are arbitrary key/value pairs passed as filters[key]=value.
	Filters map[string]string
}

// List returns the current stock levels.
func (s *StockService) List(ctx context.Context, opts ListStocksOptions) (*models.Stock, error) {
	params := make(url.Values)
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.WmsID != "" {
		params.Set("wmsId", opts.WmsID)
	}
	if opts.HideIgnoredInWMS {
		params.Set("hideIgnoredInWMS", "true")
	}
	for k, v := range opts.Filters {
		params.Set("filters["+k+"]", v)
	}
	if len(params) == 0 {
		params = nil
	}

	var result models.Stock
	if err := s.client.doRequest(ctx, http.MethodGet, "/public-api/v1/stocks", params, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GenerateWMSStockReportOptions specifies the optional parameters for
// [StockService.GenerateReport].
type GenerateWMSStockReportOptions struct {
	// WmsID restricts the report to a specific WMS ID.
	WmsID string
	// SKU restricts the report to a specific product SKU.
	SKU string
	// Group restricts the report to a specific product group.
	Group string
}

// GenerateReport triggers asynchronous generation of a WMS stock report and
// returns the UUID to be used with [StockService.GetReport].
func (s *StockService) GenerateReport(ctx context.Context, opts GenerateWMSStockReportOptions) (string, error) {
	params := make(url.Values)
	if opts.WmsID != "" {
		params.Set("wmsId", opts.WmsID)
	}
	if opts.SKU != "" {
		params.Set("sku", opts.SKU)
	}
	if opts.Group != "" {
		params.Set("group", opts.Group)
	}
	if len(params) == 0 {
		params = nil
	}
	var resp struct {
		UUID string `json:"uuid"`
	}
	if err := s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/generate-wms-stock-report", params, nil, &resp); err != nil {
		return "", err
	}
	return resp.UUID, nil
}

// GetReportOptions specifies the optional parameters for [StockService.GetReport].
type GetReportOptions struct {
	// Offset is the zero-based pagination offset.
	Offset int
	// Limit is the maximum number of records to return.
	Limit int
}

// GetReport retrieves the WMS stock report identified by uuid.
// The server returns 202 Accepted while the report is still being generated;
// in that case err will be a *[APIError] with StatusCode 202.
func (s *StockService) GetReport(ctx context.Context, uuid string, opts GetReportOptions) (*models.Stock, error) {
	params := make(url.Values)
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}
	if len(params) == 0 {
		params = nil
	}

	var result models.Stock
	path := fmt.Sprintf("/public-api/v1/wms-stock-report/%s", url.PathEscape(uuid))
	if err := s.client.doRequest(ctx, http.MethodGet, path, params, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateProductStocks updates stock levels for one or more products.
func (s *StockService) UpdateProductStocks(ctx context.Context, req models.StockUpdateRequest) error {
	return s.client.doRequest(ctx, http.MethodPut, "/public-api/v1/products-stocks", nil, req, nil)
}
