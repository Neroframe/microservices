package domain

import (
	"context"
	"strings"
)

type Category struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

type CategoryRepository interface {
	Create(ctx context.Context, c *Category) error
	GetByID(ctx context.Context, id string) (*Category, error)
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Category, error)
}

type CategoryUsecase interface {
	Create(ctx context.Context, c *Category) error
	GetByID(ctx context.Context, id string) (*Category, error)
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Category, error)
}

func (c *Category) NormalizeName() {
	c.Name = strings.ToLower(strings.TrimSpace(c.Name))
}
