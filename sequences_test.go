package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
)

func TestSequencesService_Get(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/sequences/order_number")
		testutil.RespondJSON(t, w, http.StatusOK, 1042)
	})

	got, err := c.Sequences.Get(context.Background(), "order_number")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 1042 {
		t.Errorf("got %d, want 1042", got)
	}
}

func TestSequencesService_Get_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "not found")
	})

	_, err := c.Sequences.Get(context.Background(), "nope")
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}
