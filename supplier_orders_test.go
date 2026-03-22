package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestSupplierOrdersService_List(t *testing.T) {
	want := Page[models.SupplierOrder]{
		Items:        []models.SupplierOrder{{ID: "s1"}, {ID: "s2"}},
		RecordsTotal: 2,
	}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/supplierorders")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.SupplierOrders.List(context.Background(), ListSupplierOrdersOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Items) != 2 {
		t.Errorf("len(Items): got %d, want 2", len(got.Items))
	}
}

func TestSupplierOrdersService_Get(t *testing.T) {
	want := models.SupplierOrder{ID: "s42"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertPath(t, r, "/public-api/v1/supplierorders/s42")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.SupplierOrders.Get(context.Background(), "s42")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != "s42" {
		t.Errorf("ID: got %q, want %q", got.ID, "s42")
	}
}

func TestSupplierOrdersService_Create(t *testing.T) {
	req := models.SupplierOrderType{ClientOrderNumber: "ext-1"}
	want := models.SupplierOrderType{ClientOrderNumber: "ext-1"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/supplierorders")
		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.SupplierOrders.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ClientOrderNumber != want.ClientOrderNumber {
		t.Errorf("ClientOrderNumber: got %q, want %q", got.ClientOrderNumber, want.ClientOrderNumber)
	}
}

func TestSupplierOrdersService_Confirm(t *testing.T) {
	req := models.SupplierOrderType{ClientOrderNumber: "conf-1"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/supplierorders/confirms")
		testutil.RespondJSON(t, w, http.StatusCreated, req)
	})

	got, err := c.SupplierOrders.Confirm(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ClientOrderNumber != req.ClientOrderNumber {
		t.Errorf("ClientOrderNumber: got %q, want %q", got.ClientOrderNumber, req.ClientOrderNumber)
	}
}

func TestSupplierOrdersService_Get_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "not found")
	})

	_, err := c.SupplierOrders.Get(context.Background(), "nope")
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}
