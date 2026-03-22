package linkercloud

import (
	"context"
	"net/http"
	"testing"

	"github.com/linkercloud/linker-api-golang/internal/testutil"
	"github.com/linkercloud/linker-api-golang/models"
)

func TestWorkflowsService_Get(t *testing.T) {
	want := models.Workflow{}

	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.AssertMethod(t, r, http.MethodGet)
		testutil.AssertPath(t, r, "/public-api/v1/workflows/wf-1")
		testutil.RespondJSON(t, w, http.StatusOK, want)
	})

	_, err := c.Workflows.Get(context.Background(), "wf-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWorkflowsService_Get_notFound(t *testing.T) {
	c, _ := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		testutil.RespondError(w, http.StatusNotFound, "not found")
	})

	_, err := c.Workflows.Get(context.Background(), "nope")
	if !IsNotFound(err) {
		t.Fatalf("expected IsNotFound, got: %v", err)
	}
}
