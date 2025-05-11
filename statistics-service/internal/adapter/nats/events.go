package nats

import "time"

// OrderCreatedEvent mirrors the payload published by Order Service.
type OrderCreatedEvent struct {
    UserID    string    `json:"user_id"`
    OrderID   string    `json:"order_id"`
    Timestamp time.Time `json:"timestamp"`
}
