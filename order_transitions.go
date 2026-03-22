package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// OrderTransitionsService handles communication with the Order Transitions
// endpoints of the Linker Cloud API.
type OrderTransitionsService struct {
	client *Client
}

// GetAvailable returns the transition state machine data for the order
// identified by orderID. The response includes all defined transitions,
// the currently available ones, available places, and the initial place.
func (s *OrderTransitionsService) GetAvailable(ctx context.Context, orderID string) (*models.OrderTransitionsResponse, error) {
	var result models.OrderTransitionsResponse
	path := fmt.Sprintf("/public-api/v1/orders/%s/transition", url.PathEscape(orderID))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Apply applies the named transition to the order identified by orderID.
func (s *OrderTransitionsService) Apply(ctx context.Context, orderID, transition string) error {
	path := fmt.Sprintf("/public-api/v1/orders/%s/transitions/%s", url.PathEscape(orderID), url.PathEscape(transition))
	return s.client.doRequest(ctx, http.MethodPost, path, nil, nil, nil)
}
