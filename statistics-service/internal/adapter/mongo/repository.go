package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const eventsCollection = "statistics_events"

var _ domain.StatisticsRepository = (*Repository)(nil)

type Repository struct {
	col *mongo.Collection
}

func NewRepository(db *mongo.Database) domain.StatisticsRepository {
	return &Repository{
		col: db.Collection(eventsCollection),
	}
}

// gRPC methods
func (r *Repository) CountOrdersByUser(ctx context.Context, userID string) (int32, error) {
	log.Printf("[Mongo] Counting orders for user_id=%s", userID)
	filter := bson.M{"user_id": userID, "event_type": "order_created"}
	count, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("CountOrdersByUser: %w", err)
	}
	log.Printf("[Mongo] CountOrdersByUser result: %d", count)
	return int32(count), nil
}

func (r *Repository) CountTotalUsers(ctx context.Context) (int32, error) {
	log.Println("[Mongo] Counting total users")

	filter := bson.M{"event_type": bson.M{"$in": []string{
		"order_created", "order_updated", "product_created", "product_updated",
	}}}

	ids, err := r.col.Distinct(ctx, "user_id", filter)
	if err != nil {
		return 0, fmt.Errorf("CountTotalUsers: %w", err)
	}
	log.Printf("[Mongo] Total unique users: %d", len(ids))
	return int32(len(ids)), nil
}

func (r *Repository) CountDailyActiveUsers(ctx context.Context) (int32, error) {
	since := time.Now().Add(-24 * time.Hour)
	log.Printf("[Mongo] Counting daily active users since %v", since)

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"timestamp": bson.M{"$gte": since}}}},
		{{Key: "$group", Value: bson.M{"_id": "$user_id"}}},
		{{Key: "$count", Value: "activeUsers"}},
	}

	cur, err := r.col.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("CountDailyActiveUsers.Aggregate: %w", err)
	}
	defer cur.Close(ctx)

	var res struct {
		ActiveUsers int32 `bson:"activeUsers"`
	}
	if cur.Next(ctx) {
		if err := cur.Decode(&res); err != nil {
			return 0, fmt.Errorf("CountDailyActiveUsers.Decode: %w", err)
		}
	}

	log.Printf("[Mongo] Daily active users: %d", res.ActiveUsers)
	return res.ActiveUsers, nil
}

// NATS methods
func (r *Repository) InsertOrderCreatedEvent(ctx context.Context, userID, orderID string, ts time.Time) error {
	return r.insertEvent(ctx, userID, orderID, "order_id", "order_created", ts)
}

func (r *Repository) InsertOrderUpdatedEvent(ctx context.Context, userID, orderID string, ts time.Time) error {
	return r.insertEvent(ctx, userID, orderID, "order_id", "order_updated", ts)
}

func (r *Repository) InsertOrderDeletedEvent(ctx context.Context, userID, orderID string, ts time.Time) error {
	return r.insertEvent(ctx, userID, orderID, "order_id", "order_deleted", ts)
}

func (r *Repository) InsertProductCreatedEvent(ctx context.Context, userID, productID string, ts time.Time) error {
	return r.insertEvent(ctx, userID, productID, "product_id", "product_created", ts)
}

func (r *Repository) InsertProductUpdatedEvent(ctx context.Context, userID, productID string, ts time.Time) error {
	return r.insertEvent(ctx, userID, productID, "product_id", "product_updated", ts)
}

func (r *Repository) InsertProductDeletedEvent(ctx context.Context, userID, productID string, ts time.Time) error {
	return r.insertEvent(ctx, userID, productID, "product_id", "product_deleted", ts)
}

func (r *Repository) insertEvent(ctx context.Context, userID, entityID, entityKey, eventType string, ts time.Time) error {
	// set timestamp 
	if ts.IsZero() {
		ts = time.Now().UTC()
	}

	doc := bson.M{
		"user_id":    userID,
		entityKey:    entityID,  // "order_id" or "product_id"
		"event_type": eventType, // e.g. "order_created"
		"timestamp":  ts,
	}

	res, err := r.col.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("insertEvent (%s): %w", eventType, err)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		log.Printf(
			"[Mongo] Inserted %s _id=%s, user_id=%s, %s=%s",
			eventType, oid.Hex(), userID, entityKey, entityID,
		)
	} else {
		log.Printf(
			"[Mongo] Inserted %s (non-ObjectID), user_id=%s, %s=%s; insertedID=%v",
			eventType, userID, entityKey, entityID, res.InsertedID,
		)
	}

	return nil
}