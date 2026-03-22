package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
)

func TestWebhooksService_Trigger(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/integrationwebhooks/order_created/webhooks/shopify")
		w.WriteHeader(http.StatusOK)
	})

	if err := c.Webhooks.Trigger(context.Background(), "order_created", "shopify", nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWebhooksService_Trigger_withBody(t *testing.T) {
	payload := map[string]string{"orderId": "ord-1"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		var got map[string]string
		testutil.DecodeBody(t, r, &got)
		if got["orderId"] != "ord-1" {
			t.Errorf("orderId: got %q, want %q", got["orderId"], "ord-1")
		}
		w.WriteHeader(http.StatusOK)
	})

	if err := c.Webhooks.Trigger(context.Background(), "action", "integration", payload); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWebhooksService_Trigger_badRequest(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusBadRequest, "invalid payload")
	})

	err := c.Webhooks.Trigger(context.Background(), "a", "b", nil)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*APIError)
	if !ok || apiErr.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400 APIError, got: %v", err)
	}
}
