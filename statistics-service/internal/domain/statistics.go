package domain

import (
	"context"

	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"
)

type StatisticsRepository interface {
	// gRPC
	CountOrdersByUser(ctx context.Context, userID string) (int32, error)
	CountTotalUsers(ctx context.Context) (int32, error)
	CountDailyActiveUsers(ctx context.Context) (int32, error)

	// NATS
	InsertEvent(ctx context.Context, evt Event) error
}

type StatisticsUsecase interface {
	// gRPC read methods
	GetUserOrdersStatistics(ctx context.Context, userID string) (*statisticspb.UserOrderStatisticsResponse, error)
	GetUserStatistics(ctx context.Context) (*statisticspb.UserStatisticsResponse, error)

	// NATS event handler
	HandleEvent(ctx context.Context, evt Event) error
}
