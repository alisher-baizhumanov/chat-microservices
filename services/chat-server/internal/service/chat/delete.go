package chat

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

func (s *service) Delete(ctx context.Context, id string) error {
	if err := s.chatRepo.DeleteChatByID(ctx, id); err != nil {
		logger.Warn("failed to delete chat", logger.String("error", err.Error()))
		return err
	}

	logger.Info("chat deleted", logger.String("id", id))
	return nil
}
