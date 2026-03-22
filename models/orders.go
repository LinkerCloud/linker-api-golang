package models

// AdditionalService represents an add-on service attached to an order
// (e.g. insurance, COD handling, special packaging).
type AdditionalService struct {
	Quantity       *int64   `json:"quantity,omitempty"`
	ExternalID     string   `json:"externalId,omitempty"`
	ServiceCode    string   `json:"serviceCode,omitempty"`
	SKU            string   `json:"sku,omitempty"`
	Description    string   `json:"description,omitempty"`
	CurrencyCode   string   `json:"currencyCode,omitempty"`
	Unit           string   `json:"unit,omitempty"`
	TotalNetAmount *float64 `json:"totalNetAmount,omitempty"`
	TotalAmount    *float64 `json:"totalAmount,omitempty"`
	TotalTaxAmount *float64 `json:"totalTaxAmount,omitempty"`
	UnitNetAmount  *float64 `json:"unitNetAmount,omitempty"`
	UnitAmount     *float64 `json:"unitAmount,omitempty"`
	UnitTaxAmount  *float64 `json:"unitTaxAmount,omitempty"`
}

// Customer is the customer object embedded in [Order] responses (camelCase JSON tags).
// See [Customer2] for the snake_case variant returned by write operations.
type Customer struct {
	ID         *int64 `json:"id,omitempty"`
	ExternalID *int64 `json:"externalId,omitempty"`
	DepotID    *int64 `json:"depotId,omitempty"`
	Code       string `json:"code,omitempty"`
	Name       string `json:"name,omitempty"`
	FullName   string `json:"fullName,omitempty"`
	Nip        string `json:"nip,omitempty"`
	Regon      string `json:"regon,omitempty"`
	PostCode   string `json:"postCode,omitempty"`
	City       string `json:"city,omitempty"`
	Street     string `json:"street,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
}

// Customer2 is the customer object embedded in [Order2] responses (snake_case JSON tags).
// See [Customer] for the camelCase variant returned by read operations.
type Customer2 struct {
	ID         string   `json:"id,omitempty"`
	ExternalID *int64   `json:"external_id,omitempty"`
	DepotID    *int64   `json:"depot_id,omitempty"`
	Code       string   `json:"code,omitempty"`
	Name       string   `json:"name,omitempty"`
	FullName   string   `json:"full_name,omitempty"`
	Nip        string   `json:"nip,omitempty"`
	Regon      string   `json:"regon,omitempty"`
	PostCode   string   `json:"post_code,omitempty"`
	City       string   `json:"city,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	Street     string   `json:"street,omitempty"`
	Email      string   `json:"email,omitempty"`
	Phone      string   `json:"phone,omitempty"`
	Orders     []*Order2 `json:"orders,omitempty"`
	CreatedAt  string   `json:"created_at,omitempty"`
	UpdatedAt  string   `json:"updated_at,omitempty"`
	WMSId      string   `json:"wms_id,omitempty"`
}

// DeliveryStatus is a Linker Cloud API resource.
type DeliveryStatus struct {
	// External package ID
	PackageID string `json:"packageId,omitempty"`
	// External package status code
	ExternalStatusCode string `json:"externalStatusCode,omitempty"`
	// External package status name
	ExternalStatusName string `json:"externalStatusName,omitempty"`
	// Is waiting for pickup?
	WaitingForPickup *bool `json:"waitingForPickup,omitempty"`
	// Is in delivery?
	InDelivery *bool `json:"inDelivery,omitempty"`
	// Is delivered?
	Delivered *bool `json:"delivered,omitempty"`
	// Is delivery cancelled?
	Cancelled *bool `json:"cancelled,omitempty"`
	// City name, where package atually could be find
	City string `json:"city,omitempty"`
	// Country name, where package atually could be find
	Country string `json:"country,omitempty"`
}

// InternalOrderState is a Linker Cloud API resource.
type InternalOrderState struct {
	Value string `json:"value,omitempty"`
	Cache any    `json:"cache,omitempty"`
	Name  string `json:"name,omitempty"`
}

// ItemBatchNumberType is a Linker Cloud API resource.
type ItemBatchNumberType struct {
	// SKU
	SKU string `json:"sku,omitempty"`
	// Quantity of the products
	Quantity *int64 `json:"quantity,omitempty"`
	// Packed batch number
	PackedBatchNumber string `json:"packedBatchNumber,omitempty"`
	// Requested batch number
	RequestedBatchNumber string `json:"requestedBatchNumber,omitempty"`
}

// ItemType is a Linker Cloud API resource.
type ItemType struct {
	// SKU of the product
	SKU string `json:"sku,omitempty"`
	// Quantity of the products
	Quantity *int64 `json:"quantity,omitempty"`
	// Batch numbers with quantities
	BatchNumbers []*ItemBatchNumberType `json:"batchNumbers"`
	// Serial numbers
	SerialNumbers []*UnstructuredType `json:"serial_numbers"`
}

