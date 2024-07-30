package mongo

import (
	"context"
)

// Client - a client for interacting with the database.
type Client interface {
	Collection(collectionName string) Collection
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
}

// Collection - an interface for interacting with a collection.
type Collection interface {
	InsertOne(ctx context.Context, queryName string, document any) (string, error)
	InsertMany(ctx context.Context, queryName string, documents []any) error
	UpdateByID(ctx context.Context, queryName string, id string, update map[string]any) error
}
