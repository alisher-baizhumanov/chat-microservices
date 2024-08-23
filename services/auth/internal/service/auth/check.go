package auth

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

func (s *service) CheckAccess(_ context.Context, path, accessToken string) error {
	claims, err := s.tokenManager.Verify(accessToken)
	if err != nil {
		logger.Info("invalid access token", logger.String("error", err.Error()))

		return err
	}

	logger.Info("checking access",
		logger.String("path", path),
		logger.Int64("user_id", claims.ID),
		logger.String("role", claims.Role.String()),
	)

	return nil
}
