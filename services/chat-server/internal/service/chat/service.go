package chat

import (
	serviceInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository"
)

type service struct {
	chatRepo repository.ChatRepository
}

// New creates a new instance of the ChatService implementation.
func New(chatRepo repository.ChatRepository) serviceInterface.ChatService {
	return &service{chatRepo: chatRepo}
}
