package handler

import (
	"context"
	"net/http"

	"github.com/Neroframe/ecommerce-platform/api-gateway/internal/client"
	inventorypb "github.com/Neroframe/ecommerce-platform/api-gateway/proto"

	"github.com/gin-gonic/gin"
)

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	resp, err := client.Inventory.GetProductByID(context.Background(), &inventorypb.GetProductRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func CreateProduct(c *gin.Context) {
	var req inventorypb.CreateProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := client.Inventory.CreateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateProduct(c *gin.Context) {
	var req inventorypb.UpdateProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := client.Inventory.UpdateProduct(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	req := &inventorypb.DeleteProductRequest{
		Id: id,
	}

	_, err := client.Inventory.DeleteProduct(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	c.Status(http.StatusNoContent)
}

func ListProducts(c *gin.Context) {
	req := &inventorypb.ListProductsRequest{}

	resp, err := client.Inventory.ListProducts(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list products"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
