package config

import (
	"context"
	"time"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	utils.Log.Info("Connecting to MongoDB...", "uri", "mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		utils.Log.Error("Mongo connection setup failed", "err", err)
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		utils.Log.Error("Mongo ping failed — DB is unreachable", "err", err)
		panic(err)
	}

	utils.Log.Info("MongoDB connection successful")
	return client.Database("inventory_db")
}
