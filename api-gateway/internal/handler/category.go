package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Neroframe/ecommerce-platform/api-gateway/internal/client"
	inventorypb "github.com/Neroframe/ecommerce-platform/api-gateway/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

func CreateCategory(c *gin.Context) {
	var req inventorypb.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding category data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category data"})
		return
	}

	log.Printf("Creating category with name: %s", req.Name)

	resp, err := client.Inventory.CreateCategory(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error creating category: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	log.Printf("Fetching category by ID: %s", id)

	resp, err := client.Inventory.GetCategoryByID(context.Background(), &inventorypb.GetCategoryRequest{
		Id: id,
	})
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error fetching category: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func UpdateCategory(c *gin.Context) {
	var req inventorypb.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding category update data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category data"})
		return
	}

	log.Printf("Updating category with ID: %s", req.Id)

	resp, err := client.Inventory.UpdateCategory(context.Background(), &req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error updating category: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	log.Printf("Deleting category with ID: %s", id)

	_, err := client.Inventory.DeleteCategory(context.Background(), &inventorypb.DeleteCategoryRequest{
		Id: id,
	})
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error deleting category: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func ListCategories(c *gin.Context) {
	log.Println("Listing all categories")

	resp, err := client.Inventory.ListCategories(context.Background(), &inventorypb.ListCategoriesRequest{})
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("Error listing categories: %v", st.Message())
		c.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
