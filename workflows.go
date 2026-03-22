package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// WorkflowsService handles communication with the Workflow endpoints of the
// Linker Cloud API.
type WorkflowsService struct {
	client *Client
}

// Get returns the workflow identified by id.
func (s *WorkflowsService) Get(ctx context.Context, id string) (*models.Workflow, error) {
	var result models.Workflow
	path := fmt.Sprintf("/public-api/v1/workflows/%s", url.PathEscape(id))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
