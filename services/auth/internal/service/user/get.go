package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// GetByID retrieves a user by their unique identifier by delegating the call to the user repository.
func (s *Service) GetByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userCache.Get(ctx, id)
	if err == nil {
		return &user, nil
	}

	if !errors.Is(err, model.ErrNotFound) {
		slog.ErrorContext(ctx, "not got user from cache",
			slog.String("error", err.Error()),
			slog.Int64("id", id),
		)
	}

	userDB, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if err = s.userCache.Set(ctx, *userDB); err != nil {
		slog.ErrorContext(ctx, "not created user in cache",
			slog.String("error", err.Error()),
			slog.Int64("id", id),
		)
	}

	return userDB, nil
}
