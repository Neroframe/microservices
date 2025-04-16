package usecase

import (
	"context"
	"errors"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
)

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryUsecase(repo domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: repo,
	}
}

func (u *categoryUsecase) Create(ctx context.Context, c *domain.Category) error {
	if c.Name == "" {
		return errors.New("category name cannot be empty")
	}
	if len(c.Name) > 36 {
		return errors.New("category name cannot exceed 36 characters")
	}

	c.NormalizeName()

	return u.categoryRepo.Create(ctx, c)
}

func (u *categoryUsecase) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	if id == "" {
		return nil, errors.New("category ID cannot be empty")
	}
	return u.categoryRepo.GetByID(ctx, id)
}

func (u *categoryUsecase) Update(ctx context.Context, c *domain.Category) error {
	if c.ID == "" {
		return errors.New("category ID cannot be empty")
	}
	if c.Name == "" {
		return errors.New("category name cannot be empty")
	}
	if len(c.Name) > 36 {
		return errors.New("category name cannot exceed 36 characters")
	}

	c.NormalizeName()

	return u.categoryRepo.Update(ctx, c)
}

func (u *categoryUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("category ID cannot be empty")
	}
	return u.categoryRepo.Delete(ctx, id)
}

func (u *categoryUsecase) List(ctx context.Context) ([]*domain.Category, error) {
	return u.categoryRepo.List(ctx)
}
