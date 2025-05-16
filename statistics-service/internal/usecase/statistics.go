package usecase

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/domain"
	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"
)

var _ domain.StatisticsUsecase = (*StatisticsUsecase)(nil)

type StatisticsUsecase struct {
	repo domain.StatisticsRepository
}

func NewStatisticsUsecase(repo domain.StatisticsRepository) *StatisticsUsecase {
	return &StatisticsUsecase{repo: repo}
}

func (u *StatisticsUsecase) GetUserOrdersStatistics(ctx context.Context, userID string) (*statisticspb.UserOrderStatisticsResponse, error) {
	total, err := u.repo.CountOrdersByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &statisticspb.UserOrderStatisticsResponse{TotalOrders: total}, nil
}

func (u *StatisticsUsecase) GetUserStatistics(ctx context.Context) (*statisticspb.UserStatisticsResponse, error) {
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

func (u *StatisticsUsecase) HandleOrderCreated(ctx context.Context, evt domain.OrderCreatedEvent) error {
	return u.repo.InsertOrderCreatedEvent(ctx, evt.UserID, evt.OrderID, evt.Timestamp)
}

func (u *StatisticsUsecase) HandleOrderUpdated(ctx context.Context, evt domain.OrderUpdatedEvent) error {
	return u.repo.InsertOrderUpdatedEvent(ctx, evt.UserID, evt.OrderID, evt.Timestamp)
}

func (u *StatisticsUsecase) HandleOrderDeleted(ctx context.Context, evt domain.OrderDeletedEvent) error {
	return u.repo.InsertOrderDeletedEvent(ctx, evt.UserID, evt.OrderID, evt.Timestamp)
}

func (u *StatisticsUsecase) HandleProductCreated(ctx context.Context, evt domain.ProductCreatedEvent) error {
	return u.repo.InsertProductCreatedEvent(ctx, evt.UserID, evt.ProductID, evt.Timestamp)
}

func (u *StatisticsUsecase) HandleProductUpdated(ctx context.Context, evt domain.ProductUpdatedEvent) error {
	return u.repo.InsertProductUpdatedEvent(ctx, evt.UserID, evt.ProductID, evt.Timestamp)
}

func (u *StatisticsUsecase) HandleProductDeleted(ctx context.Context, evt domain.ProductDeletedEvent) error {
	return u.repo.InsertProductDeletedEvent(ctx, evt.UserID, evt.ProductID, evt.Timestamp)
}


