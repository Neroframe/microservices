package repository

import (
	"context"
	"errors"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoProductRepo struct {
	collection *mongo.Collection
}

func NewProductMongoRepo(db *mongo.Database) domain.ProductRepository {
	return &mongoProductRepo{collection: db.Collection("products")}
}

func (r *mongoProductRepo) Create(ctx context.Context, p *domain.Product) error {
	// utils.Log.Info("Creating product", "name", p.Name)

	res, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		utils.Log.Error("InsertOne failed", "err", err)
		return err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		p.ID = oid.Hex()
		utils.Log.Info("Product inserted", "id", p.ID)
	} else {
		utils.Log.Warn("Inserted ID is not ObjectID", "raw_id", res.InsertedID)
	}

	return nil
}

func (r *mongoProductRepo) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	// utils.Log.Info("Fetching product by ID", "id", id)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.Log.Error("Failed to convert product ID", "id", id, "err", err)
		return nil, errors.New("invalid product ID")
	}

	var product domain.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.Log.Warn("Product not found", "id", id)
			return nil, nil
		}
		utils.Log.Error("FindOne failed", "id", id, "err", err)
		return nil, err
	}

	product.ID = id
	return &product, nil
}

func (r *mongoProductRepo) Update(ctx context.Context, p *domain.Product) error {
	// utils.Log.Info("Updating product", "id", p.ID)

	oid, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		utils.Log.Error("Failed to convert product ID", "id", p.ID, "err", err)
		return err
	}

	update := bson.M{
		"name":     p.Name,
		"price":    p.Price,
		"category": p.Category,
		"stock":    p.Stock,
	}

	_, err = r.collection.UpdateByID(ctx, oid, bson.M{"$set": update})
	if err != nil {
		utils.Log.Error("UpdateByID failed", "id", p.ID, "err", err)
		return err
	}

	return nil
}

func (r *mongoProductRepo) Delete(ctx context.Context, id string) error {
	// utils.Log.Info("Deleting product", "id", id)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.Log.Error("Invalid ObjectID for delete", "id", id)
		return errors.New("invalid product ID")
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		utils.Log.Error("DeleteOne failed", "id", id, "err", err)
	}
	return err
}

func (r *mongoProductRepo) List(ctx context.Context) ([]*domain.Product, error) {
	// utils.Log.Info("Listing all products")

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		utils.Log.Error("Find failed", "err", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*domain.Product
	for cursor.Next(ctx) {
		var p domain.Product
		if err := cursor.Decode(&p); err != nil {
			utils.Log.Error("Decode failed", "err", err)
			return nil, err
		}
		products = append(products, &p)
	}

	return products, nil
}
