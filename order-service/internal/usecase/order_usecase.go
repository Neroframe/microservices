package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
)

type orderUsecase struct {
	repo      domain.OrderRepository
	publisher domain.OrderEventPublisher
}

func NewOrderUsecase(r domain.OrderRepository, p domain.OrderEventPublisher) domain.OrderUsecase {
	return &orderUsecase{repo: r, publisher: p}
}

func (u *orderUsecase) Create(ctx context.Context, o *domain.Order) error {
	if o.UserID == "" || len(o.Items) == 0 {
		return errors.New("invalid order: missing user or items")
	}
	o.Status = "Pending"
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()

	err := u.repo.Create(ctx, o)
	if err != nil {
		return err
	}

	event := domain.OrderCreatedEvent{
		OrderID: o.ID,
		UserID:  o.UserID,
		Items:   o.Items,
	}
	return u.publisher.PublishOrderCreated(ctx, event)
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
	err := u.repo.Update(ctx, o)
	if err != nil {
		return err
	}

	event := domain.OrderUpdatedEvent{
		OrderID: o.ID,
		Status:  o.Status,
	}

	return u.publisher.PublishOrderUpdated(ctx, event)
}

func (u *orderUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("missing order ID")
	}

	err := u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	event := domain.OrderDeletedEvent{
		OrderID: id,
	}

	return u.publisher.PublishOrderDeleted(ctx, event)
}

func (u *orderUsecase) ListByUserID(ctx context.Context, userID string) ([]*domain.Order, error) {
	if userID == "" {
		return nil, errors.New("missing user ID")
	}
	return u.repo.ListByUserID(ctx, userID)
}