// Order is the response type for GET /public-api/v1/orders (list) and
// GET /public-api/v1/orders/{id} (get). Field names use camelCase JSON tags
// matching the V1 list/get response format.
// See [Order2] for the snake_case variant returned by write operations.
type Order struct {
	ID                           string               `json:"id,omitempty"`
	UUID                         string               `json:"uuid,omitempty"`
	ExternalID                   string               `json:"externalId,omitempty"`
	AdditionalOrderNumber        string               `json:"additionalOrderNumber,omitempty"`
	DepotID                      *int64               `json:"depotId,omitempty"`
	SrcClientID                  string               `json:"srcClientID,omitempty"`
	Number                       string               `json:"number,omitempty"`
	InvoiceNumber                string               `json:"invoiceNumber,omitempty"`
	InvoiceID                    string               `json:"invoiceId,omitempty"`
	ReceiptFiscalNumber          string               `json:"receiptFiscalNumber,omitempty"`
	ReceiptFiscalID              string               `json:"receiptFiscalID,omitempty"`
	GoodsDispatchNumber          string               `json:"goodsDispatchNumber,omitempty"`
	GoodsDispatchID              string               `json:"goodsDispatchID,omitempty"`
	ClientOrderNumber            string               `json:"clientOrderNumber,omitempty"`
	Customer                     *int64               `json:"customer,omitempty"`
	OrderDate                    string               `json:"orderDate,omitempty"`
	ExecutionDate                string               `json:"executionDate,omitempty"`
	ImportDate                   string               `json:"importDate,omitempty"`
	PriceGross                   *float64             `json:"priceGross,omitempty"`
	PriceNet                     *float64             `json:"priceNet,omitempty"`
	ShipmentPrice                *float64             `json:"shipmentPrice,omitempty"`
	CurrencySymbol               string               `json:"currencySymbol,omitempty"`
	OrderStatus                  string               `json:"orderStatus,omitempty"`
	Comments                     string               `json:"comments,omitempty"`
	// Order internal comments
	InternalComments             string               `json:"internalComments,omitempty"`
	Discount                     *float64             `json:"discount,omitempty"`
	DiscountTax                  *float64             `json:"discountTax,omitempty"`
	Carrier                      string               `json:"carrier,omitempty"`
	ExternalCarrier              string               `json:"external_carrier,omitempty"`
	// CarrierMappingHistory is the carrier mapping audit trail (array).
	CarrierMappingHistory        any                  `json:"carrier_mapping_history,omitempty"`
	DeliveryCountry              string               `json:"deliveryCountry,omitempty"`
	DeliveryPostCode             string               `json:"deliveryPostCode,omitempty"`
	DeliveryCity                 string               `json:"deliveryCity,omitempty"`
	DeliveryStreet               string               `json:"deliveryStreet,omitempty"`
	DeliveryCompany              string               `json:"deliveryCompany,omitempty"`
	DeliveryRecipient            string               `json:"deliveryRecipient,omitempty"`
	DeliveryRecipientLogin       string               `json:"deliveryRecipientLogin,omitempty"`
	DeliveryPhone                string               `json:"deliveryPhone,omitempty"`
	DeliveryEmail                string               `json:"deliveryEmail,omitempty"`
	DeliveryPointID              string               `json:"deliveryPointID,omitempty"`
	DeliveryPointName            string               `json:"deliveryPointName,omitempty"`
	DeliveryPointAddress         string               `json:"deliveryPointAddress,omitempty"`
	DeliveryPointPostcode        string               `json:"deliveryPointPostcode,omitempty"`
	DeliveryPointCity            string               `json:"deliveryPointCity,omitempty"`
	// DeliveryConfiguration is carrier-specific delivery config (object or array).
	DeliveryConfiguration        any                  `json:"deliveryConfiguration,omitempty"`
	PaymentMethod                string               `json:"paymentMethod,omitempty"`
	PaymentDate                  string               `json:"paymentDate,omitempty"`
	BillingFirstName             string               `json:"billingFirstName,omitempty"`
	BillingLastName              string               `json:"billingLastName,omitempty"`
	BillingCompany               string               `json:"billingCompany,omitempty"`
	BillingVATId                 string               `json:"billingVatID,omitempty"`
	BillingCity                  string               `json:"billingCity,omitempty"`
	BillingPostCode              string               `json:"billingPostCode,omitempty"`
	BillingStreet1               string               `json:"billingStreet1,omitempty"`
	BillingStreet2               string               `json:"billingStreet2,omitempty"`
	BillingState                 string               `json:"billingState,omitempty"`
	BillingCountry               string               `json:"billingCountry,omitempty"`
	BillingPhone                 string               `json:"billingPhone,omitempty"`
	BillingEmail                 string               `json:"billingEmail,omitempty"`
	BillingCountryCode           string               `json:"billingCountryCode,omitempty"`
	Cod                          *bool                `json:"cod,omitempty"`
	CodAmount                    *float64             `json:"codAmount,omitempty"`
	CodAccountNumber             string               `json:"codAccountNumber,omitempty"`
	Items                        []*OrderItem         `json:"items,omitempty"`
	Tags                         []string             `json:"tags,omitempty"`
	InvoicesNumbers              []any                `json:"invoicesNumbers,omitempty"`
	CustomerObject               *Customer            `json:"customerObject,omitempty"`
	ExternalDeliveryIds          []*ExternalDeliveryID `json:"externalDeliveryIds,omitempty"`
	ExternalDeliveryIdsCanceled  []*ExternalDeliveryID `json:"externalDeliveryIdsCanceled,omitempty"`
	IntegrationClientName        string               `json:"integrationClientName,omitempty"`
	WMSIntegrationName           string               `json:"wmsIntegrationName,omitempty"`
	AdditionalWMSIntegrationName string               `json:"additionalWmsIntegrationName,omitempty"`
	StatusHistory                []*StatusHistoryEntry `json:"statusHistory,omitempty"`
	InternalDeliveryMethod       string               `json:"internalDeliveryMethod,omitempty"`
	DeliveryMethodMap            string               `json:"deliveryMethodMap,omitempty"`
	ReturnDocumentsType          string               `json:"returnDocumentsType,omitempty"`
	DeliveryHour                 string               `json:"deliveryHour,omitempty"`
	TransactionID                string               `json:"transactionId,omitempty"`
	ExternalSellerID             string               `json:"externalSellerID,omitempty"`
	Origin                       string               `json:"origin,omitempty"`
	DeliveryDueDate              string               `json:"deliveryDueDate,omitempty"`
	DeliveryStatus               string               `json:"deliveryStatus,omitempty"`
	PackageType                  string               `json:"packageType,omitempty"`
	NumberOfPackages             *int64               `json:"numberOfPackages,omitempty"`
	PackedAt                     string               `json:"packedAt,omitempty"`
	PackedBy                     string               `json:"packedBy,omitempty"`
	PickedAt                     string               `json:"picked_at,omitempty"`
	PickedBy                     string               `json:"picked_by,omitempty"`
	PackedManually               *bool                `json:"packedManually,omitempty"`
	ContactPerson                string               `json:"contact_person,omitempty"`
	ContactPhone                 string               `json:"contact_phone,omitempty"`
	WMSOrderNumber               string               `json:"wms_order_number,omitempty"`
	PaymentTransactionID         string               `json:"payment_transaction_id,omitempty"`
	CourierInfo                  string               `json:"courier_info,omitempty"`
	ContractorID                 string               `json:"contractor_id,omitempty"`
	WMSId                        string               `json:"wms_id,omitempty"`
	OrderType                    string               `json:"order_type,omitempty"`
	DeliveryToPostBranch         *bool                `json:"delivery_to_post_branch,omitempty"`
	ContactFirstName             string               `json:"contact_first_name,omitempty"`
	ContactLastName              string               `json:"contact_last_name,omitempty"`
	ContactCity                  string               `json:"contact_city,omitempty"`
	ContactStreet                string               `json:"contact_street,omitempty"`
	ContactBuildingNumber        string               `json:"contact_building_number,omitempty"`
	ContactFlatNumber            string               `json:"contact_flat_number,omitempty"`
	ContactPostCode              string               `json:"contact_post_code,omitempty"`
	ContactPostOffice            string               `json:"contact_post_office,omitempty"`
	ContactCountryCode           string               `json:"contact_country_code,omitempty"`
	ContactCountryName           string               `json:"contact_country_name,omitempty"`
	ContactEmail                 string               `json:"contact_email,omitempty"`
	ShipmentPriceNet             *float64             `json:"shipment_price_net,omitempty"`
	ShipmentPriceTax             *float64             `json:"shipment_price_tax,omitempty"`
	TaxAmount                    *float64             `json:"tax_amount,omitempty"`
	TimeToDeadline               *int64               `json:"time_to_deadline,omitempty"`
	ConfirmedInWMS               *bool                `json:"confirmed_in_wms,omitempty"`
	OrderItemsCount              *int64               `json:"order_items_count,omitempty"`
	AdditionalClientOrderNumber  string               `json:"additional_client_order_number,omitempty"`
	UpdatedAt                    string               `json:"updated_at,omitempty"`
	UpdatedBy                    string               `json:"updated_by,omitempty"`
	PackingStation               string               `json:"packing_station,omitempty"`
	PackingStartedAt             string               `json:"packing_started_at,omitempty"`
	ExecutionDueDate             string               `json:"executionDueDate,omitempty"`
	PaymentStatus                string               `json:"payment_status,omitempty"`
	// CustomProperties is instance-specific custom fields (array or object).
	CustomProperties             any                  `json:"custom_properties,omitempty"`
}

