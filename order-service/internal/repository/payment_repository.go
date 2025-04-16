package repository

import (
	"context"
	"errors"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoPaymentRepo struct {
	collection *mongo.Collection
}

func NewPaymentMongoRepo(db *mongo.Database) domain.PaymentRepository {
	return &mongoPaymentRepo{
		collection: db.Collection("payments"),
	}
}

func (r *mongoPaymentRepo) Create(ctx context.Context, p *domain.Payment) error {
	res, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		p.ID = oid.Hex()
	}
	return nil
}

func (r *mongoPaymentRepo) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid payment ID format")
	}

	var payment domain.Payment
	if err := r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&payment); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	payment.ID = id
	return &payment, nil
}

func (r *mongoPaymentRepo) Update(ctx context.Context, p *domain.Payment) error {
	oid, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		return errors.New("invalid payment ID format")
	}
	update := bson.M{
		"order_id":       p.OrderID,
		"amount":         p.Amount,
		"payment_method": p.PaymentMethod,
		"status":         p.Status,
		"updated_at":     p.UpdatedAt,
	}
	_, err = r.collection.UpdateByID(ctx, oid, bson.M{"$set": update})
	return err
}

func (r *mongoPaymentRepo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid payment ID format")
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}
