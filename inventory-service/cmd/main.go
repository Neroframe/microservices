package main

import (
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/config"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/handler"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/repository"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// utils.LoggerInit()
	// utils.Log.Info("Inventory service started")

	db := config.ConnectToMongo()

	productRepo := repository.NewProductMongoRepo(db)

	productUsecase := usecase.NewProductUsecase(productRepo)

	handler := handler.NewProductHandler(productUsecase)

	r := gin.Default()
	r.GET("/products/:id", handler.GetProduct)
	r.POST("/products/", handler.CreateProduct)
	r.PATCH("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)
	r.GET("/products/", handler.ListProducts)

	r.Run(":8081")
}
