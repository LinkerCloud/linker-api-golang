// Package testutil provides httptest helpers for unit-testing the Linker Cloud
// client without making real network calls.
package testutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// NewServer starts a test HTTP server backed by handler and returns it.
// The server is automatically closed via t.Cleanup when the test ends.
func NewServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(handler))
	t.Cleanup(srv.Close)
	return srv
}

// RespondJSON writes status and body (JSON-encoded) to w.
func RespondJSON(t *testing.T, w http.ResponseWriter, status int, body any) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			t.Errorf("testutil: encode response: %v", err)
		}
	}
}

// RespondError writes an HTTP error status with a plain-text body.
func RespondError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}

// AssertHeader fails the test if the request does not carry the expected header value.
func AssertHeader(t *testing.T, r *http.Request, key, want string) {
	t.Helper()
	if got := r.Header.Get(key); got != want {
		t.Errorf("header %q: got %q, want %q", key, got, want)
	}
}

// AssertMethod fails the test if the request method does not match.
func AssertMethod(t *testing.T, r *http.Request, method string) {
	t.Helper()
	if r.Method != method {
		t.Errorf("method: got %q, want %q", r.Method, method)
	}
}

// AssertPath fails the test if the request URL path does not match.
func AssertPath(t *testing.T, r *http.Request, path string) {
	t.Helper()
	if r.URL.Path != path {
		t.Errorf("path: got %q, want %q", r.URL.Path, path)
	}
}

// AssertQuery fails the test if the request URL query param does not match.
func AssertQuery(t *testing.T, r *http.Request, key, want string) {
	t.Helper()
	if got := r.URL.Query().Get(key); got != want {
		t.Errorf("query param %q: got %q, want %q", key, got, want)
	}
}

// DecodeBody decodes the JSON request body into dst.
func DecodeBody(t *testing.T, r *http.Request, dst any) {
	t.Helper()
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		t.Fatalf("testutil: decode request body: %v", err)
	}
}
