package grpcserver

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
	orderpb "github.com/Neroframe/ecommerce-platform/order-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderGRPCServer) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.OrderResponse, error) {
	order := &domain.Order{
		UserID: req.UserId,
		Items:  mapOrderItems(req.Items),
	}
	if err := s.orderUsecase.Create(ctx, order); err != nil {
		return nil, status.Errorf(codes.Internal, "create order failed: %v", err)
	}
	return toOrderResponse(order), nil
}

func (s *OrderGRPCServer) GetOrderByID(ctx context.Context, req *orderpb.GetOrderRequest) (*orderpb.OrderResponse, error) {
	order, err := s.orderUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "order not found: %v", err)
	}
	return toOrderResponse(order), nil
}

func (s *OrderGRPCServer) UpdateOrderStatus(ctx context.Context, req *orderpb.UpdateOrderStatusRequest) (*orderpb.OrderResponse, error) {
	order, err := s.orderUsecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "order not found: %v", err)
	}
	if order == nil {
		return nil, status.Errorf(codes.NotFound, "order not found")
	}

	order.Status = req.Status
	if err := s.orderUsecase.Update(ctx, order); err != nil {
		return nil, status.Errorf(codes.Internal, "update failed: %v", err)
	}
	return toOrderResponse(order), nil
}

func (s *OrderGRPCServer) ListUserOrders(ctx context.Context, req *orderpb.ListOrdersRequest) (*orderpb.ListOrdersResponse, error) {
	orders, err := s.orderUsecase.ListByUserID(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list failed: %v", err)
	}
	var resp []*orderpb.OrderResponse
	for _, o := range orders {
		resp = append(resp, toOrderResponse(o))
	}
	return &orderpb.ListOrdersResponse{Orders: resp}, nil
}

func mapOrderItems(items []*orderpb.OrderItem) []domain.OrderItem {
	var result []domain.OrderItem
	for _, i := range items {
		result = append(result, domain.OrderItem{
			ProductID: i.ProductId,
			Quantity:  int(i.Quantity),
		})
	}
	return result
}

func toOrderResponse(o *domain.Order) *orderpb.OrderResponse {
	var items []*orderpb.OrderItem
	for _, i := range o.Items {
		items = append(items, &orderpb.OrderItem{
			ProductId: i.ProductID,
			Quantity:  int32(i.Quantity),
		})
	}
	return &orderpb.OrderResponse{
		Id:     o.ID,
		UserId: o.UserID,
		Status: o.Status,
		Items:  items,
	}
}
