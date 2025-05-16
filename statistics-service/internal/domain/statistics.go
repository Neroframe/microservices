package domain

import (
	"context"
	"time"

	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"
)

type StatisticsRepository interface {
	// gRPC
	CountOrdersByUser(ctx context.Context, userID string) (int32, error)
	CountTotalUsers(ctx context.Context) (int32, error)
	CountDailyActiveUsers(ctx context.Context) (int32, error)

	// NATS
	InsertOrderCreatedEvent(ctx context.Context, userID, orderID string, ts time.Time) error
	InsertOrderUpdatedEvent(ctx context.Context, userID, orderID string, ts time.Time) error
	InsertOrderDeletedEvent(ctx context.Context, userID, orderID string, ts time.Time) error

	InsertProductCreatedEvent(ctx context.Context, userID, productID string, ts time.Time) error
	InsertProductUpdatedEvent(ctx context.Context, userID, productID string, ts time.Time) error
	InsertProductDeletedEvent(ctx context.Context, userID, productID string, ts time.Time) error
}

type StatisticsUsecase interface {
	// gRPC read methods
	GetUserOrdersStatistics(ctx context.Context, userID string) (*statisticspb.UserOrderStatisticsResponse, error)
	GetUserStatistics(ctx context.Context) (*statisticspb.UserStatisticsResponse, error)

	// NATS event handlers
	HandleOrderCreated(ctx context.Context, evt OrderCreatedEvent) error
	HandleOrderUpdated(ctx context.Context, evt OrderUpdatedEvent) error
	HandleOrderDeleted(ctx context.Context, evt OrderDeletedEvent) error

	HandleProductCreated(ctx context.Context, evt ProductCreatedEvent) error
	HandleProductUpdated(ctx context.Context, evt ProductUpdatedEvent) error
	HandleProductDeleted(ctx context.Context, evt ProductDeletedEvent) error
}
