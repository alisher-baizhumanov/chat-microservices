package message

import (
	repositoryInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
)

type repository struct {
}

// New creates a new instance of the MessageRepository implementation.
func New(_ any) repositoryInterface.MessageRepository {
	return &repository{}
}
