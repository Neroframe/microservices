package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Neroframe/ecommerce-platform/statistics-service/config"
	grpcadapter "github.com/Neroframe/ecommerce-platform/statistics-service/internal/adapter/grpc"
	cache "github.com/Neroframe/ecommerce-platform/statistics-service/internal/adapter/inmemory"
	mongoadapter "github.com/Neroframe/ecommerce-platform/statistics-service/internal/adapter/mongo"
	natsadapter "github.com/Neroframe/ecommerce-platform/statistics-service/internal/adapter/nats"
	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	mongocon "github.com/Neroframe/ecommerce-platform/statistics-service/pkg/mongo"
	natsconn "github.com/Neroframe/ecommerce-platform/statistics-service/pkg/nats"
	natsconsumer "github.com/Neroframe/ecommerce-platform/statistics-service/pkg/nats/consumer"
)

const serviceName = "statistics-service"

type App struct {
	grpcServer   *grpcadapter.API
	natsConsumer *natsconsumer.PubSub
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Printf("Starting %s...", serviceName)

	// MongoDB
	mdb, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo connect: %w", err)
	}

	// Repository, inmemory cache & Usecase
	repo := mongoadapter.NewRepository(mdb.Conn)
	evtCache := cache.NewInMemoryEventCache()

	uc := usecase.NewStatisticsUsecase(repo, evtCache)

	// gRPC API
	grpcAPI := grpcadapter.New(cfg.Server.GRPCServer, uc)

	// NATS Client
	nc, err := natsconn.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		return nil, fmt.Errorf("nats connect: %w", err)
	}
	log.Printf("NATS status: %s", nc.Conn.Status())

	// NATS conn & Handler
	pubsub := natsconsumer.NewPubSub(nc)
	handler := natsadapter.NewStatisticsHandler(uc, nc.Conn)

	subjects := []string{
		cfg.Nats.NatsSubjects.OrderCreated,
		cfg.Nats.NatsSubjects.OrderUpdated,
		cfg.Nats.NatsSubjects.OrderDeleted,
		cfg.Nats.NatsSubjects.ProductCreated,
		cfg.Nats.NatsSubjects.ProductUpdated,
		cfg.Nats.NatsSubjects.ProductDeleted,
	}

	for _, subj := range subjects {
		pubsub.Subscribe(natsconsumer.PubSubSubscriptionConfig{
			Subject: subj,
			Handler: handler.Handle,
		})
	}

	return &App{
		grpcServer:   grpcAPI,
		natsConsumer: pubsub,
	}, nil
}

// starts the gRPC server and NATS consumer - blocking until an error or OS signal
func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start servers
	a.grpcServer.Run(ctx, errCh)
	a.natsConsumer.Start(ctx, errCh)
	log.Println("Statistics service is running")

	// handle termination
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		return fmt.Errorf("runtime error: %w", err)

	case sig := <-sigCh:
		log.Printf("Received signal %v, shutting down...", sig)
		// graceful shutdown
		if cerr := a.grpcServer.Stop(ctx); cerr != nil {
			log.Printf("gRPC stop error: %v", cerr)
		}
		a.natsConsumer.Stop()
		return nil
	}
}
