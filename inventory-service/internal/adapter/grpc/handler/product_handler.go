package handler

import (
	"context"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/utils"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *InventoryHandler) GetProductByID(ctx context.Context, req *inventorypb.GetProductRequest) (*inventorypb.ProductResponse, error) {
	// utils.Log.Info("gRPC GetProductByID", "id", req.Id)

	product, err := s.productUsecase.GetByID(ctx, req.Id)
	if err != nil {
		utils.Log.Error("product not found", "id", req.Id, "err", err)
		return nil, status.Errorf(codes.NotFound, "product not found: %v", err)
	}

	if product == nil {
		utils.Log.Warn("Product returned nil", "id", req.Id)
		return nil, status.Errorf(codes.NotFound, "product not found")
	}

	utils.Log.Info("product found", "id", product.ID)
	return &inventorypb.ProductResponse{
		Id:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category,
		Stock:    int32(product.Stock),
	}, nil
}

func (s *InventoryHandler) CreateProduct(ctx context.Context, req *inventorypb.CreateProductRequest) (*inventorypb.ProductResponse, error) {
	// utils.Log.Info("gRPC CreateProduct", "name", req.Name)

	product := &domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
		Stock:    int(req.Stock),
	}

	if err := s.productUsecase.Create(ctx, product); err != nil {
		utils.Log.Error("failed to create product", "err", err)
		return nil, status.Errorf(codes.Internal, "failed to create product: %v", err)
	}

	utils.Log.Info("product created", "id", product.ID)
	return &inventorypb.ProductResponse{
		Id:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category,
		Stock:    int32(product.Stock),
	}, nil
}

func (s *InventoryHandler) UpdateProduct(ctx context.Context, req *inventorypb.UpdateProductRequest) (*inventorypb.ProductResponse, error) {
	// utils.Log.Info("gRPC UpdateProduct", "id", req.Id)

	current, err := s.productUsecase.GetByID(ctx, req.Id)
	if err != nil {
		utils.Log.Error("product not found for update", "id", req.Id, "err", err)
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

	if err := s.productUsecase.Update(ctx, current); err != nil {
		utils.Log.Error("failed to update product", "id", current.ID, "err", err)
		return nil, status.Errorf(codes.Internal, "failed to update product: %v", err)
	}

	utils.Log.Info("product updated", "id", current.ID)
	return &inventorypb.ProductResponse{
		Id:       current.ID,
		Name:     current.Name,
		Price:    current.Price,
		Category: current.Category,
		Stock:    int32(current.Stock),
	}, nil
}

func (s *InventoryHandler) DeleteProduct(ctx context.Context, req *inventorypb.DeleteProductRequest) (*emptypb.Empty, error) {
	// utils.Log.Info("gRPC DeleteProduct", "id", req.Id)

	if err := s.productUsecase.Delete(ctx, req.Id); err != nil {
		utils.Log.Error("failed to delete product", "id", req.Id, "err", err)
		return nil, status.Errorf(codes.Internal, "failed to delete product: %v", err)
	}

	utils.Log.Info("product deleted", "id", req.Id)
	return &emptypb.Empty{}, nil
}

func (s *InventoryHandler) ListProducts(ctx context.Context, _ *inventorypb.ListProductsRequest) (*inventorypb.ListProductsResponse, error) {
	// utils.Log.Info("gRPC ListProducts")

	products, err := s.productUsecase.List(ctx)
	if err != nil {
		utils.Log.Error("failed to list products", "err", err)
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
	}

	utils.Log.Info("products listed", "count", len(products))

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
