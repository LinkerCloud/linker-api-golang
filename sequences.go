package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// SequencesService handles communication with the Sequence endpoints of the
// Linker Cloud API.
type SequencesService struct {
	client *Client
}

// Get returns the next value of the sequence identified by name.
func (s *SequencesService) Get(ctx context.Context, name string) (int64, error) {
	var result int64
	path := fmt.Sprintf("/public-api/v1/sequences/%s", url.PathEscape(name))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return 0, err
	}
	return result, nil
}
