package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Neroframe/ecommerce-platform/inventory-service/config"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/adapter/grpc/handler"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type API struct {
	server     *grpc.Server
	cfg        config.GRPCServer
	productUC  domain.ProductUsecase
	categoryUC domain.CategoryUsecase
	addr       string
}

func New(cfg config.GRPCServer, pu domain.ProductUsecase, cu domain.CategoryUsecase) *API {
	return &API{
		cfg:        cfg,
		productUC:  pu,
		categoryUC: cu,
		addr:       fmt.Sprintf("0.0.0.0:%d", cfg.Port),
	}
}

func (api *API) Run(ctx context.Context, errCh chan<- error) {
	go func() {
		log.Printf("gRPC server starting on %s", api.addr)
		if err := api.run(ctx); err != nil {
			errCh <- fmt.Errorf("error starting gRPC server: %w", err)
		}
	}()
}

// Provide context to force stop on timeout
func (api *API) Stop(ctx context.Context) error {
	if api.server == nil {
		return nil
	}

	done := make(chan struct{})
	go func() {
		api.server.GracefulStop()
		close(done)
	}()

	select {
	case <-ctx.Done(): // Stop immediately if the context is terminated
		api.server.Stop()
	case <-done:
	}

	return nil
}

func (api *API) run(ctx context.Context) error {
	// build server opts
	opts := api.setOptions(ctx)
	api.server = grpc.NewServer(opts...)

	// register product, category service
	InventoryHandler := handler.NewInventoryHandler(api.productUC, api.categoryUC)
	inventorypb.RegisterInventoryServiceServer(api.server, InventoryHandler)

	reflection.Register(api.server)

	lis, err := net.Listen("tcp", api.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", api.addr, err)
	}

	return api.server.Serve(lis)
}

func (a *API) setOptions(ctx context.Context) []grpc.ServerOption {
	opts := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      a.cfg.MaxConnectionAge,
			MaxConnectionAgeGrace: a.cfg.MaxConnectionAgeGrace,
		}),
		grpc.MaxRecvMsgSize(a.cfg.MaxRecvMsgSizeMiB * 1024 * 1024), // MaxRecvSize * 1 MB
	}

	return opts
}
