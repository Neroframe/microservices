package mongo

import (
	"context"
	"fmt"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *OrderRepository) Create(ctx context.Context, o *domain.Order) error {
	// InsertOne will set o.ID automatically if you're using _id tags in your struct
	_, err := r.collection.InsertOne(ctx, o)
	return err
}

func (r *OrderRepository) GetByID(ctx context.Context, id string) (*domain.Order, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id %q: %w", id, err)
	}

	var o domain.Order
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&o); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	// ensure the string ID is set
	o.ID = oid.Hex()
	return &o, nil
}

func (r *OrderRepository) Update(ctx context.Context, o *domain.Order) error {
	if o.ID == "" {
		return domain.ErrNotFound
	}
	oid, err := primitive.ObjectIDFromHex(o.ID)
	if err != nil {
		return fmt.Errorf("invalid id %q: %w", o.ID, err)
	}

	// update only the status and updated timestamp
	update := bson.M{
		"$set": bson.M{
			"status":     o.Status,
			"updated_at": o.UpdatedAt,
		},
	}
	_, err = r.collection.UpdateByID(ctx, oid, update)
	return err
}

func (r *OrderRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid id %q: %w", id, err)
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (r *OrderRepository) ListByUserID(ctx context.Context, userID string) ([]*domain.Order, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []*domain.Order
	for cursor.Next(ctx) {
		var o domain.Order
		if err := cursor.Decode(&o); err != nil {
			return nil, err
		}
		// set string ID from ObjectID
		if oid, err := primitive.ObjectIDFromHex(o.ID); err == nil {
			o.ID = oid.Hex()
		}
		orders = append(orders, &o)
	}
	return orders, nil
}
