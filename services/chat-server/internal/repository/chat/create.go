package chat

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/chat/converter"
	data "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/chat/model"
)

func (r *repository) Create(ctx context.Context, chatConverted model.ChatCreate, userIDList []int64) (string, error) {
	chat := converter.ChatCreateModelToData(chatConverted)
	chat.ID = primitive.NewObjectID()

	participants := data.NewParticipantList(userIDList, chat.ID, chat.CreatedAt)
	documents := make([]interface{}, len(participants))
	for i, participant := range participants {
		documents[i] = participant
	}

	if _, err := r.collectionChat.InsertOne(ctx, chat); err != nil {
		return "", fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	if _, err := r.collectionParticipants.InsertMany(ctx, documents); err != nil {
		return "", fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	return chat.ID.Hex(), nil
}
