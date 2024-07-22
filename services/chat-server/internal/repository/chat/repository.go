package chat

import (
	repositoryInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
)

type repository struct {
}

// New creates a new instance of the ChatRepository implementation.
func New(_ any) repositoryInterface.ChatRepository {
	return &repository{}
}
