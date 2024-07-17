package chat

import (
	"go.mongodb.org/mongo-driver/mongo"

	repositoryInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
)

type repository struct {
	client *mongo.Client
}

// New creates a new instance of the ChatRepository implementation.
func New(client *mongo.Client) repositoryInterface.ChatRepository {
	return &repository{client: client}
}
