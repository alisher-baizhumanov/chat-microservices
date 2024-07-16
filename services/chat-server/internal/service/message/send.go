package message

import (
	"context"
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

func (s *service) Send(ctx context.Context, messageSave model.MessageSave) (string, error) {
	msg := model.MessageCreate{
		ID:          "generated-id",
		MessageSave: messageSave,
		CreatedAt:   time.Now(),
	}

	if err := s.messageRepo.CreateMessage(ctx, msg); err != nil {
		return "", err
	}

	return msg.ID, nil
}
