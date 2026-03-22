package linkercloud

import (
	"errors"
	"fmt"
)

// APIError represents an error response from the Linker Cloud API.
// It carries the HTTP status code, an extracted human-readable message (when
// available), and the full raw response body.
type APIError struct {
	// StatusCode is the HTTP status code returned by the server (e.g. 400, 401, 404).
	StatusCode int
	// Message is the human-readable error extracted from the JSON response
	// body, if available. Empty when the body isn't JSON or has no "message" field.
	Message string
	// Body is the full raw response body from the server.
	Body string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("linkercloud: HTTP %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("linkercloud: HTTP %d", e.StatusCode)
}

// IsNotFound reports whether err is a 404 Not Found API error.
func IsNotFound(err error) bool { return isAPIErrorWithCode(err, 404) }

// IsUnauthorized reports whether err is a 401 Unauthorized API error.
func IsUnauthorized(err error) bool { return isAPIErrorWithCode(err, 401) }

// IsForbidden reports whether err is a 403 Forbidden API error.
func IsForbidden(err error) bool { return isAPIErrorWithCode(err, 403) }

// IsConflict reports whether err is a 409 Conflict API error.
func IsConflict(err error) bool { return isAPIErrorWithCode(err, 409) }

// IsTooManyRequests reports whether err is a 429 Too Many Requests API error.
func IsTooManyRequests(err error) bool { return isAPIErrorWithCode(err, 429) }

func isAPIErrorWithCode(err error, code int) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.StatusCode == code
}
