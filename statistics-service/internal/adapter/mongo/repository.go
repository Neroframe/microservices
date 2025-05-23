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
func (r *Repository) InsertEvent(ctx context.Context, evt domain.Event) error {
	doc := bson.M{
		"user_id":    evt.UserID,
		"entity_key": evt.EntityKey,
		"entity_id":  evt.EntityID,
		"event_type": evt.EventType,
		"timestamp":  evt.Timestamp,
		"data": evt.Data,
	}

	_, err := r.col.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("mongo insert event: %w", err)
	}
	return nil
}

func (r *Repository) ListEvents(ctx context.Context) ([]domain.Event, error) {
	cur, err := r.col.Find(ctx, bson.D{}) 
	if err != nil {
		return nil, fmt.Errorf("mongo find events: %w", err)
	}
	defer cur.Close(ctx)

	var out []domain.Event
	for cur.Next(ctx) {
		var doc struct {
			ID        primitive.ObjectID     `bson:"_id"`
			UserID    string                 `bson:"user_id"`
			EntityKey string                 `bson:"entity_key"`
			EntityID  string                 `bson:"entity_id"`
			EventType string                 `bson:"event_type"`
			Timestamp time.Time              `bson:"timestamp"`
			Data      map[string]interface{} `bson:"data,omitempty"`
		}
		if err := cur.Decode(&doc); err != nil {
			return nil, fmt.Errorf("decode event doc: %w", err)
		}

		out = append(out, domain.Event{
			UserID:    doc.UserID,
			EntityKey: doc.EntityKey,
			EntityID:  doc.EntityID,
			EventType: doc.EventType,
			Timestamp: doc.Timestamp,
			Data:      doc.Data,
		})
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}
	return out, nil
}
