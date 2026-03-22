package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestOrderReturnsService_Get(t *testing.T) {
	want := models.OrderReturn{ID: "ret-1"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orderreturns/ret-1")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.OrderReturns.Get(context.Background(), "ret-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != want.ID {
		t.Errorf("ID: got %q, want %q", got.ID, want.ID)
	}
}

func TestOrderReturnsService_List(t *testing.T) {
	want := Page[models.OrderReturn]{
		Items:        []models.OrderReturn{{ID: "r1"}, {ID: "r2"}},
		RecordsTotal: 2,
	}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orderreturns")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.OrderReturns.List(context.Background(), ListOrderReturnsOptions{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Items) != len(want.Items) {
		t.Errorf("len(Items): got %d, want %d", len(got.Items), len(want.Items))
	}
}

func TestOrderReturnsService_Create(t *testing.T) {
	want := models.OrderReturn{ID: "new-ret"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/orderreturns")
		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.OrderReturns.Create(context.Background(), models.OrderReturnType{
		OrderNumber: "ORD-1",
		ReturnType:  "RETURN",
		Items: []models.OrderReturnItemType{
			{SKU: "SKU-1", Quantity: 1, QuantityByCondition: map[string]int{"good": 1}},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != want.ID {
		t.Errorf("ID: got %q, want %q", got.ID, want.ID)
	}
}

func TestOrderReturnsService_Accept(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/orderreturns/ret-1/accept")
		w.WriteHeader(http.StatusCreated)
	})

	if err := c.OrderReturns.Accept(context.Background(), "ret-1"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrderReturnsService_Get_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "not found")
	})

	_, err := c.OrderReturns.Get(context.Background(), "nope")
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}
