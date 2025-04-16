package domain

import (
	"context"
	"time"
)

type Payment struct {
	ID            string    `bson:"_id,omitempty"`
	OrderID       string    `bson:"order_id"`
	Amount        float64   `bson:"amount"`         // The amount paid
	PaymentMethod string    `bson:"payment_method"` // e.g. Credit Card, PayPal
	Status        string    `bson:"status"`         // e.g. Completed, Failed
	CreatedAt     time.Time `bson:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at"`
}

type PaymentRepository interface {
	Create(ctx context.Context, p *Payment) error
	GetByID(ctx context.Context, id string) (*Payment, error)
	Update(ctx context.Context, p *Payment) error
	Delete(ctx context.Context, id string) error
}

type PaymentUsecase interface {
	Create(ctx context.Context, p *Payment) error
	GetByID(ctx context.Context, id string) (*Payment, error)
	Update(ctx context.Context, p *Payment) error
	Delete(ctx context.Context, id string) error
}
