package handler

import (
	"context"
	"log"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"
)

type StatisticsHandler struct {
	statisticspb.UnimplementedStatisticsServiceServer
	uc *usecase.StatisticsUsecase
}

func NewStatisticsHandler(uc *usecase.StatisticsUsecase) *StatisticsHandler {
	return &StatisticsHandler{
		uc: uc,
	}
}

func (h *StatisticsHandler) GetUserOrdersStatistics(ctx context.Context, req *statisticspb.UserOrderStatisticsRequest) (*statisticspb.UserOrderStatisticsResponse, error) {
	log.Printf("[gRPC] GetUserOrdersStatistics called for user_id=%s", req.UserId)

	resp, err := h.uc.GetUserOrdersStatistics(ctx, req.UserId)
	if err != nil {
		log.Printf("[gRPC] GetUserOrdersStatistics error: %v", err)
		return nil, err
	}
	log.Printf("[gRPC] GetUserOrdersStatistics result: total_orders=%d", resp.TotalOrders)
	return &statisticspb.UserOrderStatisticsResponse{TotalOrders: resp.TotalOrders}, nil
}

func (h *StatisticsHandler) GetUserStatistics(ctx context.Context, req *statisticspb.UserStatisticsRequest) (*statisticspb.UserStatisticsResponse, error) {
	log.Println("[gRPC] GetUserStatistics called")

	resp, err := h.uc.GetUserStatistics(ctx)
	if err != nil {
		log.Printf("[gRPC] GetUserStatistics error: %v", err)
		return nil, err
	}
	log.Printf("[gRPC] GetUserStatistics result: total_users=%d, daily_active_users=%d",
		resp.TotalUsers, resp.DailyActiveUsers)
	return &statisticspb.UserStatisticsResponse{
		TotalUsers:       resp.TotalUsers,
		DailyActiveUsers: resp.DailyActiveUsers,
	}, nil
}
