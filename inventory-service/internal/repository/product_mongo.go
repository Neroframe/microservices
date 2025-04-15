package repository

import (
	"context"
	"errors"
	"log"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *mongoProductRepo) Create(ctx context.Context, p *domain.Product) error {
	log.Printf("[Repo] Inserting product: %+v\n", p)

	res, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		log.Printf("[Repo] InsertOne failed: %v\n", err)
		return err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		p.ID = oid.Hex()
		log.Printf("[Repo] Inserted with ID: %s\n", p.ID)
	} else {
		log.Printf("[Repo] Inserted but ID type unexpected: %v\n", res.InsertedID)
	}

	return nil
}

func (r *mongoProductRepo) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	var product domain.Product
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	product.ID = id
	return &product, nil
}

func (r *mongoProductRepo) Update(ctx context.Context, p *domain.Product) error {
	update := bson.M{
		"name":     p.Name,
		"price":    p.Price,
		"category": p.Category,
		"stock":    p.Stock,
	}

	_, err := r.collection.UpdateByID(ctx, p.ID, bson.M{"$set": update})
	return err
}

func (r *mongoProductRepo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid product ID")
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (r *mongoProductRepo) List(ctx context.Context) ([]*domain.Product, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*domain.Product
	for cursor.Next(ctx) {
		var p domain.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}
