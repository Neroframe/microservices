package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const eventsCollection = "statistics_events"

// implements usecase.StatisticsRepository
type Repository struct {
	col *mongo.Collection
}

func NewRepository(db *mongo.Database) usecase.StatisticsRepository {
	return &Repository{
		col: db.Collection(eventsCollection),
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Read methods for gRPC
// ─────────────────────────────────────────────────────────────────────────────

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
	filter := bson.M{"event_type": "user_registered"}
	ids, err := r.col.Distinct(ctx, "user_id", filter)
	if err != nil {
		return 0, fmt.Errorf("CountTotalUsers: %w", err)
	}
	log.Printf("[Mongo] CountUsers result: %d", len(ids))
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

	log.Printf("[Mongo] CountDailyActiveUsers result: %d", res.ActiveUsers)
	return res.ActiveUsers, nil
}

// ─────────────────────────────────────────────────────────────────────────────
// Write methods for NATS event handling
// ─────────────────────────────────────────────────────────────────────────────

func (r *Repository) InsertOrderCreatedEvent(ctx context.Context, userID, orderID string, ts time.Time) error {
	doc := bson.M{
		"user_id":    userID,
		"order_id":   orderID,
		"event_type": "order_created",
		"timestamp":  ts,
	}

	res, err := r.col.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("InsertOrderCreatedEvent: %w", err)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		log.Printf(
			"[Mongo] Inserted OrderCreatedEvent _id=%s, user_id=%s, order_id=%s",
			oid.Hex(),
			userID,
			orderID,
		)
	} else {
		log.Printf(
			"[Mongo] Inserted OrderCreatedEvent (non-ObjectID) for order_id=%s; insertedID=%v",
			orderID,
			res.InsertedID,
		)
	}

	return nil
}

func (r *Repository) InsertUserRegisteredEvent(ctx context.Context, userID string, ts time.Time) error {
	doc := bson.M{
		"user_id":    userID,
		"event_type": "user_registered",
		"timestamp":  ts,
	}
	if _, err := r.col.InsertOne(ctx, doc); err != nil {
		return fmt.Errorf("InsertUserRegisteredEvent: %w", err)
	}
	return nil
}