// Order2 is the response type for POST, PUT, and PATCH operations on
// /public-api/v1/orders. Field names use snake_case JSON tags matching
// the V1 write response format. The Linker API returns different JSON
// naming conventions for read vs write operations.
// See [Order] for the camelCase variant returned by read operations.
type Order2 struct {
	ExternalID                    *int64               `json:"external_id,omitempty"`
	TagID                         string               `json:"tag_id,omitempty"`
	AdditionalOrderNumber         string               `json:"additional_order_number,omitempty"`
	DepotID                       *int64               `json:"depot_id,omitempty"`
	SrcClientID                   string               `json:"src_client_id,omitempty"`
	Number                        string               `json:"number,omitempty"`
	InvoiceNumber                 string               `json:"invoice_number,omitempty"`
	InvoicesNumbers               []any                `json:"invoices_numbers,omitempty"`
	InvoiceID                     string               `json:"invoice_id,omitempty"`
	ReceiptFiscalNumber           string               `json:"receipt_fiscal_number,omitempty"`
	ReceiptFiscalID               string               `json:"receipt_fiscal_id,omitempty"`
	GoodsDispatchNumber           string               `json:"goods_dispatch_number,omitempty"`
	GoodsDispatchID               string               `json:"goods_dispatch_id,omitempty"`
	ClientOrderNumber             string               `json:"client_order_number,omitempty"`
	Customer                      *int64               `json:"customer,omitempty"`
	OrderDate                     string               `json:"order_date,omitempty"`
	ExecutionDate                 string               `json:"execution_date,omitempty"`
	ImportDate                    string               `json:"import_date,omitempty"`
	PriceGross                    *float64             `json:"price_gross,omitempty"`
	PriceNet                      *float64             `json:"price_net,omitempty"`
	TaxAmount                     *float64             `json:"tax_amount,omitempty"`
	ShipmentPrice                 *float64             `json:"shipment_price,omitempty"`
	ShipmentPriceNet              *float64             `json:"shipment_price_net,omitempty"`
	ShipmentPriceTax              *float64             `json:"shipment_price_tax,omitempty"`
	CurrencySymbol                string               `json:"currency_symbol,omitempty"`
	OrderStatus                   string               `json:"order_status,omitempty"`
	Comments                      string               `json:"comments,omitempty"`
	InternalComments              string               `json:"internal_comments,omitempty"`
	Discount                      *float64             `json:"discount,omitempty"`
	DiscountTax                   *float64             `json:"discount_tax,omitempty"`
	Carrier                       string               `json:"carrier,omitempty"`
	ExternalCarrier               string               `json:"external_carrier,omitempty"`
	// CarrierMappingHistory is the carrier mapping audit trail (array).
	CarrierMappingHistory         any                  `json:"carrier_mapping_history,omitempty"`
	DeliveryCountry               string               `json:"delivery_country,omitempty"`
	DeliveryPostCode              string               `json:"delivery_post_code,omitempty"`
	DeliveryCity                  string               `json:"delivery_city,omitempty"`
	DeliveryStreet                string               `json:"delivery_street,omitempty"`
	DeliveryCompany               string               `json:"delivery_company,omitempty"`
	DeliveryRecipient             string               `json:"delivery_recipient,omitempty"`
	DeliveryRecipientLogin        string               `json:"delivery_recipient_login,omitempty"`
	DeliveryPhone                 string               `json:"delivery_phone,omitempty"`
	DeliveryEmail                 string               `json:"delivery_email,omitempty"`
	DeliveryPointID               string               `json:"delivery_point_id,omitempty"`
	DeliveryPointName             string               `json:"delivery_point_name,omitempty"`
	DeliveryPointAddress          string               `json:"delivery_point_address,omitempty"`
	DeliveryPointPostcode         string               `json:"delivery_point_postcode,omitempty"`
	DeliveryPointCity             string               `json:"delivery_point_city,omitempty"`
	// DeliveryConfiguration is carrier-specific delivery config (object or array).
	DeliveryConfiguration         any                  `json:"delivery_configuration,omitempty"`
	PaymentMethod                 string               `json:"payment_method,omitempty"`
	PaymentDate                   string               `json:"payment_date,omitempty"`
	PaymentStatus                 string               `json:"payment_status,omitempty"`
	BillingFirstName              string               `json:"billing_first_name,omitempty"`
	BillingLastName               string               `json:"billing_last_name,omitempty"`
	BillingCompany                string               `json:"billing_company,omitempty"`
	BillingVATId                  string               `json:"billing_vat_id,omitempty"`
	BillingCity                   string               `json:"billing_city,omitempty"`
	BillingPostCode               string               `json:"billing_post_code,omitempty"`
	BillingStreet1                string               `json:"billing_street1,omitempty"`
	BillingStreet2                string               `json:"billing_street2,omitempty"`
	BillingState                  string               `json:"billing_state,omitempty"`
	BillingCountry                string               `json:"billing_country,omitempty"`
	BillingPhone                  string               `json:"billing_phone,omitempty"`
	BillingEmail                  string               `json:"billing_email,omitempty"`
	BillingCountryCode            string               `json:"billing_country_code,omitempty"`
	Cod                           *bool                `json:"cod,omitempty"`
	CodAmount                     *float64             `json:"cod_amount,omitempty"`
	CodAccountNumber              string               `json:"cod_account_number,omitempty"`
	Items                         []*OrderItem2        `json:"items,omitempty"`
	Tags                          []string             `json:"tags,omitempty"`
	CustomerObject                *Customer2           `json:"customerObject,omitempty"`
	InternalState                 *InternalOrderState  `json:"internal_state,omitempty"`
	ExternalDeliveryIds           []*ExternalDeliveryID `json:"external_delivery_ids,omitempty"`
	ExternalDeliveryIdsCanceled   []*ExternalDeliveryID `json:"external_delivery_ids_canceled,omitempty"`
	IntegrationClientName         string               `json:"integration_client_name,omitempty"`
	WMSIntegrationName            string               `json:"wms_integration_name,omitempty"`
	AdditionalWMSIntegrationName  string               `json:"additional_wms_integration_name,omitempty"`
	StatusHistory                 []*StatusHistoryEntry `json:"status_history,omitempty"`
	InternalDeliveryMethod        string               `json:"internal_delivery_method,omitempty"`
	DeliveryMethodMap             string               `json:"delivery_method_map,omitempty"`
	ReturnDocumentsType           string               `json:"return_documents_type,omitempty"`
	DeliveryHour                  string               `json:"delivery_hour,omitempty"`
	TransactionID                 string               `json:"transaction_id,omitempty"`
	ExternalSellerID              string               `json:"external_seller_id,omitempty"`
	Origin                        string               `json:"origin,omitempty"`
	DeliveryDueDate               string               `json:"delivery_due_date,omitempty"`
	DeliveryStatus                string               `json:"delivery_status,omitempty"`
	PackageType                   string               `json:"package_type,omitempty"`
	NumberOfPackages              *int64               `json:"number_of_packages,omitempty"`
	PackedAt                      string               `json:"packed_at,omitempty"`
	PackedBy                      string               `json:"packed_by,omitempty"`
	PackingStartedAt              string               `json:"packing_started_at,omitempty"`
	PackingStation                string               `json:"packing_station,omitempty"`
	PackedManually                *bool                `json:"packed_manually,omitempty"`
	PickedAt                      string               `json:"picked_at,omitempty"`
	PickedBy                      string               `json:"picked_by,omitempty"`
	IsValid                       *bool                `json:"isValid,omitempty"`
	// ValidationErrors contains validation error details (array).
	ValidationErrors              any                  `json:"validationErrors,omitempty"`
	WMSState                      string               `json:"wmsState,omitempty"`
	ExternalState                 string               `json:"externalState,omitempty"`
	StationID                     string               `json:"station_id,omitempty"`
	CreateInvoice                 *bool                `json:"create_invoice,omitempty"`
	ContactPerson                 string               `json:"contact_person,omitempty"`
	ContactPhone                  string               `json:"contact_phone,omitempty"`
	WMSOrderNumber                string               `json:"wms_order_number,omitempty"`
	ConfirmedInWMS                *bool                `json:"confirmed_in_wms,omitempty"`
	PaymentTransactionID          string               `json:"payment_transaction_id,omitempty"`
	// Attachments is order attachment metadata (array).
	Attachments                   any                  `json:"attachments,omitempty"`
	ExternalOrdersPackageID       string               `json:"external_orders_package_id,omitempty"`
	ExternalOrdersPackageQuantity *int64               `json:"external_orders_package_quantity,omitempty"`
	Merchant                      string               `json:"merchant,omitempty"`
	DataHash                      string               `json:"data_hash,omitempty"`
	CourierInfo                   string               `json:"courier_info,omitempty"`
	WMSId                         string               `json:"wms_id,omitempty"`
	ContractorID                  string               `json:"contractor_id,omitempty"`
	OrderType                     string               `json:"order_type,omitempty"`
	DeliveryToPostBranch          *bool                `json:"delivery_to_post_branch,omitempty"`
	ContactFirstName              string               `json:"contact_first_name,omitempty"`
	ContactLastName               string               `json:"contact_last_name,omitempty"`
	ContactCity                   string               `json:"contact_city,omitempty"`
	ContactStreet                 string               `json:"contact_street,omitempty"`
	ContactBuildingNumber         string               `json:"contact_building_number,omitempty"`
	ContactFlatNumber             string               `json:"contact_flat_number,omitempty"`
	ContactPostCode               string               `json:"contact_post_code,omitempty"`
	ContactPostOffice             string               `json:"contact_post_office,omitempty"`
	ContactCountryCode            string               `json:"contact_country_code,omitempty"`
	ContactCountryName            string               `json:"contact_country_name,omitempty"`
	ContactEmail                  string               `json:"contact_email,omitempty"`
	TimeToDeadline                *int64               `json:"time_to_deadline,omitempty"`
	BatchType                     string               `json:"batch_type,omitempty"`
	OrderItemsCount               *int64               `json:"order_items_count,omitempty"`
	AdditionalClientOrderNumber   string               `json:"additional_client_order_number,omitempty"`
	BatchRealizationProgress      *float64             `json:"batch_realization_progress,omitempty"`
	ExecutionDueDate              string               `json:"execution_due_date,omitempty"`
	// CustomProperties is instance-specific custom fields (array or object).
	CustomProperties              any                  `json:"custom_properties,omitempty"`
	Priority                      *int64               `json:"priority,omitempty"`
	Type                          *int64               `json:"type,omitempty"`
	Status                        *OrderStatus         `json:"status,omitempty"`
	// Jobs is the list of background jobs associated with this order.
	Jobs                          []any                `json:"jobs,omitempty"`
	CreatedAt                     string               `json:"created_at,omitempty"`
	UpdatedAt                     string               `json:"updated_at,omitempty"`
	ClonedID                      string               `json:"cloned_id,omitempty"`
	ShipmentDocumentID            string               `json:"shipment_document_id,omitempty"`
	PackageValue                  *float64             `json:"package_value,omitempty"`
	PackageValueCurrency          string               `json:"package_value_currency,omitempty"`
	WaybillCreated                *bool                `json:"waybill_created,omitempty"`
	ExportedAt                    string               `json:"exported_at,omitempty"`
	CreatedBy                     string               `json:"created_by,omitempty"`
	UpdatedBy                     string               `json:"updated_by,omitempty"`
	AssignedTo                    string               `json:"assigned_to,omitempty"`
	Client                        *OrderProject        `json:"client,omitempty"`
	ErpOrder                      *bool                `json:"erp_order,omitempty"`
	// AdditionalDeliveryConfiguration is extra carrier delivery config (object or array).
	AdditionalDeliveryConfiguration any                `json:"additional_delivery_configuration,omitempty"`
	PackageGenerationError        *bool                `json:"package_generation_error,omitempty"`
	PackageGenerationErrorMessage string               `json:"package_generation_error_message,omitempty"`
	WMSErrorMessage               string               `json:"wms_error_message,omitempty"`
	Validator                     string               `json:"validator,omitempty"`
	RejectedBy                    string               `json:"rejected_by,omitempty"`
	RejectionReason               string               `json:"rejection_reason,omitempty"`
	// SourceOrders is the list of source order IDs (for consolidated orders).
	SourceOrders                  []any                `json:"source_orders,omitempty"`
	// SourceOrdersObjects is the list of source order objects (for consolidated orders).
	SourceOrdersObjects           []any                `json:"source_orders_objects,omitempty"`
	// ChildOrders is the list of child order IDs (for divided orders).
	ChildOrders                   []any                `json:"child_orders,omitempty"`
	// ChildOrdersObjects is the list of child order objects (for divided orders).
	ChildOrdersObjects            []any                `json:"child_orders_objects,omitempty"`
	CreatingOperation             string               `json:"creating_operation,omitempty"`
	IsConsolidated                *bool                `json:"is_consolidated,omitempty"`
	IsDivided                     *bool                `json:"is_divided,omitempty"`
	Video                         *OrderVideo          `json:"video,omitempty"`
	AssignedFulfillmentSite       string               `json:"assigned_fulfillment_site,omitempty"`
	AssignedFulfillmentSiteObject *FulfillmentSite     `json:"assigned_fulfillment_site_object,omitempty"`
	// WorkflowHistory is the state machine transition history (nullable array).
	WorkflowHistory               any                  `json:"workflow_history,omitempty"`
	// AdditionalWMSOrders is the list of additional WMS order references.
	AdditionalWMSOrders           []any                `json:"additional_wms_orders,omitempty"`
	IsAdditionalWMSOrder          *bool                `json:"is_additional_wms_order,omitempty"`
	SortingRack                   string               `json:"sorting_rack,omitempty"`
	ReceiptFiscalPrintStatus      string               `json:"receipt_fiscal_print_status,omitempty"`
	IsDelivered                   *bool                `json:"is_delivered,omitempty"`
	IsCodReceived                 *bool                `json:"is_cod_received,omitempty"`
	// CodCustomProperties is COD payment metadata (object with paymentMethod, paymentDate).
	CodCustomProperties           any                  `json:"cod_custom_properties,omitempty"`
	AllowedActions                map[string]bool      `json:"allowed_actions,omitempty"`
	PackingSuggestion             *PackingSuggestion   `json:"packing_suggestion,omitempty"`
	AdditionalServices            []*AdditionalService `json:"additional_services,omitempty"`
	DeliveryRecipientFirstName    string               `json:"delivery_recipient_first_name,omitempty"`
	DeliveryRecipientLastName     string               `json:"delivery_recipient_last_name,omitempty"`
	LastTimeClosedAt              string               `json:"last_time_closed_at,omitempty"`
}

