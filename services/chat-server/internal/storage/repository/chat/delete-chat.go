package chat

import (
	"context"
	"errors"
	"fmt"
	"time"

	mg "github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo/mg"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

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
