package mongodb

import (
	"context"
	"getcare-notification/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

// NewMongoDBConn Create new MongoDB client
func New(ctx context.Context, cfg *config.Config) (*mongo.Database, error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(cfg.MongoDB.URI).
			// SetAuth(options.Credential{
			// 	Username: cfg.MongoDB.User,
			// 	Password: cfg.MongoDB.Password,
			// }).
			SetConnectTimeout(connectTimeout).
			SetMaxConnIdleTime(maxConnIdleTime).
			SetMinPoolSize(minPoolSize).
			SetMaxPoolSize(maxPoolSize))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(cfg.MongoDB.DB), nil
}
