package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// PackingService handles communication with the Packing endpoints of the
// Linker Cloud API.
type PackingService struct {
	client *Client
}

// GetContainers returns the packing containers for the order identified by
// orderID.
func (s *PackingService) GetContainers(ctx context.Context, orderID string) (*models.PackingContainer, error) {
	var result models.PackingContainer
	path := fmt.Sprintf("/public-api/v1/orders/%s/packingcontainers", url.PathEscape(orderID))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMaterials returns the packing materials for the order identified by
// orderID.
func (s *PackingService) GetMaterials(ctx context.Context, orderID string) (*models.OrderMaterial, error) {
	var result models.OrderMaterial
	path := fmt.Sprintf("/public-api/v1/orders/%s/packingmaterials", url.PathEscape(orderID))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
