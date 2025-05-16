package main

import (
	"context"
	"log"

	"github.com/Neroframe/ecommerce-platform/inventory-service/config"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/app"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/utils"
)

func main() {
	utils.InitLogger()
	utils.Log.Info("starting inventory-service…")

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	application, err := app.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("app init error: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("service error: %v", err)
	}
}
