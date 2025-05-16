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
	// Connect to microservices
	conn, err := grpc.Dial("inventory-service:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatal("failed to connect to inventory service:", err)
	}
	defer conn.Close()

	orderConn, err := grpc.Dial("order-service:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatal("failed to connect to order service:", err)
	}
	defer orderConn.Close()

	statsConn, err := grpc.Dial("statistics-service:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatal("failed to connect to statistics service:", err)
	}
	defer statsConn.Close()

	// Init microservices
	client.InitInventoryClient(conn)
	client.InitOrderClient(orderConn)
	client.InitStatisticsClient(statsConn)

	r := gin.Default()
	api := r.Group("/v1")
	{
		inventory := api.Group("/inventory")
		{
			inventory.GET("/product/:id", handler.GetProductByID)
			inventory.POST("/product", handler.CreateProduct)
			inventory.PUT("/product", handler.UpdateProduct)
			inventory.DELETE("/product/:id", handler.DeleteProduct)
			inventory.GET("/products", handler.ListProducts)

			inventory.GET("/category/:id", handler.GetCategoryByID)
			inventory.POST("/category", handler.CreateCategory)
			inventory.PUT("/category", handler.UpdateCategory)
			inventory.DELETE("/category/:id", handler.DeleteCategory)
			inventory.GET("/categories", handler.ListCategories)
		}

		orders := api.Group("/orders")
		{
			orders.POST("/", handler.CreateOrder)
			orders.GET("/:id", handler.GetOrderByID)
			orders.PUT("/:id/status", handler.UpdateOrderStatus)
			orders.GET("/user/:userId", handler.ListUserOrders)
		}

		payments := api.Group("/payments")
		{
			payments.POST("/", handler.CreatePayment)
			payments.GET("/:id", handler.GetPaymentByID)
		}

		statistics := api.Group("/statistics")
		{
			statistics.GET("/user/:userId/orders", handler.GetUserOrdersStatistics)
			statistics.GET("/users", handler.GetUserStatistics)
		}
	}

	log.Println("API Gateway running on :8080")
	r.Run(":8080")
}
