package chat

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

func (s *service) Save(ctx context.Context, chatSave model.ChatSave) (string, error) {
	createdAt := time.Now()
	id := primitive.NewObjectIDFromTimestamp(createdAt)

	chatCreate := model.Chat{
		ID:        id,
		Name:      chatSave.Name,
		CreatedAt: createdAt,
	}

	if err := s.chatRepo.CreateChat(ctx, chatCreate); err != nil {
		logger.Warn("failed to create a new chat", logger.String("error", err.Error()))
		return "", err
	}

	participants := model.NewParticipantList(chatSave.UserIDList, id, createdAt)

	if err := s.chatRepo.JoinParticipants(ctx, participants); err != nil {
		logger.Error("failed to join participants", logger.String("error", err.Error()))
		return "", err
	}

	idStr := id.Hex()
	logger.Info("chat created", logger.String("id", idStr))

	return idStr, nil
}
