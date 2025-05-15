package nats

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/order-service/pkg/nats"
)

type OrderEventPublisher struct {
	client *nats.Client
}

var _ domain.OrderEventPublisher = (*OrderEventPublisher)(nil) // compile-time check

func NewOrderEventPublisher(client *nats.Client) *OrderEventPublisher {
	return &OrderEventPublisher{client: client}
}

func (p *OrderEventPublisher) PublishOrderCreated(ctx context.Context, payload any) error {
	return p.publish(ctx, "order.created", payload)
}

func (p *OrderEventPublisher) PublishOrderUpdated(ctx context.Context, payload any) error {
	return p.publish(ctx, "order.updated", payload)
}

func (p *OrderEventPublisher) PublishOrderDeleted(ctx context.Context, payload any) error {
	return p.publish(ctx, "order.deleted", payload)
}

func (p *OrderEventPublisher) publish(ctx context.Context, subject string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("payload marshal error: %w", err)
	}
	return p.client.Conn.Publish(subject, data)
}
