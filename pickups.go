package linkercloud

import (
	"context"
	"net/http"

	"github.com/linkercloud/linker-api-golang/models"
)

// PickupsService handles communication with the Pickup endpoints of the
// Linker Cloud API.
type PickupsService struct {
	client *Client
}

// Create schedules a new pickup.
func (s *PickupsService) Create(ctx context.Context, req models.PickupRequest) error {
	return s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/pickups", nil, req, nil)
}
