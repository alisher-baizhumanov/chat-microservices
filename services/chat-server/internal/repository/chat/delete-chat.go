package chat

import (
	"context"
	"errors"
	"fmt"
	"time"

	mg "github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo/mg"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

// DeleteChatByID marks a chat document as deleted by setting a timestamp in the "deletedAt" field.
// This function updates the chat document identified by the given ID with the current time,
// marking it as deleted. The actual document is not removed from the database.
func (r *repository) DeleteChatByID(ctx context.Context, id string) error {
	update := map[string]any{
		"$set": map[string]any{
			"deletedAt": time.Now(),
		},
	}

	if err := r.collectionChat.UpdateByID(ctx, "chat.Delete", id, update); err != nil {
		if errors.Is(err, mg.ErrInvalidID) {
			return fmt.Errorf("%w, message: %w", model.ErrInvalidID, err)
		}

		return fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	return nil
}
