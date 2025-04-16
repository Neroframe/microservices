package grpcserver

import (
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"
)

type InventoryGRPCServer struct {
	categoryUsecase domain.CategoryUsecase
	productUsecase  domain.ProductUsecase
	inventorypb.UnimplementedInventoryServiceServer
}

func NewInventoryGRPCServer(
	productUsecase domain.ProductUsecase,
	categoryUsecase domain.CategoryUsecase,
) *InventoryGRPCServer {
	return &InventoryGRPCServer{
		productUsecase:  productUsecase,
		categoryUsecase: categoryUsecase,
	}
}
