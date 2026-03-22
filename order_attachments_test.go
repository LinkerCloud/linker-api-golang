package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestOrderAttachmentsService_Upload(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/orders/order-1/orderattachments")
		w.WriteHeader(http.StatusCreated)
	})

	if err := c.OrderAttachments.Upload(context.Background(), "order-1", models.AttachmentType{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrderAttachmentsService_GetByOrderNumber(t *testing.T) {
	want := []models.Attachment{{ID: "att-1"}}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orders/ORD-001/orderattachment")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.OrderAttachments.GetByOrderNumber(context.Background(), "ORD-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 {
		t.Errorf("len: got %d, want 1", len(got))
	}
}

func TestOrderAttachmentsService_GetByID(t *testing.T) {
	want := models.OrderAttachment{ID: "att-99"}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/orders/attachment/att-99")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	got, err := c.OrderAttachments.GetByID(context.Background(), "att-99")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != want.ID {
		t.Errorf("ID: got %q, want %q", got.ID, want.ID)
	}
}

func TestOrderAttachmentsService_PrintLabels(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodPost)
		testutil.AssertPath(t, r, "/public-api/v1/orderattachments/prints")
		w.WriteHeader(http.StatusOK)
	})

	if err := c.OrderAttachments.PrintLabels(context.Background(), models.AttachmentPrintType{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
