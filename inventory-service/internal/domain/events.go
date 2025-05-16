package domain

import "context"

type ProductCreatedEvent struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}

type ProductUpdatedEvent struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}

type ProductDeletedEvent struct {
	ID string `json:"id"`
}

type CategoryCreatedEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryUpdatedEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryDeletedEvent struct {
	ID string `json:"id"`
}

type InventoryEventPublisher interface {
	PublishProductCreated(ctx context.Context, payload ProductCreatedEvent) error
	PublishProductUpdated(ctx context.Context, payload ProductUpdatedEvent) error
	PublishProductDeleted(ctx context.Context, payload ProductDeletedEvent) error

	PublishCategoryCreated(ctx context.Context, payload CategoryCreatedEvent) error
	PublishCategoryUpdated(ctx context.Context, payload CategoryUpdatedEvent) error
	PublishCategoryDeleted(ctx context.Context, payload CategoryDeletedEvent) error
}
