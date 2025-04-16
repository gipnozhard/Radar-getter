package repository

import (
	"Radar-getter/internal/config"
	"context"
	"time"

	"Radar-getter/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoRepository(cfg *config.MongoDBConfig) (*MongoRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	return &MongoRepository{
		client: client,
		db:     client.Database(cfg.DBName),
	}, nil
}

func (r *MongoRepository) SaveRealTimeData(ctx context.Context, data *domain.RealTimeData) error {
	collection := r.db.Collection("realtime_data")

	doc := map[string]interface{}{
		"equipment_id": data.EquipmentID,
		"data":         data.LaneData,
		"timestamp":    data.Timestamp,
	}

	_, err := collection.InsertOne(ctx, doc)
	return err
}

func (r *MongoRepository) SaveStatsData(ctx context.Context, data *domain.StatsData) error {
	collection := r.db.Collection("stats_data")

	doc := map[string]interface{}{
		"equipment_id": data.EquipmentID,
		"data":         data.Stats,
		"timestamp":    data.Timestamp,
	}

	_, err := collection.InsertOne(ctx, doc)
	return err
}

func (r *MongoRepository) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return r.client.Disconnect(ctx)
}
