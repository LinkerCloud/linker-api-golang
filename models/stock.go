package models

// Stock is a Linker Cloud API resource.
type Stock struct {
	ExpireIn   *int64       `json:"expireIn,omitempty"`
	TotalCount *int64       `json:"totalCount,omitempty"`
	Items      []*StockItem `json:"items,omitempty"`
}

// StockItem is a Linker Cloud API resource.
type StockItem struct {
	SKU              string `json:"sku,omitempty"`
	FinalSKU         string `json:"final_sku,omitempty"`
	EAN              string `json:"ean,omitempty"`
	Name             string `json:"name,omitempty"`
	Total            *int64 `json:"total,omitempty"`
	Reserved         *int64 `json:"reserved,omitempty"`
	Damaged          *int64 `json:"damaged,omitempty"`
	Free             *int64 `json:"free,omitempty"`
	ReservedInOrders *int64 `json:"reserved_in_orders,omitempty"`
	BatchNumber      string `json:"batch_number,omitempty"`
	ExpirationDate   string `json:"expiration_date,omitempty"`
}

// StockUpdateItemRequest is a Linker Cloud API resource.
type StockUpdateItemRequest struct {
	SKU           string `json:"sku"`
	EAN           string `json:"ean,omitempty"`
	Name          string `json:"name,omitempty"`
	TotalQuantity int64  `json:"totalQuantity"`
}

// StockUpdateRequest is a Linker Cloud API resource.
type StockUpdateRequest struct {
	Items          []*StockUpdateItemRequest `json:"items"`
	FulfilmentSite string                   `json:"fulfilmentSite,omitempty"`
	PublishStocks  string                   `json:"publishStocks,omitempty"`
}
