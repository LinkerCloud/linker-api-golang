package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestShipmentsService_Create(t *testing.T) {
	want := models.CreateShipmentsResponse{LastID: "del-1", Packages: []string{"del-1"}}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/deliveries")
		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.Shipments.Create(context.Background(), models.CreateShipmentsRequest{IDs: []string{"1"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.LastID != "del-1" {
		t.Errorf("LastID: got %q, want %q", got.LastID, "del-1")
	}
}

func TestShipmentsService_CreatePackages(t *testing.T) {
	want := []models.ParcelType{{TrackingNumber: "TRACK-001"}}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/deliveries/packages")
		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.Shipments.CreatePackages(context.Background(), models.ShipmentType{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || got[0].TrackingNumber != "TRACK-001" {
		t.Errorf("unexpected result: %+v", got)
	}
}

func TestShipmentsService_UpdateStatus(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/deliveries/statuses")
		w.WriteHeader(http.StatusNoContent)
	})

	if err := c.Shipments.UpdateStatus(context.Background(), models.DeliveryStatus{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestShipmentsService_CancelPackages(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPatch)
		testutil.AssertPath(t, r, "/public-api/v1/deliveries/order-42")
		w.WriteHeader(http.StatusNoContent)
	})

	if err := c.Shipments.CancelPackages(context.Background(), "order-42", models.CancelPackagesRequest{IDs: []string{"pkg-1"}}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestShipmentsService_GetLabels(t *testing.T) {
	// "bGFiZWwx" is base64 for "label1", "bGFiZWwy" is base64 for "label2"
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/order-1/pkg-1/PDF/parcel-1")
		testutil.RespondJSON(t, w, http.StatusOK, []string{"bGFiZWwx", "bGFiZWwy"})
	})

	got, err := c.Shipments.GetLabels(context.Background(), GetLabelsOptions{
		OrderID:   "order-1",
		PackageID: "pkg-1",
		Format:    LabelFormatPDF,
		ParcelID:  "parcel-1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len(labels): got %d, want 2", len(got))
	}
	if string(got[0]) != "label1" {
		t.Errorf("label[0]: got %q, want %q", got[0], "label1")
	}
	if string(got[1]) != "label2" {
		t.Errorf("label[1]: got %q, want %q", got[1], "label2")
	}
}

func TestShipmentsService_GetLabels_optionalParcelID(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertPath(t, r, "/public-api/v1/ord/pkg/PNG/")
		testutil.RespondJSON(t, w, http.StatusOK, []string{})
	})

	got, err := c.Shipments.GetLabels(context.Background(), GetLabelsOptions{
		OrderID:   "ord",
		PackageID: "pkg",
		Format:    LabelFormatPNG,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected empty labels, got %d", len(got))
	}
}

func TestShipmentsService_CancelPackages_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "not found")
	})

	err := c.Shipments.CancelPackages(context.Background(), "nope", models.CancelPackagesRequest{IDs: []string{"pkg-1"}})
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}
