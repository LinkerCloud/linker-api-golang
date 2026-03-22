package models

import "encoding/json"

// Money represents a monetary amount with its currency, used in request bodies
// for price fields (e.g. CodAmount, ShipmentPrice).
//
// Amount uses [json.Number] rather than float64 to preserve the exact decimal
// representation from JSON with no floating-point precision loss.
// Callers can convert via Amount.Float64() or Amount.String() when needed.
type Money struct {
	Amount   json.Number `json:"amount"`
	Currency string      `json:"currency"`
}

// UnstructuredType represents an arbitrary key-value object in the Linker API.
// Fields using this type (serial_numbers, custom_properties, storage_units, etc.)
// accept any structure — the API validates them with allow_extra_fields: true.
type UnstructuredType map[string]any
