package auth

import (
	"context"
	"errors"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func (s *service) GetRefreshToken(_ context.Context, refreshToken string) (string, error) {
	token, err := s.tokenManager.RegenerateRefreshToken(refreshToken)
	if err != nil {
		if errors.Is(err, model.ErrInvalidToken) {
			logger.Info("invalid refresh token")
		} else {
			logger.Warn("refresh token regeneration error", logger.String("error", err.Error()))
		}

		return "", err
	}

	logger.Info("refresh token regenerated")
	return token, nil
}
