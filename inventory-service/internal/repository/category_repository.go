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

type mongoCategoryRepo struct {
	collection *mongo.Collection
}

func NewCategoryMongoRepo(db *mongo.Database) domain.CategoryRepository {
	return &mongoCategoryRepo{collection: db.Collection("categories")}
}

func (r *mongoCategoryRepo) Create(ctx context.Context, c *domain.Category) error {
	res, err := r.collection.InsertOne(ctx, c)
	if err != nil {
		utils.Log.Error("Insert category failed", "err", err)
		return err
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		c.ID = oid.Hex()
		utils.Log.Info("Category inserted", "id", c.ID)
	}

	return nil
}

func (r *mongoCategoryRepo) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid category ID format")
	}

	var c domain.Category
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&c)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	c.ID = id
	return &c, nil
}

func (r *mongoCategoryRepo) Update(ctx context.Context, c *domain.Category) error {
	oid, err := primitive.ObjectIDFromHex(c.ID)
	if err != nil {
		return errors.New("invalid category ID format")
	}

	update := bson.M{"name": c.Name}

	_, err = r.collection.UpdateByID(ctx, oid, bson.M{"$set": update})
	if err != nil {
		utils.Log.Error("Update category failed", "id", c.ID, "err", err)
	}
	return err
}

func (r *mongoCategoryRepo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid category ID format")
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (r *mongoCategoryRepo) List(ctx context.Context) ([]*domain.Category, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []*domain.Category
	for cursor.Next(ctx) {
		var c domain.Category
		if err := cursor.Decode(&c); err != nil {
			return nil, err
		}
		categories = append(categories, &c)
	}

	return categories, nil
}
