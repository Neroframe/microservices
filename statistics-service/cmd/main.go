package main

import (
	"context"
	"log"

	"github.com/Neroframe/ecommerce-platform/statistics-service/config"
	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/app"
)

func main() {
	// 1. Load config
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	// 2. Bootstrap the whole application
	application, err := app.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("app init error: %v", err)
	}

	// 3. Run (blocks until shutdown)
	if err := application.Run(); err != nil {
		log.Fatalf("service error: %v", err)
	}
}
