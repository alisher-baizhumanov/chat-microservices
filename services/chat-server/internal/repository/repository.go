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
	CreateChat(ctx context.Context, chat model.Chat) error
	JoinParticipants(ctx context.Context, participants []model.Participant) error
	DeleteChatByID(ctx context.Context, id string) error
}
