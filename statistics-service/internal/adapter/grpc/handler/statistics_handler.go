package handler

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"
)

type StatisticsHandler struct {
	statisticspb.UnimplementedStatisticsServiceServer
	uc usecase.StatisticsUsecase
}

// Inject usecase into the handler
func NewStatisticsHandler(uc usecase.StatisticsUsecase) *StatisticsHandler {
	return &StatisticsHandler{
		uc: uc,
	}
}

func (s *StatisticsHandler) GetUserOrdersStatistics(ctx context.Context, req *statisticspb.UserOrderStatisticsRequest) (*statisticspb.UserOrderStatisticsResponse, error) {
	return s.uc.GetUserOrdersStatistics(ctx, req.UserId)
}

func (s *StatisticsHandler) GetUserStatistics(ctx context.Context, req *statisticspb.UserStatisticsRequest) (*statisticspb.UserStatisticsResponse, error) {
	return s.uc.GetUserStatistics(ctx)
}
