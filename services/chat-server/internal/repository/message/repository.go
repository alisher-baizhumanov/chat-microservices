package message

import (
	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo"
	repositoryInterface "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
)

const (
	// CollectionMessages is the name of the MongoDB collection storing messages documents.
	CollectionMessages = "messages"
)

type repository struct {
	collectionMessages mongo.Collection
}

// New creates a new instance of the MessageRepository implementation.
func New(collectionMessages mongo.Collection) repositoryInterface.MessageRepository {
	return &repository{collectionMessages: collectionMessages}
}
