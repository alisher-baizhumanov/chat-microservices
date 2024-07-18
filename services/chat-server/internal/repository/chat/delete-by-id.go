package chat

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

func (r *repository) DeleteByID(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrInvalidID, err)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"deletedAt": time.Now(),
		},
	}

	_, err = r.collectionChat.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	return nil
}
