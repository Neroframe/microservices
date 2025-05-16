package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/inventory-service/pkg/nats"
)

type InventoryEventPublisher struct {
	client *nats.Client
}

var _ domain.InventoryEventPublisher = (*InventoryEventPublisher)(nil) 

func NewInventoryEventPublisher(client *nats.Client) *InventoryEventPublisher {
	return &InventoryEventPublisher{client: client}
}

func (p *InventoryEventPublisher) PublishProductCreated(ctx context.Context, payload domain.ProductCreatedEvent) error {
	return p.publish(ctx, "product.created", payload)
}

func (p *InventoryEventPublisher) PublishProductUpdated(ctx context.Context, payload domain.ProductUpdatedEvent) error {
	return p.publish(ctx, "product.updated", payload)
}

func (p *InventoryEventPublisher) PublishProductDeleted(ctx context.Context, payload domain.ProductDeletedEvent) error {
	return p.publish(ctx, "product.deleted", payload)
}

func (p *InventoryEventPublisher) PublishCategoryCreated(ctx context.Context, payload domain.CategoryCreatedEvent) error {
	return p.publish(ctx, "category.created", payload)
}
func (p *InventoryEventPublisher) PublishCategoryUpdated(ctx context.Context, payload domain.CategoryUpdatedEvent) error {
	return p.publish(ctx, "category.updated", payload)
}
func (p *InventoryEventPublisher) PublishCategoryDeleted(ctx context.Context, payload domain.CategoryDeletedEvent) error {
	return p.publish(ctx, "category.deleted", payload)
}

func (p *InventoryEventPublisher) publish(ctx context.Context, subject string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[NATS] Marshal error: %v", err)
		return fmt.Errorf("payload marshal error: %w", err)
	}
	log.Printf("[NATS] Publishing to subject '%s': %s", subject, data)

	if err := p.client.Conn.Publish(subject, data); err != nil {
		log.Printf("[NATS] Publish failed on subject '%s': %v", subject, err)
		return fmt.Errorf("nats publish error: %w", err)
	}

	log.Printf("[NATS] Successfully published to subject '%s'", subject)
	return nil
}
