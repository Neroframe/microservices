package handler

import (
	"log"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/handler/dto"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/usecase"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(u usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	product, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": "product not found"})
		return
	}

	c.JSON(200, product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	log.Println("[Handler] Incoming POST /products")
	// dto validation
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[Handler] JSON binding failed: %v", err)
		c.JSON(400, gin.H{"error": "invalid payload"})
		return
	}

	// transform to domain
	product := &domain.Product{
		Name:     req.Name,
		Price:    req.Price,
		Category: req.Category,
		Stock:    req.Stock,
	}

	// send to usecase
	if err := h.usecase.Create(c.Request.Context(), product); err != nil {
		c.JSON(500, gin.H{"error": "failed to create product--"})
		return
	}

	c.JSON(201, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid payload"})
		return
	}

	current, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": "product not found"})
		return
	}

	product := &domain.Product{
		ID:       id,
		Name:     utils.DefaultIfNilString(req.Name, current.Name),
		Price:    utils.DefaultIfNilFloat(req.Price, current.Price),
		Category: utils.DefaultIfNilString(req.Category, current.Category),
		Stock:    utils.DefaultIfNilInt(req.Stock, current.Stock),
	}

	if err := h.usecase.Update(c.Request.Context(), product); err != nil {
		c.JSON(500, gin.H{"error": "failed to update product"})
		return
	}

	c.JSON(200, gin.H{"message": "product updated"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {

}

func (h *ProductHandler) ListProducts(c *gin.Context) {

}
