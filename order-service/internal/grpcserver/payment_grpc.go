package grpcserver

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
	orderpb "github.com/Neroframe/ecommerce-platform/order-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OrderGRPCServer) CreatePayment(ctx context.Context, req *orderpb.CreatePaymentRequest) (*orderpb.PaymentResponse, error) {
	payment := &domain.Payment{
		OrderID:       req.OrderId,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
	}

	if err := s.paymentUsecase.Create(ctx, payment); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create payment: %v", err)
	}

	return &orderpb.PaymentResponse{
		PaymentId: payment.ID,
		Status:    payment.Status,
		Message:   "Payment created successfully",
	}, nil
}

func (s *OrderGRPCServer) GetPaymentByID(ctx context.Context, req *orderpb.GetPaymentRequest) (*orderpb.PaymentResponse, error) {
	payment, err := s.paymentUsecase.GetByID(ctx, req.PaymentId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "payment not found: %v", err)
	}

	return &orderpb.PaymentResponse{
		PaymentId: payment.ID,
		Status:    payment.Status,
		Message:   "Payment retrieved successfully",
	}, nil
}
