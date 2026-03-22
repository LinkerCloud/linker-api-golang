package linkercloud

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNew_defaults(t *testing.T) {
	c := New("https://api-example.linker.shop", "test-key")

	if c.baseURL != "https://api-example.linker.shop" {
		t.Errorf("baseURL: got %q", c.baseURL)
	}
	if c.Orders == nil || c.Products == nil || c.Stock == nil {
		t.Error("expected all services to be initialised")
	}
}

func TestNew_trailingSlashStripped(t *testing.T) {
	c := New("https://api-example.linker.shop/", "key")
	if c.baseURL != "https://api-example.linker.shop" {
		t.Errorf("trailing slash not stripped: %q", c.baseURL)
	}
}

func TestNew_withTimeout(t *testing.T) {
	c := New("https://api-example.linker.shop", "key", WithTimeout(5*time.Second))
	if c.httpClient.Timeout != 5*time.Second {
		t.Errorf("timeout: got %v, want 5s", c.httpClient.Timeout)
	}
}

func TestNew_withCustomHTTPClient(t *testing.T) {
	custom := &http.Client{Timeout: 10 * time.Second}
	c := New("https://api-example.linker.shop", "key", WithHTTPClient(custom))
	if c.httpClient == nil {
		t.Fatal("httpClient is nil")
	}
}

func TestAuthTransport_injectsAPIKey(t *testing.T) {
	const wantKey = "my-secret-api-key"
	var gotKey string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotKey = r.Header.Get(headerAPIKey)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("null"))
	}))
	defer srv.Close()

	c := New(srv.URL, wantKey)
	_ = c.doRequest(context.Background(), http.MethodGet, "/test", nil, nil, nil)

	if gotKey != wantKey {
		t.Errorf("apikey header: got %q, want %q", gotKey, wantKey)
	}
}

func TestDoRequest_decodesJSONResponse(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload{Name: "hello"})
	}))
	defer srv.Close()

	c := New(srv.URL, "key")
	var got payload
	if err := c.doRequest(context.Background(), http.MethodGet, "/", nil, nil, &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "hello" {
		t.Errorf("Name: got %q, want %q", got.Name, "hello")
	}
}

func TestDoRequest_returnsAPIErrorOnNon2xx(t *testing.T) {
	cases := []struct {
		code int
		body string
	}{
		{400, "bad request"},
		{401, "unauthorized"},
		{404, "not found"},
		{500, "server error"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(http.StatusText(tc.code), func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, tc.body, tc.code)
			}))
			defer srv.Close()

			c := New(srv.URL, "key")
			err := c.doRequest(context.Background(), http.MethodGet, "/", nil, nil, nil)
			if err == nil {
				t.Fatal("expected error, got nil")
			}

			apiErr, ok := err.(*APIError)
			if !ok {
				t.Fatalf("expected *APIError, got %T: %v", err, err)
			}
			if apiErr.StatusCode != tc.code {
				t.Errorf("StatusCode: got %d, want %d", apiErr.StatusCode, tc.code)
			}
		})
	}
}

func TestIsNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer srv.Close()

	c := New(srv.URL, "key")
	err := c.doRequest(context.Background(), http.MethodGet, "/", nil, nil, nil)

	if !IsNotFound(err) {
		t.Errorf("IsNotFound: got false, want true for err=%v", err)
	}
	if IsUnauthorized(err) {
		t.Error("IsUnauthorized should be false")
	}
}

func TestDoRequest_encodesRequestBody(t *testing.T) {
	type payload struct {
		Value int `json:"value"`
	}

	var gotBody payload
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&gotBody)
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	c := New(srv.URL, "key")
	_ = c.doRequest(context.Background(), http.MethodPost, "/", nil, payload{Value: 42}, nil)

	if gotBody.Value != 42 {
		t.Errorf("body value: got %d, want 42", gotBody.Value)
	}
}

