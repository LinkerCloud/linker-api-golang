package linkercloud

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	srv := testutil.NewServer(t, handler)
	return New(srv.URL, "test-key"), srv
}

func TestOrdersService_Get(t *testing.T) {
	want := models.Order{ID: "42", Number: "ORD-001"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orders/42")
		testutil.AssertHeader(t, r, headerAPIKey, "test-key")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.Orders.Get(context.Background(), "42")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != want.ID || got.Number != want.Number {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestOrdersService_Get_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "not found")
	})

	_, err := c.Orders.Get(context.Background(), "999")
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}

func TestOrdersService_List(t *testing.T) {
	want := Page[models.Order]{
		Items:        []models.Order{{ID: "1"}, {ID: "2"}},
		RecordsTotal: 2,
	}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orders")
		testutil.AssertQuery(t, r, "limit", "10")
		testutil.AssertQuery(t, r, "sortDir", "ASC")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.Orders.List(context.Background(), ListOrdersOptions{Limit: 10, SortDir: "ASC"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Items) != len(want.Items) {
		t.Errorf("len(Items): got %d, want %d", len(got.Items), len(want.Items))
	}
	if got.RecordsTotal != want.RecordsTotal {
		t.Errorf("RecordsTotal: got %d, want %d", got.RecordsTotal, want.RecordsTotal)
	}
}

func TestOrdersService_List_withFilters(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertQuery(t, r, "filters[order_date]", "15.11.2021")
		testutil.RespondJSON(t, w, http.StatusOK, Page[models.Order]{})
	})

	_, err := c.Orders.List(context.Background(), ListOrdersOptions{
		Filters: map[string]string{"order_date": "15.11.2021"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrdersService_Create(t *testing.T) {
	req := models.OrderType{ClientOrderNumber: "EXT-001"}
	want := models.Order2{Number: "EXT-001"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/orders")

		var body models.OrderType
		testutil.DecodeBody(t, r, &body)
		if body.ClientOrderNumber != req.ClientOrderNumber {
			t.Errorf("clientOrderNumber: got %q, want %q", body.ClientOrderNumber, req.ClientOrderNumber)
		}

		testutil.RespondJSON(t, w, http.StatusCreated, want)
	})

	got, err := c.Orders.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Number != want.Number {
		t.Errorf("Number: got %q, want %q", got.Number, want.Number)
	}
}

func TestOrdersService_Cancel(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPut)
		testutil.AssertPath(t, r, "/public-api/v1/orders/77/cancel")
		w.WriteHeader(http.StatusOK)
	})

	if err := c.Orders.Cancel(context.Background(), "77"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrdersService_UpdatePaymentStatus(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPut)
		testutil.AssertPath(t, r, "/public-api/v1/payment-status")

		var body map[string]any
		testutil.DecodeBody(t, r, &body)

		w.WriteHeader(http.StatusAccepted)
	})

	req := models.PaymentRequest{}
	if err := c.Orders.UpdatePaymentStatus(context.Background(), req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrdersService_ConfirmPicking(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPut)
		testutil.AssertPath(t, r, "/public-api/v1/order/pickingconfirmation")
		w.WriteHeader(http.StatusCreated)
	})

	if err := c.Orders.ConfirmPicking(context.Background(), models.PickingConfirmation{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrdersService_GetOriginalItems(t *testing.T) {
	want := []models.OrderItem{{SKU: "SKU-A"}, {SKU: "SKU-B"}}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertPath(t, r, "/public-api/v1/orders/5/original/items")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.Orders.GetOriginalItems(context.Background(), "5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("len: got %d, want 2", len(got))
	}
}

func TestOrdersService_errorCases(t *testing.T) {
	cases := []struct {
		name     string
		code     int
		checkErr func(error) bool
	}{
		{"unauthorized", 401, IsUnauthorized},
		{"forbidden", 403, IsForbidden},
		{"conflict", 409, IsConflict},
		{"too_many_requests", 429, IsTooManyRequests},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, tc.name, tc.code)
			})

			_, err := c.Orders.Get(context.Background(), "1")
			if !tc.checkErr(err) {
				t.Errorf("expected error check to pass for code %d, got: %v", tc.code, err)
			}
		})
	}
}

// Ensure the JSON tags round-trip correctly for a key Order field.
func TestOrderType_jsonRoundTrip(t *testing.T) {
	orig := models.OrderType{
		ClientOrderNumber: "TEST-123",
	}
	b, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded models.OrderType
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if decoded.ClientOrderNumber != orig.ClientOrderNumber {
		t.Errorf("round-trip: got %q, want %q", decoded.ClientOrderNumber, orig.ClientOrderNumber)
	}
}
