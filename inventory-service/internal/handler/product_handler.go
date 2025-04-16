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
		log.Printf("Product not found: %v", err)
		c.JSON(404, gin.H{"error": "product not found"})
		return
	}

	log.Printf("Fetched product: %+v\n", product)
	c.JSON(200, gin.H{"message": "product found"})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	// dto validation
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("JSON binding failed: %v", err)
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

	log.Printf("Fetched product: %+v\n", product)
	c.JSON(201, gin.H{"message": "product created"})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Log.Error("JSON binding failed", "err", err)
		c.JSON(400, gin.H{"error": "invalid payload"})
		return
	}

	current, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		utils.Log.Error("product not found", "id", id, "err", err)
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
		utils.Log.Error("Failed to update product", "id", id, "err", err)
		c.JSON(500, gin.H{"error": "failed to update product"})
		return
	}

	utils.Log.Info("Product updated", "p", product)
	c.JSON(200, gin.H{"message": "product updated"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := h.usecase.Delete(c.Request.Context(), id); err != nil {
		log.Printf("Failed to delete product %s: %v", id, err)
		c.JSON(500, gin.H{"error": "failed to delete product"})
		return
	}

	log.Printf("Deleted product ID: %s\n", id)
	c.JSON(200, gin.H{"message": "product deleted"})
}

func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.usecase.List(c.Request.Context())
	if err != nil {
		log.Printf("Failed to list products: %v", err)
		c.JSON(500, gin.H{"error": "failed to list products"})
		return
	}

	log.Printf("Listed %d products", len(products))
	c.JSON(200, products)
}
