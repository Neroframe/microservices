package main

import (
	"log"
	"time"

	"github.com/Neroframe/ecommerce-platform/api-gateway/internal/client"
	"github.com/Neroframe/ecommerce-platform/api-gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatal("failed to connect to inventory service:", err)
	}
	defer conn.Close()

	client.InitInventoryClient(conn)

	r := gin.Default()

	api := r.Group("/v1/inventory")
	{
		api.GET("/product/:id", handler.GetProductByID)
		api.POST("/product", handler.CreateProduct)
		api.PUT("/product", handler.UpdateProduct)
		api.DELETE("/product/:id", handler.DeleteProduct)
		api.GET("/products", handler.ListProducts)

		api.GET("/category/:id", handler.GetCategoryByID)
		api.POST("/category", handler.CreateCategory)
		api.PUT("/category", handler.UpdateCategory)
		api.DELETE("/category/:id", handler.DeleteCategory)
		api.GET("/categories", handler.ListCategories)
	}

	log.Println("API Gateway running on :8080")
	r.Run(":8080")
}
