package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// WebhooksService handles communication with the Integration Webhook endpoints
// of the Linker Cloud API.
type WebhooksService struct {
	client *Client
}

// Trigger fires the webhook for the given action and integration.
// The body structure is integration-specific — consult the documentation for
// your integration adapter. It typically contains order identification fields
// (external_id or client_order_number) and action-specific data.
// Supported actions: FULFILLED, CANCELLED, DELETED.
func (s *WebhooksService) Trigger(ctx context.Context, action, integrationName string, body any) error {
	path := fmt.Sprintf("/public-api/v1/integrationwebhooks/%s/webhooks/%s", url.PathEscape(action), url.PathEscape(integrationName))
	return s.client.doRequest(ctx, http.MethodPost, path, nil, body, nil)
}
