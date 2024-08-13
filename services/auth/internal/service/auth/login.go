package auth

import (
	"context"
	"errors"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	credentials, err := s.userRepository.GetCredentials(ctx, email)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			_, _ = s.hasher.Compare([]byte(password), []byte{}) // to prevent timing attacks

			return "", model.ErrInvalidCredentials
		}

		return "", err
	}

	ok, err := s.hasher.Compare([]byte(password), credentials.HashedPassword)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", model.ErrInvalidCredentials
	}

	return s.tokenManager.GenerateRefreshToken(model.UserClaims{
		ID:   credentials.ID,
		Role: credentials.Role,
	})
}
