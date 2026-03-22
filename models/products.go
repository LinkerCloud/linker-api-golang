package models

// FulfillmentSite is a Linker Cloud API resource.
type FulfillmentSite struct {
	Destinations any `json:"destinations"`
}

// FulfillmentSiteMinimalStock defines the minimum stock threshold for a product
// at a specific fulfillment site.
type FulfillmentSiteMinimalStock struct {
	FulfillmentSiteName string `json:"fulfillmentSiteName,omitempty"`
	MinimalStock        *int64 `json:"minimalStock,omitempty"`
}

// FulfillmentSiteStock holds stock quantities for a product at a specific
// fulfillment site, including reservations and availability.
type FulfillmentSiteStock struct {
	FulfillmentSiteName string   `json:"fulfillmentSiteName,omitempty"`
	UpdatedAt           string   `json:"updatedAt,omitempty"`
	TotalQuantity       *int64   `json:"totalQuantity,omitempty"`
	Reserved            *int64   `json:"reserved,omitempty"`
	Available           *int64   `json:"available,omitempty"`
	AvailabilityLevel   string   `json:"availabilityLevel,omitempty"`
	SerialNumbers       []string `json:"serialNumbers,omitempty"`
	Location            string   `json:"location,omitempty"`
}

// Product is a Linker Cloud API resource.
type Product struct {
	ID                          string   `json:"id,omitempty"`
	UUID                        string   `json:"uuid,omitempty"`
	ExternalID                  *int64   `json:"external_id,omitempty"`
	VariantExternalID           *int64   `json:"variant_external_id,omitempty"`
	StockExternalID             *int64   `json:"stock_external_id,omitempty"`
	CategoryExternalID          *int64   `json:"category_external_id,omitempty"`
	SKU                         string   `json:"sku,omitempty"`
	ParentSKU                   string   `json:"parent_sku,omitempty"`
	DepotID                     *int64   `json:"depot_id,omitempty"`
	MinimalStock                *int64   `json:"minimal_stock,omitempty"`
	Name                        string   `json:"name,omitempty"`
	NameAliases                 []string `json:"name_aliases,omitempty"`
	Barcode                     string   `json:"barcode,omitempty"`
	AdditionalCodes             []any    `json:"additional_codes,omitempty"`
	Unit                        string   `json:"unit,omitempty"`
	Type                        string   `json:"type,omitempty"`
	NetWeight                   *float64 `json:"net_weight,omitempty"`
	Weight                      *float64 `json:"weight,omitempty"`
	BoxWeight                   *float64 `json:"box_weight,omitempty"`
	NetBoxWeight                *float64 `json:"net_box_weight,omitempty"`
	LayerWeight                 *float64 `json:"layer_weight,omitempty"`
	LayerNetWeight              *float64 `json:"layer_net_weight,omitempty"`
	PaletteWeight               *float64 `json:"palette_weight,omitempty"`
	PaletteNetWeight            *float64 `json:"palette_net_weight,omitempty"`
	WeightUnit                  string   `json:"weight_unit,omitempty"`
	Volume                      *float64 `json:"volume,omitempty"`
	VolumeUnit                  string   `json:"volume_unit,omitempty"`
	Length                      *float64 `json:"length,omitempty"`
	Width                       *float64 `json:"width,omitempty"`
	Depth                       *float64 `json:"depth,omitempty"`
	BoxLength                   *float64 `json:"box_length,omitempty"`
	BoxWidth                    *float64 `json:"box_width,omitempty"`
	BoxDepth                    *float64 `json:"box_depth,omitempty"`
	LayerLength                 *float64 `json:"layer_length,omitempty"`
	LayerWidth                  *float64 `json:"layer_width,omitempty"`
	LayerDepth                  *float64 `json:"layer_depth,omitempty"`
	PaletteLength               *float64 `json:"palette_length,omitempty"`
	PaletteWidth                *float64 `json:"palette_width,omitempty"`
	PaletteDepth                *float64 `json:"palette_depth,omitempty"`
	DimensionsUnit              string   `json:"dimensions_unit,omitempty"`
	Stock                       *int64   `json:"stock,omitempty"`
	IntegrationClientName       string   `json:"integration_client_name,omitempty"`
	ExternalProductCopies       any      `json:"external_product_copies,omitempty"`
	WMSExternalID               string   `json:"wms_external_id,omitempty"`
	ExportedToWMS               *bool    `json:"exported_to_wms,omitempty"`
	IgnoreInWMS                 *bool    `json:"ignore_in_w_m_s,omitempty"`
	IgnoreWhenPacking           *bool    `json:"ignore_when_packing,omitempty"`
	Category                    string   `json:"category,omitempty"`
	StorageUnits                []any    `json:"storage_units,omitempty"`
	StorageUnitWMSId            string   `json:"storage_unit_wms_id,omitempty"`
	MergedWithWMSStorageUnit    *bool    `json:"merged_with_wms_storage_unit,omitempty"`
	Storehouse                  string   `json:"storehouse,omitempty"`
	StorehouseLocation          string   `json:"storehouse_location,omitempty"`
	CreatedAt                   string   `json:"created_at,omitempty"`
	UpdatedAt                   string   `json:"updated_at,omitempty"`
	ExportedToWMSAt             string   `json:"exported_to_wms_at,omitempty"`
	AlwaysAskForSerialNumber    *bool    `json:"always_ask_for_serial_number,omitempty"`
	IsExpirable                 *bool    `json:"is_expirable,omitempty"`
	HasBatchNumber              *bool    `json:"has_batch_number,omitempty"`
	IsBio                       *bool    `json:"is_bio,omitempty"`
	IsFood                      *bool    `json:"is_food,omitempty"`
	IsInsert                    *bool    `json:"is_insert,omitempty"`
	IsFragile                   *bool    `json:"is_fragile,omitempty"`
	HarmonizedSystemCode        string   `json:"harmonized_system_code,omitempty"`
	CountryOfOrigin             string   `json:"country_of_origin,omitempty"`
	ProvinceOfOrigin            string   `json:"province_of_origin,omitempty"`
	FulfillmentSiteStocks       any      `json:"fulfillment_site_stocks,omitempty"`
	FulfillmentSiteMinimalStocks []FulfillmentSiteMinimalStock `json:"fulfillment_site_minimal_stocks,omitempty"`
	CustomProperties            []any    `json:"custom_properties,omitempty"`
	MapIntegrationExternalID    []any    `json:"map_integration_external_id,omitempty"`
}

