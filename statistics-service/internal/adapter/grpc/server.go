package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Neroframe/ecommerce-platform/statistics-service/config"
	handler "github.com/Neroframe/ecommerce-platform/statistics-service/internal/adapter/grpc/handler"
	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	statisticspb "github.com/Neroframe/ecommerce-platform/statistics-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// API wraps the gRPC server for the statistics service.
type API struct {
	s    *grpc.Server
	cfg  config.GRPCServer
	uc   usecase.StatisticsUsecase
	addr string
}

// New creates a new gRPC API with the given configuration and usecase.
func New(cfg config.GRPCServer, uc usecase.StatisticsUsecase) *API {
	return &API{
		cfg:  cfg,
		uc:   uc,
		addr: fmt.Sprintf("0.0.0.0:%d", cfg.Port),
	}
}

// Run starts the gRPC server asynchronously, reporting startup errors on errCh.
func (a *API) Run(ctx context.Context, errCh chan<- error) {
	go func() {
		log.Printf("gRPC server starting on %s", a.addr)
		if err := a.run(ctx); err != nil {
			errCh <- fmt.Errorf("can't start grpc server: %w", err)
		}
	}()
}

// Stop gracefully stops the gRPC server. Blocks until done or ctx is cancelled.
func (a *API) Stop(ctx context.Context) error {
	if a.s == nil {
		return nil
	}
	done := make(chan struct{})
	go func() {
		a.s.GracefulStop()
		close(done)
	}()

	select {
	case <-ctx.Done():
		a.s.Stop()
	case <-done:
	}
	return nil
}

// run configures and starts the underlying gRPC Server.
func (a *API) run(ctx context.Context) error {
	// build server options
	opts := a.setOptions(ctx)
	a.s = grpc.NewServer(opts...)

	// register service handler
	h := handler.NewStatisticsHandler(a.uc)
	statisticspb.RegisterStatisticsServiceServer(a.s, h)

	// register reflection for debugging
	reflection.Register(a.s)

	// listen for requests
	ln, err := net.Listen("tcp", a.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", a.addr, err)
	}

	// serve incoming connections
	if err := a.s.Serve(ln); err != nil {
		return fmt.Errorf("failed to serve grpc: %w", err)
	}
	return nil
}

// setOptions constructs the gRPC server options from config.
func (a *API) setOptions(ctx context.Context) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      a.cfg.MaxConnectionAge,
			MaxConnectionAgeGrace: a.cfg.MaxConnectionAgeGrace,
		}),
		grpc.MaxRecvMsgSize(int(a.cfg.MaxRecvMsgSizeMiB) * 1024 * 1024),
	}
}