// OrderAttachment is a Linker Cloud API resource.
type OrderAttachment struct {
	ID string `json:"id,omitempty"`
	// File encoded as base64
	Content string `json:"content,omitempty"`
	// File encoded as md5
	ContentMd5 string `json:"content_md5,omitempty"`
	// Project ID
	DepotID *int64 `json:"depot_id,omitempty"`
	// Document number, eg. invoice number
	DocumentNumber string `json:"document_number,omitempty"`
	// Attachment reference
	Reference string `json:"reference,omitempty"`
	// Attachment type
	Type string `json:"type,omitempty"`
}

// OrderItem is the order line item as returned in [Order] (camelCase JSON tags).
// See [OrderItem2] for the snake_case variant.
type OrderItem struct {
	ID                       *int64   `json:"id,omitempty"`
	ExternalID               string   `json:"externalId,omitempty"`
	ProductExternalID        string   `json:"productExternalID,omitempty"`
	ProductVariantExternalID string   `json:"productVariantExternalID,omitempty"`
	VATCode                  string   `json:"vatCode,omitempty"`
	PriceGross               *float64 `json:"priceGross,omitempty"`
	PriceNet                 *float64 `json:"priceNet,omitempty"`
	PriceNetAfterDiscount    *float64 `json:"priceNetAfterDiscount,omitempty"`
	PriceGrossAfterDiscount  *float64 `json:"priceGrossAfterDiscount,omitempty"`
	DiscountType             string   `json:"discountType,omitempty"`
	DiscountValue            *float64 `json:"discountValue,omitempty"`
	Unit                     string   `json:"unit,omitempty"`
	Quantity                 *float64 `json:"quantity,omitempty"`
	SKU                      string   `json:"sku,omitempty"`
	EAN                      string   `json:"ean,omitempty"`
	Description              string   `json:"description,omitempty"`
	AdditionalInfo           string   `json:"additionalInfo,omitempty"`
	Weight                   *float64 `json:"weight,omitempty"`
	WeightUnit               string   `json:"weightUnit,omitempty"`
	Length                   *float64 `json:"length,omitempty"`
	Width                    *float64 `json:"width,omitempty"`
	Depth                    *float64 `json:"depth,omitempty"`
	DimensionsUnit           string   `json:"dimensionsUnit,omitempty"`
	Origin                   string   `json:"origin,omitempty"`
	SourceData               []any    `json:"source_data,omitempty"`
	Carrier                  string   `json:"carrier,omitempty"`
	CustomProperties         []any    `json:"custom_properties,omitempty"`
	AvailableQuantity        *int64   `json:"available_quantity,omitempty"`
	ProductKitSKU            string   `json:"productKitSku,omitempty"`
	SerialNumbers            []any    `json:"serial_numbers,omitempty"`
	BatchNumbers             []any    `json:"batch_numbers,omitempty"`
}

