package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Neroframe/ecommerce-platform/order-service/config"
	grpcadapter "github.com/Neroframe/ecommerce-platform/order-service/internal/adapter/grpc"
	mongoadapter "github.com/Neroframe/ecommerce-platform/order-service/internal/adapter/mongo"
	natsadapter "github.com/Neroframe/ecommerce-platform/order-service/internal/adapter/nats"
	"github.com/Neroframe/ecommerce-platform/order-service/internal/usecase"

	mongoconn "github.com/Neroframe/ecommerce-platform/order-service/pkg/mongo"
	natsconn "github.com/Neroframe/ecommerce-platform/order-service/pkg/nats"
)

const serviceName = "order-service"

type App struct {
	grpcServer *grpcadapter.API
	// natsConsumer *natsconsumer.PubSub
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Printf("Starting %s...", serviceName)

	mongoDB, err := mongoconn.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	orderRepo := mongoadapter.NewOrderRepository(mongoDB.Conn)
	paymentRepo := mongoadapter.NewPaymentRepository(mongoDB.Conn)

	// NATS client
	natsClient, err := natsconn.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		return nil, fmt.Errorf("nats connect: %w", err)
	}
	log.Printf("NATS status: %s", natsClient.Conn.Status())

	// NATS publisher
	eventPublisher := natsadapter.NewOrderEventPublisher(natsClient)

	orderUC := usecase.NewOrderUsecase(orderRepo, eventPublisher)
	paymentUC := usecase.NewPaymentUsecase(paymentRepo)

	grpcAPI := grpcadapter.New(cfg.Server.GRPCServer, orderUC, paymentUC)

	return &App{grpcServer: grpcAPI}, nil
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	a.grpcServer.Run(ctx, errCh)
	log.Println("Order service is running")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return fmt.Errorf("runtime error: %w", err)
	case sig := <-sigCh:
		log.Printf("Received signal %v, shutting down...", sig)
		if cerr := a.grpcServer.Stop(ctx); cerr != nil {
			log.Printf("gRPC stop error: %v", cerr)
		}
		return nil
	}
}
