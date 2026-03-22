package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestPackingService_GetContainers(t *testing.T) {
	want := models.PackingContainer{}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orders/ord-1/packingcontainers")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	_, err := c.Packing.GetContainers(context.Background(), "ord-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPackingService_GetMaterials(t *testing.T) {
	want := models.OrderMaterial{}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orders/ord-1/packingmaterials")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	_, err := c.Packing.GetMaterials(context.Background(), "ord-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPackingService_GetContainers_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "order not found")
	})

	_, err := c.Packing.GetContainers(context.Background(), "unknown")
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}
