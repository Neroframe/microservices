package main

import (
	"log"
	"net"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/config.go"
	"github.com/Neroframe/ecommerce-platform/order-service/internal/grpcserver"
	"github.com/Neroframe/ecommerce-platform/order-service/internal/repository"
	"github.com/Neroframe/ecommerce-platform/order-service/internal/usecase"
	orderpb "github.com/Neroframe/ecommerce-platform/order-service/proto"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Order Service: Starting up")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Order Service: Failed to listen: %v", err)
	}

	db := config.ConnectToMongo()
	orderRepo := repository.NewOrderMongoRepo(db)
	paymentRepo := repository.NewPaymentMongoRepo(db)

	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo)

	s := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(s, grpcserver.NewOrderGRPCServer(orderUsecase, paymentUsecase))
	orderpb.RegisterPaymentServiceServer(s, grpcserver.NewOrderGRPCServer(orderUsecase, paymentUsecase))

	log.Println("Order Service: gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Order Service: Failed to serve: %v", err)
	}
}
