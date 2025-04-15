package grpcserver

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/usecase"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InventoryGRPCServer struct {
	usecase usecase.ProductUsecase
	inventorypb.UnimplementedInventoryServiceServer
}

func NewInventoryGRPCServer(u usecase.ProductUsecase) *InventoryGRPCServer {
	return &InventoryGRPCServer{
		usecase: u,
	}
}

func (s *InventoryGRPCServer) GetProductByID(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.ProductResponse, error) {
	product, err := s.usecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	return &inventorypb.ProductResponse{
		Id:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category,
		Stock:    int32(product.Stock),
	}, nil
}
