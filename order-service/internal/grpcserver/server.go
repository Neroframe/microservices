package grpcserver

import (
	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
	orderpb "github.com/Neroframe/ecommerce-platform/order-service/proto"
)

type OrderGRPCServer struct {
	orderUsecase   domain.OrderUsecase
	paymentUsecase domain.PaymentUsecase
	orderpb.UnimplementedOrderServiceServer
	orderpb.UnimplementedPaymentServiceServer
}

func NewOrderGRPCServer(
	orderUsecase domain.OrderUsecase,
	paymentUsecase domain.PaymentUsecase,
) *OrderGRPCServer {
	return &OrderGRPCServer{
		orderUsecase:   orderUsecase,
		paymentUsecase: paymentUsecase,
	}
}
