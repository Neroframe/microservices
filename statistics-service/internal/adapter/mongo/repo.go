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

// Repository implements usecase.StatisticsRepository using MongoDB.
type Repository struct {
	col *mongo.Collection
}

func NewRepository(db *mongo.Database) usecase.StatisticsRepository {
	return &Repository{
		col: db.Collection(eventsCollection),
	}
}

func (r *Repository) CountOrdersByUser(ctx context.Context, userID string) (int32, error) {
	filter := bson.M{
		"user_id":    userID,
		"event_type": "order_created",
	}
	cnt, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("CountOrdersByUser: %w", err)
	}
	return int32(cnt), nil
}

func (r *Repository) CountTotalUsers(ctx context.Context) (int32, error) {
	ids, err := r.col.Distinct(ctx, "user_id", bson.M{"event_type": "user_registered"})
	if err != nil {
		return 0, fmt.Errorf("CountTotalUsers: %w", err)
	}
	return int32(len(ids)), nil
}

func (r *Repository) CountDailyActiveUsers(ctx context.Context) (int32, error) {
	since := time.Now().Add(-24 * time.Hour)
	pipeline := mongo.Pipeline{
		{{"$match", bson.M{"timestamp": bson.M{"$gte": since}}}},
		{{"$group", bson.M{"_id": "$user_id"}}},
		{{"$count", "activeUsers"}},
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
