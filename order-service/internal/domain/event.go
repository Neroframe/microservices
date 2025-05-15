package domain

import "context"

type OrderEventPublisher interface {
	PublishOrderCreated(ctx context.Context, payload any) error
	PublishOrderUpdated(ctx context.Context, payload any) error
	PublishOrderDeleted(ctx context.Context, payload any) error
}
