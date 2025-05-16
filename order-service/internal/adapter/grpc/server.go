package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Neroframe/ecommerce-platform/order-service/config"
	"github.com/Neroframe/ecommerce-platform/order-service/internal/adapter/grpc/handler"
	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
	orderpb "github.com/Neroframe/ecommerce-platform/order-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type API struct {
	server    *grpc.Server
	cfg       config.GRPCServer
	orderUC   domain.OrderUsecase
	paymentUC domain.PaymentUsecase
	addr      string
}

func New(cfg config.GRPCServer, ou domain.OrderUsecase, pu domain.PaymentUsecase) *API {
	return &API{
		cfg:       cfg,
		orderUC:   ou,
		paymentUC: pu,
		addr:      fmt.Sprintf("0.0.0.0:%d", cfg.Port),
	}
}

func (api *API) Run(ctx context.Context, errCh chan<- error) {
	go func() {
		log.Printf("gRPC server starting on %s", api.addr)
		if err := api.run(ctx); err != nil {
			errCh <- fmt.Errorf("error starting grpc server: %w", err)
		}
	}()
}

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
	case <-ctx.Done():
		api.server.Stop()
	case <-done:
	}
	return nil
}

func (api *API) run(ctx context.Context) error {
	opts := api.setOptions(ctx)
	api.server = grpc.NewServer(opts...)

	// register order service
	oh := handler.NewOrderHandler(api.orderUC)
	orderpb.RegisterOrderServiceServer(api.server, oh)

	// register payment service
	ph := handler.NewPaymentHandler(api.paymentUC)
	orderpb.RegisterPaymentServiceServer(api.server, ph)

	reflection.Register(api.server)

	lis, err := net.Listen("tcp", api.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", api.addr, err)
	}
	return api.server.Serve(lis)
}

func (api *API) setOptions(ctx context.Context) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      api.cfg.MaxConnectionAge,
			MaxConnectionAgeGrace: api.cfg.MaxConnectionAgeGrace,
		}),
		grpc.MaxRecvMsgSize(int(api.cfg.MaxRecvMsgSizeMiB) * 1024 * 1024), // MaxRecvSize * 1 MB
	}
}
