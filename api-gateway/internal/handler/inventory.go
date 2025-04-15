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
