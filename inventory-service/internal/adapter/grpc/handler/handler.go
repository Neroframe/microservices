package handler

import (
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	inventorypb "github.com/Neroframe/ecommerce-platform/inventory-service/proto"
)

type InventoryHandler struct {
	inventorypb.UnimplementedInventoryServiceServer
	productUsecase domain.ProductUsecase
	categoryUsecase domain.CategoryUsecase
}

func NewInventoryHandler(pu domain.ProductUsecase, cu domain.CategoryUsecase) *InventoryHandler {
	return &InventoryHandler{productUsecase: pu, categoryUsecase: cu,}
}

