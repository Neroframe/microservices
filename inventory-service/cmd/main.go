package main

import (
	"log"
	"net"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/config"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/grpcserver"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/repository"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/usecase"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := config.ConnectToMongo()
	productRepo := repository.NewProductMongoRepo(db)

	productUsecase := usecase.NewProductUsecase(productRepo)
	
	s := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(s, grpcserver.NewInventoryGRPCServer(productUsecase))

	log.Println("gRPC server running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
