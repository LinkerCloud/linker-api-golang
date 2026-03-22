package models

// ParcelType is a Linker Cloud API resource.
type ParcelType struct {
	// ID of the package
	PackageID string `json:"packageId,omitempty"`
	// ID of the parcel
	ParcelID string `json:"parcelId,omitempty"`
	// Label in base64 format
	Label string `json:"label,omitempty"`
	// URL to track a package
	TrackingURL string `json:"trackingUrl,omitempty"`
	// Tracking number
	TrackingNumber string `json:"trackingNumber,omitempty"`
	// Name of the operator
	OperatorName string `json:"operatorName,omitempty"`
	// Wms ID
	WMSId string `json:"wmsId,omitempty"`
}

// ShipmentPackageType is a Linker Cloud API resource.
type ShipmentPackageType struct {
	// Weight of the package
	Weight *float64 `json:"weight,omitempty"`
	// Delivery package type
	Type string `json:"type,omitempty"`
	// Width of the package
	Width *float64 `json:"width,omitempty"`
	// Height of the package
	Height *float64 `json:"height,omitempty"`
	// Depth of the package
	Depth *float64 `json:"depth,omitempty"`
	// Wms ID
	WMSId string `json:"wmsId,omitempty"`
	// Products. Required, only if packOrder flag is set to true
	Items []*ItemType `json:"items"`
}

// CreateShipmentsRequest is the request body for POST /deliveries.
type CreateShipmentsRequest struct {
	// IDs is a list of order IDs to create shipments for (required).
	IDs []string `json:"ids"`
	// CreateAdditional creates an additional package per order when true.
	CreateAdditional bool `json:"createAdditional"`
}

// CreateShipmentsResponse is the response from POST /deliveries.
type CreateShipmentsResponse struct {
	LastID   string   `json:"last_id"`
	Packages []string `json:"packages"`
}

// CancelPackagesRequest is the request body for PATCH /deliveries/{id}.
type CancelPackagesRequest struct {
	// IDs is a list of package IDs to cancel.
	IDs []string `json:"ids,omitempty"`
	// PackIDs is an alternative list of package IDs to cancel.
	PackIDs []string `json:"packIds,omitempty"`
}

// ShipmentType is a Linker Cloud API resource.
type ShipmentType struct {
	// Order number
	OrderNumber string `json:"orderNumber"`
	// Collection of packages
	Packages []*ShipmentPackageType `json:"packages"`
	// Batch numbers with quantities
	BatchNumbers []*ItemBatchNumberType `json:"batchNumbers"`
	// Format of the label that endpoint will return
	LabelFormat string `json:"labelFormat,omitempty"`
	// Set this to true if you want to pack the order
	PackOrder *bool `json:"packOrder,omitempty"`
	// Set this to true if you want to pack the order
	MarkAsPacked bool `json:"markAsPacked"`
	// Set this to true if you want to create additional package
	CreateAdditional bool `json:"createAdditional"`
}
