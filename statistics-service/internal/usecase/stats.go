package usecase

import (
	"context"

	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"
)

// StatisticsRepository defines the methods the usecase needs from the data layer.
type StatisticsRepository interface {
	CountOrdersByUser(ctx context.Context, userID string) (int32, error)
	CountTotalUsers(ctx context.Context) (int32, error)
	CountDailyActiveUsers(ctx context.Context) (int32, error)
}

// StatisticsUsecase defines the business operations available.
type StatisticsUsecase interface {
	GetUserOrdersStatistics(ctx context.Context, userID string) (*statisticspb.UserOrderStatisticsResponse, error)
	GetUserStatistics(ctx context.Context) (*statisticspb.UserStatisticsResponse, error)
}

// statisticsUsecase is the default implementation of StatisticsUsecase.
type statisticsUsecase struct {
	repo StatisticsRepository
}

// NewStatisticsUsecase constructs a new StatisticsUsecase.
func NewStatisticsUsecase(repo StatisticsRepository) StatisticsUsecase {
	return &statisticsUsecase{repo: repo}
}

// GetUserOrdersStatistics returns the total orders for a given user.
func (u *statisticsUsecase) GetUserOrdersStatistics(ctx context.Context, userID string) (*statisticspb.UserOrderStatisticsResponse, error) {
	total, err := u.repo.CountOrdersByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &statisticspb.UserOrderStatisticsResponse{
		TotalOrders: total,
	}, nil
}

// GetUserStatistics returns global statistics about users.
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
