package mongo_impl

import (
	"context"
	"time"

	mongoLibrary "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo"
)

type mongoClient struct {
	client *mongoLibrary.Client
	database *mongoLibrary.Database
}

func NewClient(ctx context.Context, dsn string, databaseName string) (mongo.Client, error) {
	clientOptions := options.Client().
		ApplyURI(dsn).
		SetConnectTimeout(30 * time.Second)

	client, err := mongoLibrary.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return &mongoClient{
		client: client,
		database: client.Database(databaseName),
	}, nil
}

func (m *mongoClient) Collection(collectionName string) mongo.Collection {
	return &mongoCollection{
		collection: m.database.Collection(collectionName),
	}
}

func (m *mongoClient) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *mongoClient) Ping(ctx context.Context) error {
	return m.client.Ping(ctx, readpref.Primary())
}
