package models

// Event is a Linker Cloud API resource.
type Event struct {
	ID            *int64 `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	Category      string `json:"category,omitempty"`
	EventName     string `json:"event_name,omitempty"`
	SelectOptions any    `json:"select_options,omitempty"`
}

// MarkingStore is a Linker Cloud API resource.
type MarkingStore struct {
	ID        *int64                `json:"id,omitempty"`
	Type      string                `json:"type,omitempty"`
	Arguments []*MarkingStoreArgument `json:"arguments,omitempty"`
}

// MarkingStoreArgument is a Linker Cloud API resource.
type MarkingStoreArgument struct {
	ID           *int64        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	MarkingStore *MarkingStore `json:"markingStore,omitempty"`
}

// Place is a Linker Cloud API resource.
type Place struct {
	ID       *int64       `json:"id,omitempty"`
	Name     string       `json:"name,omitempty"`
	Workflow *Workflow    `json:"workflow,omitempty"`
	Events   []*PlaceEvent `json:"events,omitempty"`
}

// PlaceEvent is a Linker Cloud API resource.
type PlaceEvent struct {
	ID                *int64  `json:"id,omitempty"`
	OnEnter           *bool   `json:"onEnter,omitempty"`
	Place             *Place  `json:"place,omitempty"`
	Event             *Event  `json:"event,omitempty"`
	Priority          *int64  `json:"priority,omitempty"`
	ExpressionContent string  `json:"expressionContent,omitempty"`
	Expression        string  `json:"expression,omitempty"`
}

// SupportedObject is a Linker Cloud API resource.
type SupportedObject struct {
	ID       *int64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Workflow *Workflow `json:"workflow,omitempty"`
}

// Transition is a Linker Cloud API resource.
type Transition struct {
	ID        *int64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	FromPlace any       `json:"fromPlace,omitempty"`
	ToPlace   any       `json:"toPlace,omitempty"`
	Workflow  *Workflow `json:"workflow,omitempty"`
	Roles     []string  `json:"roles,omitempty"`
}

// Workflow is a Linker Cloud API resource.
type Workflow struct {
	ID           *int64             `json:"id,omitempty"`
	Name         string             `json:"name,omitempty"`
	Key          string             `json:"key,omitempty"`
	Type         string             `json:"type,omitempty"`
	Places       []*Place           `json:"places,omitempty"`
	Supports     []*SupportedObject `json:"supports,omitempty"`
	MarkingStore *MarkingStore      `json:"markingStore,omitempty"`
	Transitions  []*Transition      `json:"transitions,omitempty"`
}
