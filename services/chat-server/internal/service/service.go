package service

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

// MessageService defines the interface for message-related operations.
type MessageService interface {
	Send(ctx context.Context, message model.MessageSave) (string, error)
}

// ChatService defines the interface for chat-related operations.
type ChatService interface {
	Save(ctx context.Context, chat model.ChatSave) (string, error)
	Delete(ctx context.Context, id string) error
}
