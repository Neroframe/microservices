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
	r.GET("/products/:id", handler.GetProductByID)

	log.Println("API Gateway running on :8080")
	r.Run(":8080")
}

// func main() {
// 	r := gin.Default()

// 	inventoryURL, _ := url.Parse("http://inventory-service:8081")
// 	inventoryProxy := httputil.NewSingleHostReverseProxy(inventoryURL)
// 	r.Any("/products/*proxyPath", func(c *gin.Context) {
// 		inventoryProxy.ServeHTTP(c.Writer, c.Request)
// 	})

// 	orderURL, _ := url.Parse("http://order-service:8082")
// 	orderProxy := httputil.NewSingleHostReverseProxy(orderURL)
// 	r.Any("/orders/*proxyPath", func(c *gin.Context) {
// 		orderProxy.ServeHTTP(c.Writer, c.Request)
// 	})

// 	r.Run(":8080")
// }
