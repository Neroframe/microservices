package grpcserver

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/usecase"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *InventoryGRPCServer) CreateProduct(ctx context.Context, req *inventorypb.CreateProductRequest) (*inventorypb.ProductResponse, error) {
	product := &domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
		Stock:    int(req.Stock),
	}

	if err := s.usecase.Create(ctx, product); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}

	return &inventorypb.ProductResponse{
		Id:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category,
		Stock:    int32(product.Stock),
	}, nil
}

func (s *InventoryGRPCServer) UpdateProduct(ctx context.Context, req *inventorypb.UpdateProductRequest) (*inventorypb.ProductResponse, error) {
	current, err := s.usecase.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	if req.Name != "" {
		current.Name = req.Name
	}
	if req.Price != 0 {
		current.Price = req.Price
	}
	if req.Category != "" {
		current.Category = req.Category
	}
	if req.Stock != 0 {
		current.Stock = int(req.Stock)
	}

	if err := s.usecase.Update(ctx, current); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update product: %v", err)
	}

	return &inventorypb.ProductResponse{
		Id:       current.ID,
		Name:     current.Name,
		Price:    current.Price,
		Category: current.Category,
		Stock:    int32(current.Stock),
	}, nil
}

func (s *InventoryGRPCServer) DeleteProduct(ctx context.Context, req *inventorypb.DeleteProductRequest) (*emptypb.Empty, error) {
	if err := s.usecase.Delete(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete product: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *InventoryGRPCServer) ListProducts(ctx context.Context, _ *inventorypb.ListProductsRequest) (*inventorypb.ListProductsResponse, error) {
	products, err := s.usecase.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
	}

	var result []*inventorypb.ProductResponse
	for _, p := range products {
		result = append(result, &inventorypb.ProductResponse{
			Id:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Category: p.Category,
			Stock:    int32(p.Stock),
		})
	}

	return &inventorypb.ListProductsResponse{Products: result}, nil
}
