package usecase

import (
	"context"
	"errors"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
)

type productUsecase struct {
	productRepo domain.ProductRepository
}

func NewProductUsecase(r domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{productRepo: r}
}

func (u *productUsecase) Create(ctx context.Context, p *domain.Product) error {
	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if len(p.Name) > 36 {
		return errors.New("product name cannot exceed 36 characters")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if p.Stock < 0 {
		return errors.New("stock value cannot be negative")
	}

	p.NormalizeName()

	return u.productRepo.Create(ctx, p)
}

func (u *productUsecase) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	if id == "" {
		return nil, errors.New("product ID cannot be empty")
	}

	return u.productRepo.GetByID(ctx, id)
}

func (u *productUsecase) Update(ctx context.Context, p *domain.Product) error {
	if p.ID == "" {
		return errors.New("product ID cannot be empty")
	}

	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if len(p.Name) > 36 {
		return errors.New("product name cannot exceed 36 characters")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if p.Stock < 0 {
		return errors.New("stock value cannot be negative")
	}

	p.NormalizeName()

	return u.productRepo.Update(ctx, p)
}

func (u *productUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("product ID cannot be empty")
	}
	return u.productRepo.Delete(ctx, id)
}

func (u *productUsecase) List(ctx context.Context) ([]*domain.Product, error) {
	return u.productRepo.List(ctx)
}
