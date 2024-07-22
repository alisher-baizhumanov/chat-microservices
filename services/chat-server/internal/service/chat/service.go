package chat

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
	serviceInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
)

type service struct {
	chatRepo repository.ChatRepository
}

// New creates a new instance of the ChatService implementation.
func New(chatRepo repository.ChatRepository) serviceInterface.ChatService {
	return &service{chatRepo: chatRepo}
}
