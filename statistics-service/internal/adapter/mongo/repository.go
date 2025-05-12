package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/usecase"
	"go.mongodb.org/mongo-driver/bson"
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
	filter := bson.M{"user_id": userID, "event_type": "order_created"}
	cnt, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("CountOrdersByUser: %w", err)
	}
	return int32(cnt), nil
}

func (r *Repository) CountTotalUsers(ctx context.Context) (int32, error) {
	filter := bson.M{"event_type": "user_registered"}
	ids, err := r.col.Distinct(ctx, "user_id", filter)
	if err != nil {
		return 0, fmt.Errorf("CountTotalUsers: %w", err)
	}
	return int32(len(ids)), nil
}

func (r *Repository) CountDailyActiveUsers(ctx context.Context) (int32, error) {
	since := time.Now().Add(-24 * time.Hour)
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
	if _, err := r.col.InsertOne(ctx, doc); err != nil {
		return fmt.Errorf("InsertOrderCreatedEvent: %w", err)
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
