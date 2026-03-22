package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestProductsService_List(t *testing.T) {
	want := Page[models.Product]{Items: []models.Product{{ID: "p1", SKU: "SKU-001"}}, RecordsTotal: 1}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/products")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.Products.List(context.Background(), ListProductsOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Items) != 1 || got.Items[0].SKU != "SKU-001" {
		t.Errorf("unexpected result: %+v", got)
	}
}

func TestProductsService_Create(t *testing.T) {
	req := models.ProductType{SKU: "NEW-SKU"}
	want := models.Product{ID: "new-id", SKU: "NEW-SKU"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/products")

		var body models.ProductType
		testutil.DecodeBody(t, r, &body)
		if body.SKU != "NEW-SKU" {
			t.Errorf("SKU: got %q, want %q", body.SKU, "NEW-SKU")
		}

		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.Products.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != want.ID {
		t.Errorf("ID: got %q, want %q", got.ID, want.ID)
	}
}

func TestProductsService_Update(t *testing.T) {
	req := models.ProductType{SKU: "UPDATED-SKU"}
	want := models.Product{ID: "p1", SKU: "UPDATED-SKU"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPut)
		testutil.AssertPath(t, r, "/public-api/v1/products/p1")
		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.Products.Update(context.Background(), "p1", req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.SKU != want.SKU {
		t.Errorf("SKU: got %q, want %q", got.SKU, want.SKU)
	}
}

func TestProductsService_List_withFilters(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertQuery(t, r, "filters[sku]", "SKU-42")
		testutil.RespondJSON(t, w, http.StatusOK, Page[models.Product]{})
	})

	_, err := c.Products.List(context.Background(), ListProductsOptions{
		Filters: map[string]string{"sku": "SKU-42"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestProductsService_Update_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "not found")
	})

	_, err := c.Products.Update(context.Background(), "nope", models.ProductType{})
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}
