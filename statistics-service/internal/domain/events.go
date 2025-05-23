package domain

import "time"

type Event struct {
	UserID    string                 `json:"user_id"`
	EntityID  string                 `json:"entity_id"`
	EntityKey string                 `json:"entity_key"` // e.g. "order_id", "product_id"
	EventType string                 `json:"event_type"` // e.g. "order.created", "product.updated"
	Timestamp time.Time              `json:"timestamp"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

type EventCache interface {
	Get(id string) (*Event, bool)
	Set(evt *Event)
	SetMany(evts []*Event)
	Delete(id string)

	GetList() ([]*Event, bool)
}
