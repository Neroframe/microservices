package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
)

type orderUsecase struct {
	repo domain.OrderRepository
}

func NewOrderUsecase(r domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{repo: r}
}

func (u *orderUsecase) Create(ctx context.Context, o *domain.Order) error {
	if o.UserID == "" || len(o.Items) == 0 {
		return errors.New("invalid order: missing user or items")
	}
	o.Status = "Pending"
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	return u.repo.Create(ctx, o)
}

func (u *orderUsecase) GetByID(ctx context.Context, id string) (*domain.Order, error) {
	if id == "" {
		return nil, errors.New("order ID cannot be empty")
	}
	return u.repo.GetByID(ctx, id)
}

func (u *orderUsecase) Update(ctx context.Context, o *domain.Order) error {
	if o.ID == "" {
		return errors.New("missing order ID")
	}
	o.UpdatedAt = time.Now()
	return u.repo.Update(ctx, o)
}

func (u *orderUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("missing order ID")
	}
	return u.repo.Delete(ctx, id)
}

func (u *orderUsecase) ListByUserID(ctx context.Context, userID string) ([]*domain.Order, error) {
	if userID == "" {
		return nil, errors.New("missing user ID")
	}
	return u.repo.ListByUserID(ctx, userID)
}
