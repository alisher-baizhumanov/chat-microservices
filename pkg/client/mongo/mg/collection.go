package mg

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoLibrary "go.mongodb.org/mongo-driver/mongo"
)

var (
	// ErrInvalidID is used when an invalid ID is provided, for example, an ID that does not exist or is incorrectly formatted.
	ErrInvalidID = errors.New("invalid id")
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

func (m *mongoCollection) UpdateByID(ctx context.Context, queryName string, id string, update map[string]any) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidID, err)
	}

	logQuery(ctx, queryName, id, update)

	filter := bson.M{"_id": objectID}
	mongoUpdate := toBsonM(update)

	_, err = m.collection.UpdateOne(ctx, filter, mongoUpdate)
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