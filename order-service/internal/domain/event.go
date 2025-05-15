package domain

import "context"

type OrderCreatedEvent struct {
    OrderID string
    UserID  string
    Items   []OrderItem
}

type OrderUpdatedEvent struct {
    OrderID string
    Status  string
}

type OrderDeletedEvent struct {
    OrderID string
}

type OrderEventPublisher interface {
	PublishOrderCreated(ctx context.Context, payload any) error
	PublishOrderUpdated(ctx context.Context, payload any) error
	PublishOrderDeleted(ctx context.Context, payload any) error
}