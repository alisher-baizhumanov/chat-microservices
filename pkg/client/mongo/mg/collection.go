package mongo_impl

import (
	"log/slog"
	"fmt"
	"context"

/* 	"go.mongodb.org/mongo-driver/bson"
 */	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoLibrary "go.mongodb.org/mongo-driver/mongo"
)

type mongoCollection struct {
	collection *mongoLibrary.Collection
}

func (m *mongoCollection) InsertOne(ctx context.Context, queryName string, document any) (string, error) {
	logQuery(ctx, queryName, document)

	id, err := m.collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	// The _id of the inserted document. A value generated by the driver will be of type primitive.ObjectID.
	if primitiveID, ok := id.InsertedID.(primitive.ObjectID); ok {
		return primitiveID.Hex(), nil
	}

	return fmt.Sprintf("%v", id.InsertedID), nil
}

func (m *mongoCollection) InsertMany(ctx context.Context, queryName string, documents []any) error {
	logQuery(ctx, queryName, documents)

	_, err := m.collection.InsertMany(ctx, documents)
	return err
}

func (m *mongoCollection) UpdateOne(ctx context.Context, queryName string, filter map[string]any, update map[string]any) error {
	logQuery(ctx, queryName, filter, update)


/* 	mongoFilter := bson.M(filter)
	mongoUpdate := bson.M(filter)

	_, err := m.collection.UpdateOne(ctx, mongoFilter, mongoUpdate) */
	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err
}

func logQuery(ctx context.Context, queryName string, args ...any) {
	if slog.Default().Enabled(ctx, slog.LevelDebug) {
		slog.DebugContext(
			ctx,
			queryName,
			slog.Any("args", args),
		)
	}
}
