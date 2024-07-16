package message

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
	serviceInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
)

type service struct {
	messageRepo repository.MessageRepository
}

// New creates a new instance of the MessageService implementation.
func New(messageRepo repository.MessageRepository) serviceInterface.MessageService {
	return &service{messageRepo: messageRepo}
}