// OrderItem2 is the order line item as returned in [Order2] (snake_case JSON tags).
// See [OrderItem] for the camelCase variant.
type OrderItem2 struct {
	ID                        *int64   `json:"id,omitempty"`
	ExternalID                *int64   `json:"external_id,omitempty"`
	ProductExternalID         *int64   `json:"product_external_id,omitempty"`
	ProductVariantExternalID  string   `json:"product_variant_external_id,omitempty"`
	VATCode                   string   `json:"vat_code,omitempty"`
	PriceGross                *float64 `json:"price_gross,omitempty"`
	PriceNet                  *float64 `json:"price_net,omitempty"`
	PriceNetAfterDiscount     *float64 `json:"price_net_after_discount,omitempty"`
	PriceGrossAfterDiscount   *float64 `json:"price_gross_after_discount,omitempty"`
	DiscountType              string   `json:"discount_type,omitempty"`
	DiscountValue             *float64 `json:"discount_value,omitempty"`
	Unit                      string   `json:"unit,omitempty"`
	Quantity                  *float64 `json:"quantity,omitempty"`
	SKU                       string   `json:"sku,omitempty"`
	EAN                       string   `json:"ean,omitempty"`
	Description               string   `json:"description,omitempty"`
	AdditionalInfo            string   `json:"additional_info,omitempty"`
	Weight                    *float64 `json:"weight,omitempty"`
	WeightUnit                string   `json:"weight_unit,omitempty"`
	Length                    *float64 `json:"length,omitempty"`
	Width                     *float64 `json:"width,omitempty"`
	Depth                     *float64 `json:"depth,omitempty"`
	DimensionsUnit            string   `json:"dimensions_unit,omitempty"`
	Origin                    string   `json:"origin,omitempty"`
	SourceData                []any    `json:"source_data,omitempty"`
	TagID                     string   `json:"tag_id,omitempty"`
	ExpirationDates           []any    `json:"expiration_dates,omitempty"`
	SerialNumbers             []any    `json:"serial_numbers,omitempty"`
	LockCode                  string   `json:"lock_code,omitempty"`
	BatchNumbers              []any    `json:"batch_numbers,omitempty"`
	AuctionID                 string   `json:"auction_id,omitempty"`
	Carrier                   string   `json:"carrier,omitempty"`
	Merchant                  string   `json:"merchant,omitempty"`
	Fulfillments              []any    `json:"fulfillments,omitempty"`
	CustomProperties          []any    `json:"custom_properties,omitempty"`
	AvailableQuantity         *int64   `json:"available_quantity,omitempty"`
	ProductKitSKU             string   `json:"product_kit_sku,omitempty"`
	Order                     *Order2  `json:"order,omitempty"`
	VariantExternalID         *int64   `json:"variant_external_id,omitempty"`
	Ordered                   *int64   `json:"ordered,omitempty"`
	Realized                  *int64   `json:"realized,omitempty"`
	ToBook                    *int64   `json:"to_book,omitempty"`
	ToAchieve                 *int64   `json:"to_achieve,omitempty"`
	Converter                 *float64 `json:"converter,omitempty"`
	Margin                    *float64 `json:"margin,omitempty"`
	DispatchedQuantity        *int64   `json:"dispatched_quantity,omitempty"`
	DispatchedLineValueNet    *float64 `json:"dispatched_line_value_net,omitempty"`
	CreatedAt                 string   `json:"created_at,omitempty"`
	UpdatedAt                 string   `json:"updated_at,omitempty"`
	ProductObject             *Product `json:"product_object,omitempty"`
}

