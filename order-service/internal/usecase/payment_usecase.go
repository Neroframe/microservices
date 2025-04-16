package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
)

type paymentUsecase struct {
	repo domain.PaymentRepository
}

func NewPaymentUsecase(r domain.PaymentRepository) domain.PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) Create(ctx context.Context, p *domain.Payment) error {
	if p.OrderID == "" || p.Amount <= 0 || p.PaymentMethod == "" {
		return errors.New("invalid payment data")
	}
	p.Status = "Completed"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return u.repo.Create(ctx, p)
}

func (u *paymentUsecase) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	if id == "" {
		return nil, errors.New("payment ID is required")
	}
	return u.repo.GetByID(ctx, id)
}

func (u *paymentUsecase) Update(ctx context.Context, p *domain.Payment) error {
	if p.ID == "" {
		return errors.New("payment ID is required")
	}
	p.UpdatedAt = time.Now()
	return u.repo.Update(ctx, p)
}

func (u *paymentUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("payment ID is required")
	}
	return u.repo.Delete(ctx, id)
}
