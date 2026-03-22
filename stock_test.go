package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func int64Ptr(v int64) *int64 { return &v }

func TestStockService_List(t *testing.T) {
	want := models.Stock{TotalCount: int64Ptr(2)}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/stocks")
		testutil.AssertQuery(t, r, "limit", "25")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.Stock.List(context.Background(), ListStocksOptions{Limit: 25})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.TotalCount == nil || *got.TotalCount != *want.TotalCount {
		t.Errorf("TotalCount: got %v, want %v", got.TotalCount, want.TotalCount)
	}
}

func TestStockService_List_hideIgnoredInWMS(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertQuery(t, r, "hideIgnoredInWMS", "true")
		testutil.RespondJSON(t, w, http.StatusOK, models.Stock{})
	})

	_, err := c.Stock.List(context.Background(), ListStocksOptions{HideIgnoredInWMS: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestStockService_GenerateReport(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/generate-wms-stock-report")
		testutil.AssertQuery(t, r, "sku", "MY-SKU")
		testutil.RespondJSON(t, w, http.StatusCreated, map[string]string{"uuid": "abc-123"})
	})

	uuid, err := c.Stock.GenerateReport(context.Background(), GenerateWMSStockReportOptions{SKU: "MY-SKU"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if uuid != "abc-123" {
		t.Errorf("uuid: got %q, want %q", uuid, "abc-123")
	}
}

func TestStockService_GetReport(t *testing.T) {
	want := models.Stock{TotalCount: int64Ptr(100)}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/wms-stock-report/some-uuid")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.Stock.GetReport(context.Background(), "some-uuid", GetReportOptions{Limit: 50})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.TotalCount == nil || *got.TotalCount != 100 {
		t.Errorf("TotalCount: got %v, want 100", got.TotalCount)
	}
}

func TestStockService_GetReport_pending(t *testing.T) {
	// 202 Accepted means the report is still being generated.
	// Since 202 is a valid 2xx response, the client returns nil error and an
	// empty Stock. The caller should poll until the response is 200.
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})

	got, err := c.Stock.GetReport(context.Background(), "pending-uuid", GetReportOptions{})
	if err != nil {
		t.Fatalf("unexpected error for 202 response: %v", err)
	}
	// An empty body results in a zero-value struct
	if got.TotalCount != nil {
		t.Errorf("expected nil TotalCount for pending report, got %v", got.TotalCount)
	}
}

func TestStockService_UpdateProductStocks(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPut)
		testutil.AssertPath(t, r, "/public-api/v1/products-stocks")
		w.WriteHeader(http.StatusAccepted)
	})

	if err := c.Stock.UpdateProductStocks(context.Background(), models.StockUpdateRequest{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
