package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestPickupsService_Create(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/pickups")
		w.WriteHeader(http.StatusCreated)
	})

	if err := c.Pickups.Create(context.Background(), models.PickupRequest{
		PickupDate: "2026-03-22",
		Provider:   "DHL",
	}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
