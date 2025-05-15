package usecase

import (
	"context"
	"time"

	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"
)

// ────────────────────────────────────────────────────────────
// Domain events live here, not in the adapter
// ────────────────────────────────────────────────────────────

type OrderCreatedEvent struct {
	UserID    string    `json:"UserID"`
	OrderID   string    `json:"OrderID"`
	Timestamp time.Time `json:"Timestamp"` // If used
}

type UserRegisteredEvent struct {
	UserID    string    `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

// ────────────────────────────────────────────────────────────
// Repository interface
// ────────────────────────────────────────────────────────────

type StatisticsRepository interface {
	CountOrdersByUser(ctx context.Context, userID string) (int32, error)
	CountTotalUsers(ctx context.Context) (int32, error)
	CountDailyActiveUsers(ctx context.Context) (int32, error)

	InsertOrderCreatedEvent(ctx context.Context, userID, orderID string, ts time.Time) error
	InsertUserRegisteredEvent(ctx context.Context, userID string, ts time.Time) error
}

// ────────────────────────────────────────────────────────────
// Usecase interface
// ────────────────────────────────────────────────────────────

type StatisticsUsecase interface {
	// gRPC read methods
	GetUserOrdersStatistics(ctx context.Context, userID string) (*statisticspb.UserOrderStatisticsResponse, error)
	GetUserStatistics(ctx context.Context) (*statisticspb.UserStatisticsResponse, error)

	// NATS event handlers
	HandleOrderCreated(ctx context.Context, evt OrderCreatedEvent) error
	HandleUserRegistered(ctx context.Context, evt UserRegisteredEvent) error
}

// ────────────────────────────────────────────────────────────
// Usecase implementation
// ────────────────────────────────────────────────────────────

type statisticsUsecase struct {
	repo StatisticsRepository
}

func NewStatisticsUsecase(repo StatisticsRepository) StatisticsUsecase {
	return &statisticsUsecase{repo: repo}
}

func (u *statisticsUsecase) GetUserOrdersStatistics(ctx context.Context, userID string) (*statisticspb.UserOrderStatisticsResponse, error) {
	total, err := u.repo.CountOrdersByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &statisticspb.UserOrderStatisticsResponse{TotalOrders: total}, nil
}

func (u *statisticsUsecase) GetUserStatistics(ctx context.Context) (*statisticspb.UserStatisticsResponse, error) {
	totalUsers, err := u.repo.CountTotalUsers(ctx)
	if err != nil {
		return nil, err
	}
	dailyActive, err := u.repo.CountDailyActiveUsers(ctx)
	if err != nil {
		return nil, err
	}
	return &statisticspb.UserStatisticsResponse{
		TotalUsers:       totalUsers,
		DailyActiveUsers: dailyActive,
	}, nil
}

func (u *statisticsUsecase) HandleOrderCreated(ctx context.Context, evt OrderCreatedEvent) error {
	return u.repo.InsertOrderCreatedEvent(ctx, evt.UserID, evt.OrderID, evt.Timestamp)
}

func (u *statisticsUsecase) HandleUserRegistered(ctx context.Context, evt UserRegisteredEvent) error {
	return u.repo.InsertUserRegisteredEvent(ctx, evt.UserID, evt.Timestamp)
}
