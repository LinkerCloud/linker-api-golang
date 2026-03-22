// Package linkercloud provides a Go client for the Linker Cloud Public API.
//
// # Authentication
//
// All requests are authenticated via an API key passed as the "apikey" HTTP
// header. You should receive your key from your operator or Linker customer
// service.
//
// # Usage
//
//	client := linkercloud.New("https://api-mycompany.linker.shop", "my-api-key")
//
//	orders, err := client.Orders.List(ctx, linkercloud.ListOrdersOptions{Limit: 50})
package linkercloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	defaultTimeout         = 30 * time.Second
	defaultMaxResponseSize = 10 << 20 // 10 MB
	headerAPIKey           = "apikey"
	headerAccept           = "Accept"
	headerContentType      = "Content-Type"
	contentTypeJSON        = "application/json"
)

// Client is the Linker Cloud API client. Create one with [New] and use the
// service fields to interact with specific API resources.
//
// Client is safe for concurrent use.
type Client struct {
	baseURL         string
	httpClient      *http.Client
	maxResponseSize int64

	// Orders provides access to order management endpoints.
	Orders *OrdersService
	// OrderReturns provides access to order return endpoints.
	OrderReturns *OrderReturnsService
	// OrderAttachments provides access to order attachment endpoints.
	OrderAttachments *OrderAttachmentsService
	// OrderTransitions provides access to order state transition endpoints.
	OrderTransitions *OrderTransitionsService
	// Packing provides access to packing container and material endpoints.
	Packing *PackingService
	// Products provides access to product management endpoints.
	Products *ProductsService
	// Stock provides access to stock and inventory endpoints.
	Stock *StockService
	// Shipments provides access to delivery and shipment endpoints.
	Shipments *ShipmentsService
	// Tracking provides access to tracking number endpoints.
	Tracking *TrackingService
	// Pickups provides access to pickup scheduling endpoints.
	Pickups *PickupsService
	// SupplierOrders provides access to inbound/supplier order endpoints.
	SupplierOrders *SupplierOrdersService
	// Workflows provides access to workflow endpoints.
	Workflows *WorkflowsService
	// Sequences provides access to sequence endpoints.
	Sequences *SequencesService
	// Webhooks provides access to integration webhook endpoints.
	Webhooks *WebhooksService
}

// Option configures a [Client].
type Option func(*clientConfig)

type clientConfig struct {
	httpClient      *http.Client
	timeout         time.Duration
	maxResponseSize int64
	timeoutSet      bool
}

// WithHTTPClient sets a custom [http.Client] to use for all requests.
// Use this to configure proxies, custom TLS settings, or test transports.
func WithHTTPClient(hc *http.Client) Option {
	return func(cfg *clientConfig) { cfg.httpClient = hc }
}

// WithTimeout sets the per-request timeout. Defaults to 30 s.
// Note: the Linker Cloud API gateway enforces a 30 s maximum — requests
// exceeding this will receive a 50x from the gateway regardless of this setting.
func WithTimeout(d time.Duration) Option {
	return func(cfg *clientConfig) { cfg.timeout = d; cfg.timeoutSet = true }
}

// WithMaxResponseSize sets the maximum response body size in bytes.
// Responses larger than this are truncated. Defaults to 10 MB.
func WithMaxResponseSize(n int64) Option {
	return func(cfg *clientConfig) { cfg.maxResponseSize = n }
}

// New creates a new [Client] targeting the given baseURL and authenticating
// with apiKey.
//
// baseURL should be the root of the Linker Cloud instance, e.g.
// "https://api-mycompany.linker.shop". A trailing slash is handled
// automatically.
//
// apiKey is passed on every request as the "apikey" HTTP header.
func New(baseURL, apiKey string, opts ...Option) *Client {
	cfg := &clientConfig{timeout: defaultTimeout, maxResponseSize: defaultMaxResponseSize}
	for _, o := range opts {
		o(cfg)
	}

	// If a custom HTTP client was provided but no explicit timeout, inherit
	// the client's timeout. Fall back to the default if it's zero.
	if cfg.httpClient != nil && !cfg.timeoutSet {
		cfg.timeout = cfg.httpClient.Timeout
		if cfg.timeout == 0 {
			cfg.timeout = defaultTimeout
		}
	}

	base := strings.TrimRight(baseURL, "/")

	var transport http.RoundTripper
	if cfg.httpClient != nil {
		// Wrap the caller-supplied transport so the API key is always injected.
		transport = &authTransport{
			apiKey: apiKey,
			inner:  cfg.httpClient.Transport,
		}
	} else {
		transport = &authTransport{
			apiKey: apiKey,
			inner:  http.DefaultTransport,
		}
	}

	hc := &http.Client{
		Timeout:   cfg.timeout,
		Transport: transport,
	}
	if cfg.httpClient != nil {
		hc.CheckRedirect = cfg.httpClient.CheckRedirect
		hc.Jar = cfg.httpClient.Jar
	}

	c := &Client{
		baseURL:         base,
		httpClient:      hc,
		maxResponseSize: cfg.maxResponseSize,
	}

	c.Orders = &OrdersService{client: c}
	c.OrderReturns = &OrderReturnsService{client: c}
	c.OrderAttachments = &OrderAttachmentsService{client: c}
	c.OrderTransitions = &OrderTransitionsService{client: c}
	c.Packing = &PackingService{client: c}
	c.Products = &ProductsService{client: c}
	c.Stock = &StockService{client: c}
	c.Shipments = &ShipmentsService{client: c}
	c.Tracking = &TrackingService{client: c}
	c.Pickups = &PickupsService{client: c}
	c.SupplierOrders = &SupplierOrdersService{client: c}
	c.Workflows = &WorkflowsService{client: c}
	c.Sequences = &SequencesService{client: c}
	c.Webhooks = &WebhooksService{client: c}

	return c
}

