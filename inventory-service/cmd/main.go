package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/config"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/grpcserver"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/repository"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/usecase"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/utils"
	redisclient "github.com/Neroframe/ecommerce-platform/inventory-service/pkg/redis"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// init logger
	utils.InitLogger()
	utils.Log.Info("starting inventory-service…")

	ctx := context.Background()

	// connect Redis (using REDIS_HOSTS env var)
	redisHost := os.Getenv("REDIS_HOSTS")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}
	rCfg := redisclient.Config{
		Host:         redisHost,
		Password:     "",
		TLSEnable:    false,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	rCli, err := redisclient.NewClient(ctx, rCfg)
	if err != nil {
		utils.Log.Error("redis connect failed", "err", err)
		os.Exit(1)
	}
	defer func() {
		if err := rCli.Close(); err != nil {
			utils.Log.Error("redis close failed", "err", err)
		}
	}()
	utils.Log.Info("redis ping OK")

	// connect MongoDB
	db := config.ConnectToMongo()
	prodRepo := repository.NewProductMongoRepo(db)
	catRepo := repository.NewCategoryMongoRepo(db)

	// init usecases with cache TTL
	cacheTTL := 12 * time.Hour
	productUC := usecase.NewProductUsecase(prodRepo, rCli.Unwrap(), cacheTTL)
	categoryUC := usecase.NewCategoryUsecase(catRepo)

	// initialize cache at startup
	if err := productUC.RefreshProductsCache(ctx); err != nil {
		utils.Log.Error("cache init failed", "err", err)
		os.Exit(1)
	}

	// schedule periodic cache refresh every 12h
	go func() {
		ticker := time.NewTicker(cacheTTL)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := productUC.RefreshProductsCache(ctx); err != nil {
					utils.Log.Error("cache refresh failed", "err", err)
				} else {
					utils.Log.Info("cache refreshed")
				}
			}
		}
	}()

	// start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		utils.Log.Error("listen failed", "err", err)
		os.Exit(1)
	}
	grpcSrv := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(grpcSrv,
		grpcserver.NewInventoryGRPCServer(productUC, categoryUC),
	)

	utils.Log.Info("grpc server running on :50051")
	if err := grpcSrv.Serve(lis); err != nil {
		utils.Log.Error("serve failed", "err", err)
		os.Exit(1)
	}
}
