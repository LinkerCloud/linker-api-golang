package models

// Container is a Linker Cloud API resource.
type Container struct {
	Products []*ContainerProduct `json:"products,omitempty"`
}

// ContainerProduct is a product within a packing [Container].
type ContainerProduct struct {
	SKU      string `json:"sku,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
}

// PackingContainer represents a packing container from the WMS.
// Only the most commonly needed fields are mapped; the full server
// entity has 30+ fields.
type PackingContainer struct {
	ID              string   `json:"id,omitempty"`
	UUID            string   `json:"uuid,omitempty"`
	Reference       string   `json:"reference,omitempty"`
	Number          string   `json:"number,omitempty"`
	Status          string   `json:"status,omitempty"`
	TrackingNumber  string   `json:"trackingNumber,omitempty"`
	PackageOperator string   `json:"packageOperator,omitempty"`
	PackageID       string   `json:"packageId,omitempty"`
	PackageType     string   `json:"packageType,omitempty"`
	Weight          *float64 `json:"weight,omitempty"`
	Width           *float64 `json:"width,omitempty"`
	Height          *float64 `json:"height,omitempty"`
	Depth           *float64 `json:"depth,omitempty"`
	DepotID         *int64   `json:"depotId,omitempty"`
}

// PickupRequest is the request body for POST /pickups.
type PickupRequest struct {
	// PickupDate is the requested pickup date in YYYY-MM-DD format (required).
	PickupDate string `json:"pickupDate"`
	// Provider is the delivery provider name, e.g. "POCZTA_POLSKA" (required).
	Provider string `json:"provider"`
	// Notify sends an email report to configured recipients when true.
	Notify bool `json:"notify"`
	// ForcePickup forces report generation when true.
	ForcePickup bool `json:"forcePickup"`
}

// PackingSuggestion is a Linker Cloud API resource.
type PackingSuggestion struct {
	Reference   string       `json:"reference,omitempty"`
	Status      string       `json:"status,omitempty"`
	RequestedAt string       `json:"requested_at,omitempty"`
	FulfilledAt string       `json:"fulfilled_at,omitempty"`
	FulfilledBy string       `json:"fulfilled_by,omitempty"`
	Containers  []*Container `json:"containers,omitempty"`
}
