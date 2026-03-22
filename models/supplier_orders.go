package models

// InternalSupplierOrderState is a Linker Cloud API resource.
type InternalSupplierOrderState struct {
	Value string `json:"value,omitempty"`
	Cache any    `json:"cache,omitempty"`
	Name  string `json:"name,omitempty"`
}

// Supplier is a Linker Cloud API resource.
type Supplier struct {
	ID         string `json:"id,omitempty"`
	ExternalID *int64 `json:"external_id,omitempty"`
	Code       string `json:"code,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Country    string `json:"country,omitempty"`
	FullName   string `json:"full_name,omitempty"`
	Nip        string `json:"nip,omitempty"`
	Regon      string `json:"regon,omitempty"`
	PostCode   string `json:"post_code,omitempty"`
	City       string `json:"city,omitempty"`
	Street     string `json:"street,omitempty"`
	DepotID    *int64 `json:"depot_id,omitempty"`
	WMSId      string `json:"wms_id,omitempty"`
}

// SupplierOrder is a Linker Cloud API resource.
type SupplierOrder struct {
	ID                          string                      `json:"id,omitempty"`
	UUID                        string                      `json:"uuid,omitempty"`
	Supplier                    *int64                      `json:"supplier,omitempty"`
	SupplierID                  string                      `json:"supplier_id,omitempty"`
	Items                       []*SupplierOrderItem        `json:"items,omitempty"`
	ExternalID                  string                      `json:"external_id,omitempty"`
	Number                      string                      `json:"number,omitempty"`
	DepotID                     *int64                      `json:"depot_id,omitempty"`
	OrderDate                   string                      `json:"order_date,omitempty"`
	ExecutionDate               string                      `json:"execution_date,omitempty"`
	Installment                 *float64                    `json:"installment,omitempty"`
	Priority                    *int64                      `json:"priority,omitempty"`
	ClientOrderNumber           string                      `json:"client_order_number,omitempty"`
	Type                        *int64                      `json:"type,omitempty"`
	AutoNumber                  *int64                      `json:"auto_number,omitempty"`
	PriceGross                  *float64                    `json:"price_gross,omitempty"`
	PriceNet                    *float64                    `json:"price_net,omitempty"`
	PriceGrossInForeignCurrency *float64                    `json:"price_gross_in_foreign_currency,omitempty"`
	PriceNetInForeignCurrency   *float64                    `json:"price_net_in_foreign_currency,omitempty"`
	ExchangeRate                *float64                    `json:"exchange_rate,omitempty"`
	CurrencySymbol              string                      `json:"currency_symbol,omitempty"`
	DocumentInForeignCurrency   *bool                       `json:"document_in_foreign_currency,omitempty"`
	ExchangeRateDate            string                      `json:"exchange_rate_date,omitempty"`
	OrderStatus                 string                      `json:"order_status,omitempty"`
	InternalState               *InternalSupplierOrderState `json:"internal_state,omitempty"`
	Comments                    string                      `json:"comments,omitempty"`
	Discount                    *float64                    `json:"discount,omitempty"`
	ExportedToWMS               *bool                       `json:"exported_to_w_m_s,omitempty"`
	ExportedAt                  string                      `json:"exported_at,omitempty"`
	SupplierObject              *Supplier                   `json:"supplierObject,omitempty"`
	IsReturn                    *bool                       `json:"is_return,omitempty"`
	IsConsistent                *bool                       `json:"is_consistent,omitempty"`
	ReturnNumber                string                      `json:"returnNumber,omitempty"`
	OrderNumber                 string                      `json:"orderNumber,omitempty"`
	StatusHistory               []*StatusHistoryEntry       `json:"status_history,omitempty"`
	WMSId                       string                      `json:"wms_id,omitempty"`
	NumberOfParcels             *int64                      `json:"number_of_parcels,omitempty"`
	NumberOfPallets             *int64                      `json:"number_of_pallets,omitempty"`
	NumberOfContainers          *int64                      `json:"number_of_containers,omitempty"`
	CreatedBy                   string                      `json:"created_by,omitempty"`
	CreatedAt                   string                      `json:"createdAt,omitempty"`
	UpdatedAt                   string                      `json:"updatedAt,omitempty"`
	IntegrationClientName       string                      `json:"integration_client_name,omitempty"`
	RejectionReason             string                      `json:"rejection_reason,omitempty"`
	IsVerified                  *bool                       `json:"is_verified,omitempty"`
	CustomProperties            []any                       `json:"custom_properties,omitempty"`
	AssignedFulfillmentSite     string                      `json:"assigned_fulfillment_site,omitempty"`
	IsPartial                   *bool                       `json:"is_partial,omitempty"`
}

// SupplierOrderItem is a Linker Cloud API resource.
type SupplierOrderItem struct {
	ID                          string         `json:"id,omitempty"`
	Order                       *SupplierOrder `json:"order,omitempty"`
	ExternalID                  string         `json:"external_id,omitempty"`
	ProductExternalID           *int64         `json:"product_external_id,omitempty"`
	VariantExternalID           string         `json:"variant_external_id,omitempty"`
	SKU                         string         `json:"sku,omitempty"`
	EAN                         string         `json:"ean,omitempty"`
	VATCode                     string         `json:"vat_code,omitempty"`
	Ordered                     *int64         `json:"ordered,omitempty"`
	Realized                    *int64         `json:"realized,omitempty"`
	ToBook                      *int64         `json:"to_book,omitempty"`
	ToAchieve                   *int64         `json:"to_achieve,omitempty"`
	PriceGross                  *float64       `json:"price_gross,omitempty"`
	PriceNet                    *float64       `json:"price_net,omitempty"`
	PriceGrossInForeignCurrency *float64       `json:"price_gross_in_foreign_currency,omitempty"`
	PriceNetInForeignCurrency   *float64       `json:"price_net_in_foreign_currency,omitempty"`
	Converter                   *float64       `json:"converter,omitempty"`
	Unit                        string         `json:"unit,omitempty"`
	Margin                      *float64       `json:"margin,omitempty"`
	Description                 string         `json:"description,omitempty"`
	Weight                      *float64       `json:"weight,omitempty"`
	Width                       *float64       `json:"width,omitempty"`
	Length                      *float64       `json:"length,omitempty"`
	Depth                       *float64       `json:"depth,omitempty"`
	Category                    string         `json:"category,omitempty"`
	CreatedAt                   string         `json:"created_at,omitempty"`
	UpdatedAt                   string         `json:"updated_at,omitempty"`
	LotNumber                   string         `json:"lot_number,omitempty"`
	ExpirationDate              string         `json:"expiration_date,omitempty"`
	NetPurchasePrice            *float64       `json:"net_purchase_price,omitempty"`
	GrossPurchasePrice          *float64       `json:"gross_purchase_price,omitempty"`
	Sources                     any            `json:"sources,omitempty"`
	ShippedQuantity             *int64         `json:"shipped_quantity,omitempty"`
	DamagedQuantity             *int64         `json:"damaged_quantity,omitempty"`
	SingleBoxQuantity           *int64         `json:"single_box_quantity,omitempty"`
	BoxNumber                   string         `json:"box_number,omitempty"`
	IsExpirable                 *bool          `json:"is_expirable,omitempty"`
	HasBatchNumber              *bool          `json:"has_batch_number,omitempty"`
	IsBio                       *bool          `json:"is_bio,omitempty"`
	IsFood                      *bool          `json:"is_food,omitempty"`
	IsInsert                    *bool          `json:"is_insert,omitempty"`
	IsFragile                   *bool          `json:"is_fragile,omitempty"`
	ProductType                 string         `json:"product_type,omitempty"`
}

// SupplierOrderItemType is a Linker Cloud API resource.
type SupplierOrderItemType struct {
	ID                 string  `json:"id,omitempty"`
	ExternalID         string  `json:"externalId,omitempty"`
	ProductExternalID  string  `json:"productExternalID,omitempty"`
	VariantExternalID  string  `json:"variantExternalID,omitempty"`
	VATCode            string  `json:"vatCode,omitempty"`
	Ordered            int64   `json:"ordered"`
	PriceGross         *float64 `json:"priceGross,omitempty"`
	PriceNet           *float64 `json:"priceNet,omitempty"`
	Converter          string  `json:"converter,omitempty"`
	Unit               string  `json:"unit,omitempty"`
	Margin             *float64 `json:"margin,omitempty"`
	Description        string  `json:"description,omitempty"`
	SKU                string  `json:"sku,omitempty"`
	EAN                string  `json:"ean,omitempty"`
	Weight             string  `json:"weight,omitempty"`
	Width              string  `json:"width,omitempty"`
	Length             string  `json:"length,omitempty"`
	Depth              string  `json:"depth,omitempty"`
	Category           string  `json:"category,omitempty"`
	ExpirationDate     string  `json:"expirationDate,omitempty"`
	NetPurchasePrice   string  `json:"netPurchasePrice,omitempty"`
	GrossPurchasePrice string  `json:"grossPurchasePrice,omitempty"`
	LotNumber          string  `json:"lotNumber,omitempty"`
	DamagedQuantity    string  `json:"damagedQuantity,omitempty"`
	ShippedQuantity    string  `json:"shippedQuantity,omitempty"`
	SingleBoxQuantity  int64   `json:"singleBoxQuantity"`
	BoxNumber          string  `json:"boxNumber,omitempty"`
	IsExpirable        string  `json:"is_expirable,omitempty"`
	HasBatchNumber     string  `json:"has_batch_number,omitempty"`
	IsBio              string  `json:"isBio,omitempty"`
	IsFood             string  `json:"isFood,omitempty"`
	IsInsert           string  `json:"isInsert,omitempty"`
	IsFragile          string  `json:"isFragile,omitempty"`
}

// SupplierOrderType is a Linker Cloud API resource.
type SupplierOrderType struct {
	DepotID *int64 `json:"depotId,omitempty"`
	// Order date, Y-m-d H:i:s
	OrderDate string `json:"orderDate"`
	// Date when order is expected by warehouse
	ExecutionDate string `json:"executionDate"`
	Installment   string `json:"installment,omitempty"`
	Priority      string `json:"priority,omitempty"`
	AutoNumber    string `json:"autoNumber,omitempty"`
	// Order total gross price
	PriceGross *Money `json:"priceGross"`
	// Order total net price
	PriceNet *Money `json:"priceNet"`
	// Order currency specified as ISO 4217, eg. USD, EUR, PLN.
	CurrencySymbol string `json:"currencySymbol,omitempty"`
	// Order comments
	Comments string `json:"comments,omitempty"`
	// Linker Supplier ID if existing
	Supplier       int64         `json:"supplier"`
	SupplierObject *SupplierType `json:"supplierObject"`
	// Is Inbound Order based on return?
	IsReturn *bool `json:"isReturn,omitempty"`
	// Order number in source system
	ClientOrderNumber string `json:"clientOrderNumber,omitempty"`
	// Expected quantity of pallets
	NumberOfPallets string `json:"numberOfPallets,omitempty"`
	// Expected quantity of parcels
	NumberOfParcels string `json:"numberOfParcels,omitempty"`
	// Expected quantity of containers
	NumberOfContainers string `json:"numberOfContainers,omitempty"`
	// Order items
	Items []*SupplierOrderItemType `json:"items"`
	// Custom properties
	CustomProperties        []*UnstructuredType `json:"customProperties"`
	IsConsistent            string              `json:"isConsistent,omitempty"`
	IsVerified              string              `json:"isVerified,omitempty"`
	// Fulfillment Site ID
	AssignedFulfillmentSite string `json:"assignedFulfillmentSite,omitempty"`
	// Supplier ID
	SupplierID string `json:"supplierId,omitempty"`
}

// SupplierType is a Linker Cloud API resource.
type SupplierType struct {
	// Supplier ID in external system
	ExternalID string `json:"externalId,omitempty"`
	// Supplier code
	Code string `json:"code,omitempty"`
	// Supplier name
	Name string `json:"name,omitempty"`
	// Supplier contact email
	Email string `json:"email,omitempty"`
	// Supplier full name
	FullName string `json:"fullName,omitempty"`
	// Supplier tax ID
	Nip string `json:"nip,omitempty"`
	// Supplier post code
	PostCode string `json:"postCode,omitempty"`
	// Supplier city
	City string `json:"city,omitempty"`
	// Supplier street
	Street string `json:"street,omitempty"`
	// Country code
	Country string `json:"country"`
}
