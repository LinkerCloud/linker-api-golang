package linkercloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/linkercloud/linker-api-golang/models"
)

// OrderAttachmentsService handles communication with the Order Attachments
// endpoints of the Linker Cloud API.
type OrderAttachmentsService struct {
	client *Client
}

// Upload creates an attachment for the order identified by orderID.
func (s *OrderAttachmentsService) Upload(ctx context.Context, orderID string, attachment models.AttachmentType) error {
	path := fmt.Sprintf("/public-api/v1/orders/%s/orderattachments", url.PathEscape(orderID))
	return s.client.doRequest(ctx, http.MethodPost, path, nil, attachment, nil)
}

// PrintLabels triggers printing of attachments for an order.
func (s *OrderAttachmentsService) PrintLabels(ctx context.Context, req models.AttachmentPrintType) error {
	return s.client.doRequest(ctx, http.MethodPost, "/public-api/v1/orderattachments/prints", nil, req, nil)
}

// GetByOrderNumber returns all attachments for the order identified by orderNumber.
func (s *OrderAttachmentsService) GetByOrderNumber(ctx context.Context, orderNumber string) ([]models.Attachment, error) {
	var result []models.Attachment
	path := fmt.Sprintf("/public-api/v1/orders/%s/orderattachment", url.PathEscape(orderNumber))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetByID returns the attachment identified by attachmentID.
func (s *OrderAttachmentsService) GetByID(ctx context.Context, attachmentID string) (*models.OrderAttachment, error) {
	var result models.OrderAttachment
	path := fmt.Sprintf("/public-api/v1/orders/attachment/%s", url.PathEscape(attachmentID))
	if err := s.client.doRequest(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