// ProductType is the V1 request body for creating/updating a product
// via POST/PUT /public-api/v1/products.
// See [ProductType2] for the V2 variant with additional fields.
type ProductType struct {
	// Product name (required)
	Name string `json:"name,omitempty"`
	// SKU product code (required)
	SKU string `json:"sku,omitempty"`
	// Barcode data (required)
	Barcode string `json:"barcode,omitempty"`
	// Warehouse ID (required)
	DepotID string `json:"depotId,omitempty"`
	// Category where product will be assigned
	Category string `json:"category,omitempty"`
	// Item unit
	Unit string `json:"unit,omitempty"`
	// Product net weight
	NetWeight string `json:"netWeight,omitempty"`
	// Product net weight
	Weight string `json:"weight,omitempty"`
	// Product weight unit
	WeightUnit string `json:"weightUnit,omitempty"`
	// Product length
	Length string `json:"length,omitempty"`
	// Product width
	Width string `json:"width,omitempty"`
	// Product depth
	Depth string `json:"depth,omitempty"`
	// Box length
	BoxLength string `json:"boxLength,omitempty"`
	// Box Width
	BoxWidth string `json:"boxWidth,omitempty"`
	// Box depth
	BoxDepth string `json:"boxDepth,omitempty"`
	// Box weight
	BoxWeight string `json:"boxWeight,omitempty"`
	// Box net weight
	NetBoxWeight string `json:"netBoxWeight,omitempty"`
	// Layer length
	LayerLength string `json:"layerLength,omitempty"`
	// Layer width
	LayerWidth string `json:"layerWidth,omitempty"`
	// Layer depth
	LayerDepth string `json:"layerDepth,omitempty"`
	// Layer weight
	LayerWeight string `json:"layerWeight,omitempty"`
	// Layer net weight
	LayerNetWeight string `json:"layerNetWeight,omitempty"`
	// Product volume
	Volume string `json:"volume,omitempty"`
	// Product volume unit
	VolumeUnit string `json:"volumeUnit,omitempty"`
	// Storage units
	StorageUnits []*UnstructuredType `json:"storageUnits"`
	// Aliases for product name
	NameAliases []*UnstructuredType `json:"nameAliases"`
	// Images
	Images []string `json:"images"`
	// Length, Width, Depth unit
	DimensionsUnit string `json:"dimensionsUnit,omitempty"`
	// Product ID in source system
	ExternalID string `json:"externalId,omitempty"`
	// Product ID from WMS
	WMSExternalID string `json:"wms_external_id,omitempty"`
	// Additional codes
	AdditionalCodes []*UnstructuredType `json:"additionalCodes"`
	// Product localization in storehouse
	StorehouseLocation string `json:"storehouse_location,omitempty"`
	// Harmonized System Code
	HarmonizedSystemCode string `json:"harmonized_system_code,omitempty"`
	// Country of Origin, specified as SO 3166-2, eg. DE, US, PL.
	CountryOfOrigin string `json:"country_of_origin,omitempty"`
	// Province of Origin
	ProvinceOfOrigin string `json:"province_of_origin,omitempty"`
	// Should linker ignore this product when wms operations?
	IgnoreInWMS bool `json:"ignore_in_wms"`
	// When selected then product does not interact with WMS.
	IgnoreWhenPacking bool `json:"ignore_when_packing"`
	// Should linker always ask for serial number?
	AlwaysAskForSerialNumber bool `json:"always_ask_for_serial_number"`
	// Has the product batch number?
	HasBatchNumber bool `json:"has_batch_number"`
	// Is the product expirable?
	IsExpirable bool `json:"is_expirable"`
	// Is the product bio?
	IsBio bool `json:"is_bio"`
	// Is the product food?
	IsFood bool `json:"is_food"`
	// Is the product insert?
	IsInsert bool `json:"is_insert"`
	// Is the product fragile?
	IsFragile bool `json:"is_fragile"`
	// Minimal stock
	MinimalStock string `json:"minimalStock,omitempty"`
}

