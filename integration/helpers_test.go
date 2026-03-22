//go:build integration

package integration

import (
	"os"
	"testing"
	"time"

	linkercloud "github.com/linkercloud/linker-api-golang"
)

// newClient creates a [linkercloud.Client] from environment variables.
// The test is skipped if LINKER_BASE_URL or LINKER_API_KEY are not set.
func newClient(t *testing.T) *linkercloud.Client {
	t.Helper()
	baseURL := os.Getenv("LINKER_BASE_URL")
	apiKey := os.Getenv("LINKER_API_KEY")
	if baseURL == "" || apiKey == "" {
		t.Skip("LINKER_BASE_URL and LINKER_API_KEY must be set to run integration tests")
	}
	return linkercloud.New(baseURL, apiKey, linkercloud.WithTimeout(15*time.Second))
}

// requireWritesAllowed skips the test unless LINKER_ALLOW_WRITES=1.
// Call this at the top of every mutating (POST / PUT / PATCH) integration test to
// prevent accidental data changes against a production instance.
//
// Example:
//
//	func TestIntegration_Write_Orders_Create(t *testing.T) {
//	    requireWritesAllowed(t)
//	    client := newClient(t)
//	    ...
//	}
func requireWritesAllowed(t *testing.T) {
	t.Helper()
	if os.Getenv("LINKER_ALLOW_WRITES") != "1" {
		t.Skip("set LINKER_ALLOW_WRITES=1 to run write integration tests")
	}
}
