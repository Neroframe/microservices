package main

import (
	"net"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/config"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/grpcserver"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/repository"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/usecase"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/utils"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"
	"google.golang.org/grpc"
)

func main() {
	utils.InitLogger()
	utils.Log.Info("slog started")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		utils.Log.Error("failed to listen", "err", err)
	}

	db := config.ConnectToMongo()
	productRepo := repository.NewProductMongoRepo(db)
	categoryRepo := repository.NewCategoryMongoRepo(db)

	productUsecase := usecase.NewProductUsecase(productRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	s := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(s, grpcserver.NewInventoryGRPCServer(productUsecase, categoryUsecase))
	utils.Log.Info("grpc server running on :50051")
	if err := s.Serve(lis); err != nil {
		utils.Log.Error("failed to serve", "err", err)
	}
}
