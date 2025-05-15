package mongo

import (
	"context"
	"fmt"

	"github.com/Neroframe/ecommerce-platform/order-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// implements domain.PaymentRepository 
type PaymentRepository struct {
	collection *mongo.Collection
}

// NewPaymentRepository returns a new mongo-backed PaymentRepository.
func NewPaymentRepository(db *mongo.Database) domain.PaymentRepository {
	return &PaymentRepository{
		collection: db.Collection("payments"),
	}
}

// Create inserts a new Payment and sets its ID field.
func (r *PaymentRepository) Create(ctx context.Context, p *domain.Payment) error {
	res, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		return fmt.Errorf("mongo insert error: %w", err)
	}
	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		p.ID = oid.Hex()
	}
	return nil
}

// GetByID finds a Payment by its string ID.
func (r *PaymentRepository) GetByID(ctx context.Context, id string) (*domain.Payment, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid payment ID %q: %w", id, err)
	}

	var p domain.Payment
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("mongo find error: %w", err)
	}
	p.ID = oid.Hex()
	return &p, nil
}

// Update applies changes to an existing Payment.
func (r *PaymentRepository) Update(ctx context.Context, p *domain.Payment) error {
	if p.ID == "" {
		return domain.ErrNotFound
	}
	oid, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		return fmt.Errorf("invalid payment ID %q: %w", p.ID, err)
	}

	update := bson.M{
		"$set": bson.M{
			"order_id":       p.OrderID,
			"amount":         p.Amount,
			"payment_method": p.PaymentMethod,
			"status":         p.Status,
			"updated_at":     p.UpdatedAt,
		},
	}
	_, err = r.collection.UpdateByID(ctx, oid, update)
	if err != nil {
		return fmt.Errorf("mongo update error: %w", err)
	}
	return nil
}

// Delete removes a Payment by its ID.
func (r *PaymentRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid payment ID %q: %w", id, err)
	}
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return fmt.Errorf("mongo delete error: %w", err)
	}
	return nil
}
