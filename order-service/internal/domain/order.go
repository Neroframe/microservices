package domain

import (
	"context"
	"errors"
	"time"
)

var ErrNotFound = errors.New("not found")

type Order struct {
	ID        string      `bson:"_id,omitempty"`
	UserID    string      `bson:"user_id"`
	Items     []OrderItem `bson:"items"`
	Status    string      `bson:"status"` // e.g. "Pending", "Shipped", "Delivered"
	CreatedAt time.Time   `bson:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at"`
}

type OrderItem struct {
	ProductID string `bson:"product_id"`
	Quantity  int    `bson:"quantity"`
}

type OrderRepository interface {
	Create(ctx context.Context, o *Order) error
	GetByID(ctx context.Context, id string) (*Order, error)
	Update(ctx context.Context, o *Order) error
	Delete(ctx context.Context, id string) error
	ListByUserID(ctx context.Context, userID string) ([]*Order, error)
}

type OrderUsecase interface {
	Create(ctx context.Context, o *Order) error
	GetByID(ctx context.Context, id string) (*Order, error)
	Update(ctx context.Context, o *Order) error
	Delete(ctx context.Context, id string) error
	ListByUserID(ctx context.Context, userID string) ([]*Order, error)
}
