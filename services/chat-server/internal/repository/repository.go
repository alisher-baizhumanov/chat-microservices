package repository

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

// MessageRepository defines the interface for message storage operations.
type MessageRepository interface {
	CreateMessage(ctx context.Context, message model.MessageCreate) error
}

// ChatRepository defines the interface for chat storage operations.
type ChatRepository interface {
	Create(ctx context.Context, chat model.ChatCreate, userIDList []int64) (string, error)
	DeleteByID(ctx context.Context, id string) error
}
