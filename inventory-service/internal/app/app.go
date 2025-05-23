package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcadapter "github.com/Neroframe/ecommerce-platform/inventory-service/internal/adapter/grpc"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/adapter/inmemory"
	mongoadapter "github.com/Neroframe/ecommerce-platform/inventory-service/internal/adapter/mongo"
	natsadapter "github.com/Neroframe/ecommerce-platform/inventory-service/internal/adapter/nats"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/adapter/redis"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/usecase"
	mongoconn "github.com/Neroframe/ecommerce-platform/inventory-service/pkg/mongo"
	natsconn "github.com/Neroframe/ecommerce-platform/inventory-service/pkg/nats"
	redisconn "github.com/Neroframe/ecommerce-platform/inventory-service/pkg/redis"

	"github.com/Neroframe/ecommerce-platform/inventory-service/config"
)

const serviceName = "inventory-service"

type App struct {
	grpcServer *grpcadapter.API
	productUC  domain.ProductUsecase
	// natsConsumer *natsconsumer.PubSub
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Printf("stating %s...", serviceName)

	mongoDB, err := mongoconn.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	productRepo := mongoadapter.NewProductRepository(mongoDB.Conn)
	categoryRepo := mongoadapter.NewCategoryRepository(mongoDB.Conn)

	// NATS client
	natsClient, err := natsconn.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		return nil, fmt.Errorf("nats connect: %w", err)
	}
	log.Printf("NATS status: %s", natsClient.Conn.Status())

	// Redis client
	redisClient, err := redisconn.NewClient(ctx, (redisconn.Config)(cfg.Redis))
	if err != nil {
		return nil, fmt.Errorf("redisconn.NewClient: %w", err)
	}
	log.Println("Redis is connected:", redisClient.Ping(ctx) == nil)

	// Cache inmemory & redis
	productInmemoryCache := inmemory.NewProductCache()
	productRedisCache := redis.NewProductCache(redisClient, cfg.Cache.ProductTTL)

	// NATS publisher
	eventPublisher := natsadapter.NewInventoryEventPublisher(natsClient)

	// UC
	productUC := usecase.NewProductUsecase(productRepo, eventPublisher, productInmemoryCache, productRedisCache)
	categoryUC := usecase.NewCategoryUsecase(categoryRepo, eventPublisher)

	grpcAPI := grpcadapter.New(cfg.Server.GRPCServer, productUC, categoryUC)

	return &App{grpcServer: grpcAPI, productUC: productUC}, nil
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init redis cache warmup
	if err := a.productUC.RefreshProductsCache(ctx); err != nil {
		log.Printf("Error warming up product cache: %v", err)
		return fmt.Errorf("failed to warm up cache: %w", err)
	}

	// Refresh every 12 hour
	go func() {
		ticker := time.NewTicker(10 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := a.productUC.RefreshProductsCache(context.Background()); err != nil {
					log.Printf("Periodic cache refresh failed: %v", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// Start grpc server
	errCh := make(chan error, 1)
	a.grpcServer.Run(ctx, errCh)
	log.Println("Inventory service is running")

	// Handle termination
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return fmt.Errorf("runtime error: %w", err)
	case sig := <-sigCh:
		log.Printf("Received signal %v, shutting down....", sig)
		if cerr := a.grpcServer.Stop(ctx); cerr != nil {
			log.Printf("gRPC stop error: %v", cerr)
		}
		return nil
	}
}
