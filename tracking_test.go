package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestTrackingService_SetTrackingNumber(t *testing.T) {
	want := models.Order2{Number: "ord-1"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPut)
		testutil.AssertPath(t, r, "/public-api/v1/orders/ord-1/trackingnumber")
		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.Tracking.SetTrackingNumber(context.Background(), "ord-1", models.TrackingNumberRequest{TrackingNumber: "TRACK-001", Operator: "DHL"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Number != want.Number {
		t.Errorf("Number: got %q, want %q", got.Number, want.Number)
	}
}

func TestTrackingService_SetTrackingNumber_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "not found")
	})

	_, err := c.Tracking.SetTrackingNumber(context.Background(), "nope", models.TrackingNumberRequest{TrackingNumber: "T1", Operator: "DHL"})
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}

func TestTrackingService_SetTrackingNumbers(t *testing.T) {
	want := models.Order2{Number: "ord-2"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPut)
		testutil.AssertPath(t, r, "/public-api/v1/orders/ord-2/trackingnumbers")
		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.Tracking.SetTrackingNumbers(context.Background(), "ord-2", models.TrackingNumbersRequest{
		TrackingNumbers: []models.TrackingNumberRequest{
			{TrackingNumber: "T1", Operator: "DHL"},
			{TrackingNumber: "T2", Operator: "DHL"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Number != want.Number {
		t.Errorf("Number: got %q, want %q", got.Number, want.Number)
	}
}