func TestDoRequest_appendsQueryParams(t *testing.T) {
	var gotQuery string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := New(srv.URL, "key")
	params := buildParams(map[string]string{"limit": "50", "offset": "0"})
	_ = c.doRequest(context.Background(), http.MethodGet, "/", params, nil, nil)

	if gotQuery == "" {
		t.Error("expected query string to be set")
	}
}

func TestClient_maxResponseSize(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Write a JSON body that exceeds 16 bytes.
		w.Write([]byte(`{"data":"` + strings.Repeat("x", 100) + `"}`))
	})
	c.maxResponseSize = 16

	var result map[string]any
	err := c.doRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	if err == nil {
		t.Fatal("expected error from truncated response, got nil")
	}
	if !strings.Contains(err.Error(), "decode response") {
		t.Errorf("expected JSON decode error, got: %v", err)
	}
}

func TestClient_retryOn429(t *testing.T) {
	var calls int
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok":true}`))
	})

	var result map[string]any
	err := c.doRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)
	if err != nil {
		t.Fatalf("expected success after retry, got: %v", err)
	}
	if calls != 2 {
		t.Errorf("expected 2 requests, got %d", calls)
	}
}

func TestClient_retryOn429_maxOneRetry(t *testing.T) {
	var calls int
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Retry-After", "0")
		w.WriteHeader(http.StatusTooManyRequests)
	})

	err := c.doRequest(context.Background(), http.MethodGet, "/test", nil, nil, nil)
	if !IsTooManyRequests(err) {
		t.Fatalf("expected IsTooManyRequests, got: %v", err)
	}
	if calls != 2 {
		t.Errorf("expected exactly 2 requests, got %d", calls)
	}
}

func TestParseRetryAfter(t *testing.T) {
	tests := []struct {
		val  string
		want time.Duration
	}{
		{"", time.Second},
		{"2", 2 * time.Second},
		{"0", time.Second},
		{"-5", time.Second},
		{"999", 30 * time.Second},
		{"garbage", time.Second},
	}
	for _, tt := range tests {
		got := parseRetryAfter(tt.val)
		if got != tt.want {
			t.Errorf("parseRetryAfter(%q) = %v, want %v", tt.val, got, tt.want)
		}
	}
}

func TestAPIError_messageExtraction(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code":400,"message":"field required"}`))
	})

	err := c.doRequest(context.Background(), http.MethodGet, "/test", nil, nil, nil)
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Message != "field required" {
		t.Errorf("Message: got %q, want %q", apiErr.Message, "field required")
	}
	if apiErr.Body == "" {
		t.Error("Body should contain full response")
	}
	if !strings.Contains(apiErr.Error(), "field required") {
		t.Errorf("Error() should contain message: %s", apiErr.Error())
	}
}

func TestAPIError_noMessage(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
	})

	err := c.doRequest(context.Background(), http.MethodGet, "/test", nil, nil, nil)
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.Message != "" {
		t.Errorf("Message should be empty for non-JSON body, got %q", apiErr.Message)
	}
	if apiErr.Error() != "linkercloud: HTTP 500" {
		t.Errorf("Error(): got %q, want %q", apiErr.Error(), "linkercloud: HTTP 500")
	}
}

func TestNew_inheritsClientTimeout(t *testing.T) {
	custom := &http.Client{Timeout: 15 * time.Second}
	c := New("https://example.com", "key", WithHTTPClient(custom))
	if c.httpClient.Timeout != 15*time.Second {
		t.Errorf("Timeout: got %v, want 15s", c.httpClient.Timeout)
	}
}

func TestNew_explicitTimeoutOverridesClient(t *testing.T) {
	custom := &http.Client{Timeout: 15 * time.Second}
	c := New("https://example.com", "key", WithHTTPClient(custom), WithTimeout(5*time.Second))
	if c.httpClient.Timeout != 5*time.Second {
		t.Errorf("Timeout: got %v, want 5s", c.httpClient.Timeout)
	}
}