// OrderItemType is a Linker Cloud API resource.
type OrderItemType struct {
	ID                string              `json:"id,omitempty"`
	ExternalID        string              `json:"externalId,omitempty"`
	ProductExternalID string              `json:"productExternalID,omitempty"`
	VATCode           string              `json:"vat_code,omitempty"`
	// Ordered quantity
	Quantity   string `json:"quantity,omitempty"`
	PriceGross string `json:"price_gross,omitempty"`
	PriceNet   string `json:"price_net,omitempty"`
	Unit       string `json:"unit,omitempty"`
	// Product name
	Description string `json:"description,omitempty"`
	// Product SKU
	SKU      string `json:"sku,omitempty"`
	Weight   string `json:"weight,omitempty"`
	LockCode string `json:"lock_code,omitempty"`
	// Serial numbers
	SerialNumbers    []*UnstructuredType `json:"serial_numbers"`
	CustomProperties []*UnstructuredType `json:"custom_properties"`
	SourceData       []*UnstructuredType `json:"source_data"`
	BatchNumbers     []*UnstructuredType `json:"batch_numbers"`
}

// OrderMaterial is a Linker Cloud API resource.
type OrderMaterial struct {
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// OrderProject is the depot/project (client facility) associated with an order.
type OrderProject struct {
	Name                    string `json:"name,omitempty"`
	ShortName               string `json:"shortName,omitempty"`
	OnlyCapitalLettersInSku *bool  `json:"onlyCapitalLettersInSku,omitempty"`
}

// OrderReturn is a Linker Cloud API resource.
type OrderReturn struct {
	ID                     string `json:"id,omitempty"`
	OrderID                string `json:"orderId,omitempty"`
	DepotID                any    `json:"depotId,omitempty"`
	Carrier                string `json:"carrier,omitempty"`
	OrderStatus            string `json:"orderStatus,omitempty"`
	InternalDeliveryMethod any    `json:"internalDeliveryMethod,omitempty"`
	ReturnNumber           string `json:"returnNumber,omitempty"`
	OrderNumber            string `json:"orderNumber,omitempty"`
	SupplierOrderNumber    string `json:"supplierOrderNumber,omitempty"`
	ReturnReason           string `json:"returnReason,omitempty"`
	Accepted               *bool  `json:"accepted,omitempty"`
	AcceptedAt             any    `json:"acceptedAt,omitempty"`
}

// OrderStatus is a Linker Cloud API resource.
type OrderStatus struct {
	Value string `json:"value,omitempty"`
	Cache any    `json:"cache,omitempty"`
	Name  string `json:"name,omitempty"`
}

// OrderType is a Linker Cloud API resource.
type OrderType struct {
	// Order number in source system
	ClientOrderNumber string `json:"clientOrderNumber,omitempty"`
	// Order ID in source system
	ExternalID string `json:"externalId,omitempty"`
	// Additional order number used for identification in third party systems
	AdditionalOrderNumber string `json:"additionalOrderNumber,omitempty"`
	// Order payment method
	PaymentMethod string `json:"paymentMethod,omitempty"`
	// Order creation date, Y-m-d H:i:s
	OrderDate string `json:"orderDate"`
	// Date when order should be fulfilled in warehouse
	ExecutionDate string `json:"executionDate"`
	// Order carrier name
	Carrier string `json:"carrier,omitempty"`
	// Order courier info
	CourierInfo string `json:"courier_info,omitempty"`
	// Delivery company name
	DeliveryCompany string `json:"deliveryCompany,omitempty"`
	// Delivery address company name
	DeliveryRecipient string `json:"deliveryRecipient,omitempty"`
	// Delivery recipient login
	DeliveryRecipientLogin string `json:"deliveryRecipientLogin,omitempty"`
	// Delivery address contact phone
	DeliveryPhone string `json:"deliveryPhone,omitempty"`
	// Delivery address contact email
	DeliveryEmail string `json:"deliveryEmail"`
	// Delivery address street
	DeliveryStreet string `json:"deliveryStreet,omitempty"`
	// Delivery address post code
	DeliveryPostCode string `json:"deliveryPostCode,omitempty"`
	// Delivery address city
	DeliveryCity string `json:"deliveryCity,omitempty"`
	// Order delivery country specified as SO 3166-2, eg. DE, US, PL.
	DeliveryCountry string `json:"deliveryCountry,omitempty"`
	// Delivery point id used if carrier is able to handle pickup points (ex. DHL DE post number)
	DeliveryPointID string `json:"deliveryPointID,omitempty"`
	// Delivery point identifier used if carrier is able to handle pickup points
	DeliveryPointName string `json:"deliveryPointName,omitempty"`
	// Delivery point address
	DeliveryPointAddress string `json:"deliveryPointAddress,omitempty"`
	// Delivery point post code
	DeliveryPointPostcode string `json:"deliveryPointPostcode,omitempty"`
	// Delivery point city
	DeliveryPointCity string `json:"deliveryPointCity,omitempty"`
	// Warehouse ID
	DepotID string `json:"depotId,omitempty"`
	// Order Cash On Delivery Amount
	CodAmount *Money `json:"codAmount"`
	// Shipment price
	ShipmentPrice *Money `json:"shipmentPrice"`
	// Shipment price net
	ShipmentPriceNet *Money `json:"shipmentPriceNet"`
	// discount value
	Discount *Money `json:"discount"`
	// Order items
	Items []*OrderItemType `json:"items"`
	// Order total gross price
	PriceGross string `json:"priceGross,omitempty"`
	// Order total net price
	PriceNet string `json:"priceNet,omitempty"`
	// Order currency specified as ISO 4217, eg. USD, EUR, PLN.
	CurrencySymbol string `json:"currencySymbol,omitempty"`
	// Billing address company name
	BillingCompany string `json:"billingCompany,omitempty"`
	// Billing Tax ID or VAT ID
	BillingVATId string `json:"billingVatID,omitempty"`
	// Billing address contact person first name
	BillingFirstName string `json:"billingFirstName,omitempty"`
	// Billing address contact person last name
	BillingLastName string `json:"billingLastName,omitempty"`
	// Billing address contact email address
	BillingEmail string `json:"billingEmail,omitempty"`
	// Billing address contact phone number
	BillingPhone string `json:"billingPhone,omitempty"`
	// Order billing street line 1
	BillingStreet1 string `json:"billingStreet1,omitempty"`
	// Order billing street line 1
	BillingStreet2 string `json:"billingStreet2,omitempty"`
	// Billing address post code
	BillingPostCode string `json:"billingPostCode,omitempty"`
	// Billing address city
	BillingCity string `json:"billingCity,omitempty"`
	// Billing address state
	BillingState string `json:"billingState,omitempty"`
	// Billing address country, specified as description
	BillingCountry string `json:"billingCountry,omitempty"`
	// Billing address country, specified as SO 3166-2, eg. DE, US, PL.
	BillingCountryCode string `json:"billingCountryCode,omitempty"`
	// Order comments
	Comments string `json:"comments,omitempty"`
	// Order internal comments
	InternalComments string `json:"internalComments,omitempty"`
	// Order type descriptor, eg. B2B, B2C
	OrderType               string              `json:"orderType,omitempty"`
	ReturnDocumentsType     string              `json:"returnDocumentsType,omitempty"`
	InternalDeliveryMethod  string              `json:"internalDeliveryMethod,omitempty"`
	DeliveryMethodMap       string              `json:"deliveryMethodMap,omitempty"`
	NumberOfPackages        string              `json:"numberOfPackages,omitempty"`
	DeliveryConfiguration   any                 `json:"deliveryConfiguration,omitempty"`
	// Payment transaction ID
	PaymentTransactionID string              `json:"paymentTransactionID"`
	CustomProperties     []*UnstructuredType `json:"customProperties"`
	// Date when order should be send
	ExecutionDueDate string `json:"executionDueDate"`
	// Tags
	Tags []string `json:"tags"`
	// Validation errors
	ValidationErrors []*UnstructuredType `json:"validationErrors"`
	RejectionReason  string              `json:"rejectionReason,omitempty"`
}

// OrderVideo is a Linker Cloud API resource.
type OrderVideo struct {
	// Status is the video recording status: NO_RECORDING, RECORDED,
	// DOWNLOADING, CONVERTING, READY, or VIDEO_EXPIRED.
	Status string `json:"status,omitempty"`
}

// PaymentItemRequest is a Linker Cloud API resource.
type PaymentItemRequest struct {
	ID                   string `json:"id,omitempty"`
	ExternalID           string `json:"externalId,omitempty"`
	ClientOrderNumber    string `json:"clientOrderNumber,omitempty"`
	PaymentStatus        string `json:"paymentStatus"`
	PaymentDate          string `json:"paymentDate,omitempty"`
	PaymentTransactionID string `json:"paymentTransactionID,omitempty"`
	PaymentMethod        string `json:"paymentMethod,omitempty"`
}

// PaymentRequest is a Linker Cloud API resource.
type PaymentRequest struct {
	Items []*PaymentItemRequest `json:"items"`
}

// MarkPickedRequest is the request body for PUT /order/picked.
type MarkPickedRequest struct {
	// Number is the order number (required).
	Number string `json:"number"`
	// Tags for the order.
	Tags string `json:"tags,omitempty"`
	// WZNumberWMS is the WMS goods-dispatch reference number.
	WZNumberWMS string `json:"WZnumberWMS,omitempty"`
	// PickedContainers is the list of picked items/containers.
	PickedContainers []PickedContainer `json:"pickedContainers,omitempty"`
}

// PickedContainer is a single picked item within a [MarkPickedRequest].
type PickedContainer struct {
	// Name of the container/item (defaults to SKU if not provided).
	Name string `json:"name,omitempty"`
	// SKU of the picked product (required).
	SKU string `json:"sku"`
	// Quantity picked (required).
	Quantity int64 `json:"quantity"`
}

// PickingConfirmation is a Linker Cloud API resource.
type PickingConfirmation struct {
	// Order number
	OrderNumber string `json:"orderNumber,omitempty"`
	// Goods dispatch document from external application
	WzOrderNumber string `json:"wzOrderNumber,omitempty"`
	// Picked containers
	Containers any `json:"containers,omitempty"`
}

// StatusHistoryEntry records a single status transition on an order.
type StatusHistoryEntry struct {
	Time   int64  `json:"time"`
	Status string `json:"status"`
}

// OperatorsData contains per-package shipping operator data within a delivery.
type OperatorsData struct {
	PackageID    string `json:"package_id,omitempty"`
	TrackingURL  string `json:"tracking_url,omitempty"`
	ParcelID     string `json:"parcel_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Extra        []any  `json:"extra,omitempty"`
	PackageIndex *int64 `json:"package_index,omitempty"`
}

// ExternalDeliveryID represents one shipment entry as returned in the
// externalDeliveryIds field on an order.
type ExternalDeliveryID struct {
	ID            string           `json:"id,omitempty"`
	Time          int64            `json:"time,omitempty"`
	TypeName      string           `json:"type_name,omitempty"`
	OperatorsData []*OperatorsData `json:"operators_data,omitempty"`
	Status        string           `json:"status,omitempty"`
}

// TrackingNumberRequest is the request body for setting a single tracking number
// on an order via PUT /orders/{id}/trackingnumber.
type TrackingNumberRequest struct {
	TrackingNumber string `json:"tracking_number"`
	Operator       string `json:"operator"`
	ParcelID       string `json:"parcel_id,omitempty"`
	TrackingURL    string `json:"tracking_url,omitempty"`
	// PackedAt is the pack timestamp in "Y-m-d H:i:s" format.
	PackedAt string `json:"packed_at,omitempty"`
}

// TrackingNumbersRequest is the request body for setting multiple tracking numbers
// on an order via PUT /orders/{id}/trackingnumbers.
type TrackingNumbersRequest struct {
	TrackingNumbers []TrackingNumberRequest `json:"tracking_numbers"`
	// PackedAt applies to all tracking numbers, in "Y-m-d H:i:s" format.
	PackedAt string `json:"packed_at,omitempty"`
}

// OrderReturnType is the request body for creating an order return
// via POST /orderreturns.
type OrderReturnType struct {
	// OrderNumber is the source order number (required).
	OrderNumber string `json:"orderNumber"`
	// ReturnType is the return type enum value (required).
	ReturnType string `json:"returnType"`
	// IsAccepted marks the return as accepted on creation.
	IsAccepted bool `json:"isAccepted"`
	DepotID    *int64 `json:"depotId,omitempty"`
	Carrier    string `json:"carrier,omitempty"`
	// TrackingNumber is the return shipment tracking number.
	TrackingNumber string `json:"trackingNumber,omitempty"`
	ReturnReason   string `json:"returnReason,omitempty"`
	// DeliveryCountry is the ISO 3166-1 country code.
	DeliveryCountry   string `json:"deliveryCountry,omitempty"`
	DeliveryPostCode  string `json:"deliveryPostCode,omitempty"`
	DeliveryCity      string `json:"deliveryCity,omitempty"`
	DeliveryStreet    string `json:"deliveryStreet,omitempty"`
	DeliveryCompany   string `json:"deliveryCompany,omitempty"`
	DeliveryRecipient string `json:"deliveryRecipient,omitempty"`
	// DeliveryPhone must be in E.164 format, e.g. "+48111222333".
	DeliveryPhone    string         `json:"deliveryPhone,omitempty"`
	DeliveryEmail    string         `json:"deliveryEmail,omitempty"`
	CustomProperties map[string]any `json:"customProperties,omitempty"`
	// Items is the list of returned items (required, at least one).
	Items []OrderReturnItemType `json:"items"`
}

// OrderReturnItemType is a single item in an order return request.
type OrderReturnItemType struct {
	SKU      string `json:"sku"`
	Quantity int64  `json:"quantity"`
	// QuantityByCondition maps condition names to quantities; sum must equal Quantity.
	QuantityByCondition map[string]int `json:"quantityByCondition"`
}