// authTransport injects the Linker Cloud API key header on every outbound request.
type authTransport struct {
	apiKey string
	inner  http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid mutating the original (required by RoundTripper contract).
	clone := req.Clone(req.Context())
	clone.Header.Set(headerAPIKey, t.apiKey)
	return t.inner.RoundTrip(clone)
}

// doRequest is the single HTTP choke-point used by all services.
//
// It builds and executes an HTTP request, decodes a successful (2xx) response
// into result (if non-nil), and returns an [*APIError] for any non-2xx status.
//
// On a 429 (Too Many Requests) response, it honours the Retry-After header and
// retries exactly once. If the retry also fails, the error is returned normally.
//
// params are appended to the URL as query parameters.
// body, if non-nil, is JSON-encoded and sent as the request body.
// result, if non-nil, receives the JSON-decoded response body on success.
func (c *Client) doRequest(ctx context.Context, method, path string, params url.Values, body, result any) error {
	u := c.baseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}

	var buf []byte
	if body != nil {
		var err error
		buf, err = json.Marshal(body)
		if err != nil {
			return fmt.Errorf("linkercloud: marshal request body: %w", err)
		}
	}

	for attempt := 0; attempt < 2; attempt++ {
		var bodyReader io.Reader
		if buf != nil {
			bodyReader = bytes.NewReader(buf)
		}

		req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
		if err != nil {
			return fmt.Errorf("linkercloud: build request: %w", err)
		}
		req.Header.Set(headerAccept, contentTypeJSON)
		if buf != nil {
			req.Header.Set(headerContentType, contentTypeJSON)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("linkercloud: execute request: %w", err)
		}

		respBody, err := io.ReadAll(io.LimitReader(resp.Body, c.maxResponseSize))
		resp.Body.Close()
		if err != nil {
			return fmt.Errorf("linkercloud: read response body: %w", err)
		}

		// On 429, retry once after honouring Retry-After.
		if resp.StatusCode == http.StatusTooManyRequests && attempt == 0 {
			if err := sleepRetryAfter(ctx, resp.Header); err != nil {
				return newAPIError(resp.StatusCode, respBody)
			}
			continue
		}

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return newAPIError(resp.StatusCode, respBody)
		}

		if result != nil && len(respBody) > 0 {
			if err := json.Unmarshal(respBody, result); err != nil {
				return fmt.Errorf("linkercloud: decode response: %w", err)
			}
		}
		return nil
	}
	return nil // unreachable
}

const maxRetryAfter = 30 * time.Second

// sleepRetryAfter waits for the duration specified by the Retry-After header.
// Returns nil on successful wait, or ctx.Err() if the context is cancelled.
func sleepRetryAfter(ctx context.Context, h http.Header) error {
	d := parseRetryAfter(h.Get("Retry-After"))
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// parseRetryAfter parses a Retry-After header value as either integer seconds
// or an HTTP-date. Returns 1 s on missing/invalid values, capped at 30 s.
func parseRetryAfter(val string) time.Duration {
	if val == "" {
		return time.Second
	}
	if secs, err := strconv.Atoi(val); err == nil {
		d := time.Duration(secs) * time.Second
		if d > maxRetryAfter {
			return maxRetryAfter
		}
		if d <= 0 {
			return time.Second
		}
		return d
	}
	if t, err := http.ParseTime(val); err == nil {
		d := time.Until(t)
		if d > maxRetryAfter {
			return maxRetryAfter
		}
		if d <= 0 {
			return time.Second
		}
		return d
	}
	return time.Second
}

// newAPIError creates an [*APIError] from a status code and raw response body,
// extracting a human-readable message from JSON when possible.
func newAPIError(statusCode int, body []byte) *APIError {
	apiErr := &APIError{
		StatusCode: statusCode,
		Body:       strings.TrimSpace(string(body)),
	}
	var parsed struct {
		Message string `json:"message"`
	}
	if json.Unmarshal(body, &parsed) == nil && parsed.Message != "" {
		apiErr.Message = parsed.Message
	}
	return apiErr
}

// buildParams builds a [url.Values] from a map, keeping it nil when empty.
// Convenience helper used by services.
func buildParams(kv map[string]string) url.Values {
	if len(kv) == 0 {
		return nil
	}
	v := make(url.Values, len(kv))
	for k, val := range kv {
		if val != "" {
			v.Set(k, val)
		}
	}
	return v
}
