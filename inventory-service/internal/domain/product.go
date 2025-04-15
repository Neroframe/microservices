package domain

import "context"

type Product struct {
	ID       string  `bson:"_id,omitempty"`
	Name     string  `bson:"name"`
	Price    float64 `bson:"price"`
	Category string  `bson:"category"`
	Stock    int     `bson:"stock"`
}

type ProductRepository interface {
	Create(ctx context.Context, p *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Product, error)
}
