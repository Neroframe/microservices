package usecase

import (
	"context"
	"fmt"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/domain"
	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"
)

var _ domain.StatisticsUsecase = (*StatisticsUsecase)(nil)

type StatisticsUsecase struct {
	repo  domain.StatisticsRepository
	cache domain.EventCache
}

func NewStatisticsUsecase(repo domain.StatisticsRepository, cache domain.EventCache) *StatisticsUsecase {
	return &StatisticsUsecase{repo: repo, cache: cache}
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

// All data (orders & intventory items) must be stored in db and cache, but if you try to retrieve data, then it should be fetched from cache.
func (u *StatisticsUsecase) HandleEvent(ctx context.Context, evt domain.Event) error {
	// Store in Mongo
	if err := u.repo.InsertEvent(ctx, evt); err != nil {
		return fmt.Errorf("repo.InsertEvent: %w", err)
	}

	// Store inmemory cache
	switch evt.EventType {
	case "order.created", "order.updated", "product.created", "product.updated":
		u.cache.Set(&evt)

	case "order.deleted", "product.deleted":
		u.cache.Delete(evt.EntityID)

	default:
		fmt.Printf("unknown event type: %s", evt.EventType)
	}

	return nil
}
