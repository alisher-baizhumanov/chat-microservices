package auth

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

func (s *service) GetAccessToken(_ context.Context, refreshToken string) (string, error) {
	accessToken, err := s.tokenManager.GenerateAccessToken(refreshToken)
	if err != nil {
		logger.Info("invalid token for generating access token", logger.String("error", err.Error()))

		return "", err
	}

	logger.Info("access token generated")
	return accessToken, nil
}
