package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	log.Println("Connecting to MongoDB...")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Printf("Mongo connection setup failed %v\n", err)
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Printf("Mongo ping failed — DB is unreachable %v\n", err)
		panic(err)
	}

	log.Println("MongoDB connection successful")
	return client.Database("inventory_db")
}