// ProductType2 is the V2 request body for creating/updating a product
// via POST/PUT /public-api/v2/products. It includes additional fields
// (IntegrationClientName, VariantExternalID) not available in V1.
// See [ProductType] for the V1 variant.
type ProductType2 struct {
	// Product name (required)
	Name string `json:"name,omitempty"`
	// SKU product code (required)
	SKU string `json:"sku,omitempty"`
	// Barcode data (required)
	Barcode string `json:"barcode,omitempty"`
	// Warehouse ID (required)
	DepotID string `json:"depotId,omitempty"`
	// Category where product will be assigned
	Category string `json:"category,omitempty"`
	// Item unit
	Unit string `json:"unit,omitempty"`
	// Product net weight
	NetWeight string `json:"netWeight,omitempty"`
	// Product net weight
	Weight string `json:"weight,omitempty"`
	// Product weight unit
	WeightUnit string `json:"weightUnit,omitempty"`
	// Product length
	Length string `json:"length,omitempty"`
	// Product width
	Width string `json:"width,omitempty"`
	// Product depth
	Depth string `json:"depth,omitempty"`
	// Box length
	BoxLength string `json:"boxLength,omitempty"`
	// Box Width
	BoxWidth string `json:"boxWidth,omitempty"`
	// Box depth
	BoxDepth string `json:"boxDepth,omitempty"`
	// Box weight
	BoxWeight string `json:"boxWeight,omitempty"`
	// Box net weight
	NetBoxWeight string `json:"netBoxWeight,omitempty"`
	// Layer length
	LayerLength string `json:"layerLength,omitempty"`
	// Layer width
	LayerWidth string `json:"layerWidth,omitempty"`
	// Layer depth
	LayerDepth string `json:"layerDepth,omitempty"`
	// Layer weight
	LayerWeight string `json:"layerWeight,omitempty"`
	// Layer net weight
	LayerNetWeight string `json:"layerNetWeight,omitempty"`
	// Product volume
	Volume string `json:"volume,omitempty"`
	// Product volume unit
	VolumeUnit string `json:"volumeUnit,omitempty"`
	// Storage units
	StorageUnits []*UnstructuredType `json:"storageUnits"`
	// Aliases for product name
	NameAliases []*UnstructuredType `json:"nameAliases"`
	// Images
	Images []string `json:"images"`
	// Length, Width, Depth unit
	DimensionsUnit string `json:"dimensionsUnit,omitempty"`
	// Product ID in source system
	ExternalID string `json:"externalId,omitempty"`
	// Additional codes
	AdditionalCodes []*UnstructuredType `json:"additionalCodes"`
	// Minimal stock
	MinimalStock string `json:"minimalStock,omitempty"`
	// Product ID from WMS
	WMSExternalID string `json:"wmsExternalID,omitempty"`
	// Product localization in storehouse
	StorehouseLocation string `json:"storehouseLocation,omitempty"`
	// Harmonized System Code
	HarmonizedSystemCode string `json:"harmonizedSystemCode,omitempty"`
	// Country of Origin, specified as SO 3166-2, eg. DE, US, PL.
	CountryOfOrigin string `json:"countryOfOrigin,omitempty"`
	// Province of Origin
	ProvinceOfOrigin string `json:"provinceOfOrigin,omitempty"`
	// Should linker ignore this product when wms operations?
	IgnoreInWMS bool `json:"ignoreInWms"`
	// When selected then product does not interact with WMS.
	IgnoreWhenPacking bool `json:"ignoreWhenPacking"`
	// Should linker always ask for serial number?
	AlwaysAskForSerialNumber bool `json:"alwaysAskForSerialNumber"`
	// Has the product batch number?
	HasBatchNumber bool `json:"hasBatchNumber"`
	// Is the product expirable?
	IsExpirable bool `json:"isExpirable"`
	// Is the product bio?
	IsBio bool `json:"isBio"`
	// Is the product food?
	IsFood bool `json:"isFood"`
	// Is the product insert?
	IsInsert bool `json:"isInsert"`
	// Is the product fragile?
	IsFragile bool `json:"isFragile"`
	// Integration name (filled if product comes from integration)
	IntegrationClientName string `json:"integrationClientName,omitempty"`
	// Product variant ID in source system
	VariantExternalID string `json:"variantExternalID,omitempty"`
}
