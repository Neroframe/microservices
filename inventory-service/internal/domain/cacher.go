package domain

import "context"

type ProductMemoryCache interface {
	Get(productID string) (*Product, bool)
	Set(product *Product)
	SetMany(products []*Product)
	Delete(productID string)

	GetList() ([]*Product, bool)
}

type ProductRedisCache interface {
	Get(ctx context.Context, productID string) (*Product, error)
	Set(ctx context.Context, product *Product) error
	SetMany(ctx context.Context, products []*Product) error
	Delete(ctx context.Context, productID string) error

	GetList(ctx context.Context) ([]*Product, error)
	SetList(ctx context.Context, products []*Product) error
}
