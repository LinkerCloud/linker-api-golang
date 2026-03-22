package models

// OrderTransitionItem represents a single transition as returned by
// GET /public-api/v1/orders/{id}/transition.
type OrderTransitionItem struct {
	Name     string   `json:"name"`
	Froms    []string `json:"froms"`
	Tos      []string `json:"tos"`
	Metadata []any    `json:"metadata"`
}

// OrderTransitionsResponse is the envelope returned by
// GET /public-api/v1/orders/{id}/transition.
type OrderTransitionsResponse struct {
	Transitions          []*OrderTransitionItem `json:"transitions"`
	AvailableTransitions []string               `json:"availableTransitions"`
	Places               any                    `json:"places"`
	Initial              string                 `json:"initial"`
}
