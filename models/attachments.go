package models

// Attachment is a Linker Cloud API resource.
type Attachment struct {
	ID string `json:"id,omitempty"`
	// Project ID
	DepotID *int64 `json:"depot_id,omitempty"`
	// Quantity of document pages
	PagesQuantity *int64 `json:"pages_quantity,omitempty"`
	// Attachment type
	Type string `json:"type,omitempty"`
	// Attachment reference
	Reference string `json:"reference,omitempty"`
	// Document number
	DocumentNumber string `json:"document_number,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	// Metadata of the attachment
	Metadata     []map[string]any  `json:"metadata,omitempty"`
	OriginalName string            `json:"original_name,omitempty"`
	CreatedBy    string            `json:"created_by,omitempty"`
	MetadataHash map[string]string `json:"metadata_hash,omitempty"`
}

// AttachmentPrintType is a Linker Cloud API resource.
type AttachmentPrintType struct {
	// Order number for identification
	OrderNumber string `json:"orderNumber"`
	// Attachment type
	Type string `json:"type"`
	// Packing station ID
	StationID string `json:"stationID"`
}

// AttachmentType is a Linker Cloud API resource.
type AttachmentType struct {
	// Attachment type
	Type string `json:"type"`
	// File encoded as base64
	Content string `json:"content"`
	// Document number, eg. invoice number
	DocumentNumber string `json:"documentNumber,omitempty"`
}
