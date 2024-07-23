package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewConnectionPool initializes a new MongoDB connection pool.
func NewConnectionPool(ctx context.Context, dsn string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(dsn).
		SetConnectTimeout(30 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

// CloseConnectionPool gracefully closes the given connection pool if it is not nil.
func CloseConnectionPool(ctx context.Context, pool *mongo.Client) error {
	if pool == nil {
		return nil
	}

	return pool.Disconnect(ctx)
}

// Client - a client for interacting with the database.
type Client interface {
	Collection(collectionName string) Collection
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
}

// Collection - an interface for interacting with a collecion.
type Collection interface {
	InsertOne(ctx context.Context, queryName string, document any) (string, error)
	InsertMany(ctx context.Context, queryName string, documents []any) error
	UpdateOne(ctx context.Context, queryName string, filter map[string]any, update map[string]any) error
}
