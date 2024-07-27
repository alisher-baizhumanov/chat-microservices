package chat

import (
	"context"
	"fmt"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/chat/converter"
)

func (r *repository) CreateChat(ctx context.Context, chatConverted model.Chat) error {
	chat := converter.ChatCreateModelToData(chatConverted)

	if _, err := r.collectionChat.InsertOne(ctx, "chat.Create", chat); err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	return nil
}
