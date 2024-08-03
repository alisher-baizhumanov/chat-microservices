package message

import (
	serviceInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository"
)

type service struct {
	messageRepo repository.MessageRepository
}

// New creates a new instance of the MessageService implementation.
func New(messageRepo repository.MessageRepository) serviceInterface.MessageService {
	return &service{messageRepo: messageRepo}
}
