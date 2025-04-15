package repository

import (
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoProductRepo struct {
	collection *mongo.Collection
}

func NewProductMongoRepo(db *mongo.Database) domain.ProductRepository {
	return &mongoProductRepo{collection: db.Collection("products")}
}
