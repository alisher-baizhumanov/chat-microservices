package chat

import (
	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo"
	repositoryInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository"
)

const (
	// CollectionChat is the name of the MongoDB collection storing chat documents.
	CollectionChat = "chats"

	// CollectionParticipants is the name of the MongoDB collection storing participant documents.
	CollectionParticipants = "participants"
)

type repository struct {
	collectionChat         mongo.Collection
	collectionParticipants mongo.Collection
}

// New creates a new instance of the ChatRepository implementation.
func New(chat mongo.Collection, participants mongo.Collection) repositoryInterface.ChatRepository {
	return &repository{collectionChat: chat, collectionParticipants: participants}
}
