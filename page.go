package linkercloud

// Page is the pagination envelope returned by all Linker Cloud list endpoints.
// The API always wraps collection responses in this structure rather than
// returning a bare JSON array.
type Page[T any] struct {
	// Items contains the records for the current page.
	Items []T `json:"items"`
	// RecordsTotal is the total number of records matching the query
	// (before pagination).
	RecordsTotal int `json:"recordsTotal"`
	// RecordsFiltered is the number of records after applying any filters.
	RecordsFiltered int `json:"recordsFiltered"`
}
