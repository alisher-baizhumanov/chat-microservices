package auth

import (
	"context"
	"errors"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	credentials, err := s.userRepository.GetCredentials(ctx, email)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			logger.Info("invalid credentials")

			_, _ = s.hasher.Compare([]byte(password), []byte{}) // to prevent timing attacks
			return "", model.ErrInvalidCredentials
		}

		logger.Info("database error", logger.String("error", err.Error()))
		return "", err
	}

	ok, err := s.hasher.Compare([]byte(password), credentials.HashedPassword)
	if err != nil {
		logger.Warn("hasher error", logger.String("error", err.Error()))
		return "", err
	}

	if !ok {
		logger.Info("invalid credentials", logger.String("email", email))
		return "", model.ErrInvalidCredentials
	}

	token, err := s.tokenManager.GenerateRefreshToken(model.UserClaims{
		ID:   credentials.ID,
		Role: credentials.Role,
	})
	if err != nil {
		logger.Warn("token generation error", logger.String("error", err.Error()))
		return "", err
	}

	logger.Info("login successful", logger.Int64("user_id", credentials.ID))
	return token, nil
}
