package message

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

// Send creates and saves a new message, generating a unique ID for it.
func (s *service) Send(ctx context.Context, messageSave model.MessageSave) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		logger.Info("failed to generate a new UUID", logger.String("error", err.Error()))
		return "", fmt.Errorf("%w, message: %w", model.ErrGeneratingID, err)
	}

	msg := model.MessageCreate{
		ID:          id.String(),
		MessageSave: messageSave,
		CreatedAt:   time.Now(),
	}

	if err = s.messageRepo.CreateMessage(ctx, msg); err != nil {
		logger.Warn("failed to create a new message", logger.String("error", err.Error()))
		return "", err
	}

	logger.Info("message created", logger.String("id", msg.ID))
	return msg.ID, nil
}
