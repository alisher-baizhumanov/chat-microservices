package chat

import (
	"go.mongodb.org/mongo-driver/mongo"

	repositoryInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
)

const (
	CollectionChat         = "chats"
	CollectionParticipants = "participants"
)

type repository struct {
	collectionChat         *mongo.Collection
	collectionParticipants *mongo.Collection
}

// New creates a new instance of the ChatRepository implementation.
func New(chat *mongo.Collection, participants *mongo.Collection) repositoryInterface.ChatRepository {
	return &repository{collectionChat: chat, collectionParticipants: participants}
}
