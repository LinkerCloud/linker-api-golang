package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestOrderTransitionsService_GetAvailable(t *testing.T) {
	want := models.OrderTransitionsResponse{
		Transitions: []*models.OrderTransitionItem{
			{Name: "ship", Froms: []string{"V"}, Tos: []string{"Y"}},
			{Name: "cancel", Froms: []string{"V"}, Tos: []string{"A"}},
		},
		AvailableTransitions: []string{"ship", "cancel"},
		Initial:              "N",
	}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orders/ord-1/transition")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.OrderTransitions.GetAvailable(context.Background(), "ord-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Transitions) != 2 {
		t.Errorf("len(Transitions): got %d, want 2", len(got.Transitions))
	}
	if got.Initial != "N" {
		t.Errorf("Initial: got %q, want %q", got.Initial, "N")
	}
}

func TestOrderTransitionsService_Apply(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/orders/ord-1/transitions/ship")
		w.WriteHeader(http.StatusOK)
	})

	if err := c.OrderTransitions.Apply(context.Background(), "ord-1", "ship"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
