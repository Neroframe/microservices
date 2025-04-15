package usecase

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
)

type ProductUsecase interface {
	Create(ctx context.Context, p *domain.Product) error
	GetByID(ctx context.Context, id string) (*domain.Product, error)
	Update(ctx context.Context, p *domain.Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*domain.Product, error)
}

type productUsecase struct {
	productRepo domain.ProductRepository
}

func NewProductUsecase(r domain.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo: r}
}

func (u *productUsecase) Create(ctx context.Context, p *domain.Product) error {
	return u.productRepo.Create(ctx, p)
}

func (u *productUsecase) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	return u.productRepo.GetByID(ctx, id)
}

func (u *productUsecase) Update(ctx context.Context, p *domain.Product) error {
	return u.productRepo.Update(ctx, p)
}

func (u *productUsecase) Delete(ctx context.Context, id string) error {
	return u.productRepo.Delete(ctx, id)
}

func (u *productUsecase) List(ctx context.Context) ([]*domain.Product, error) {
	return u.productRepo.List(ctx)
}
