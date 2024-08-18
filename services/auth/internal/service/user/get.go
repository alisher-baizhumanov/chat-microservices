package user

import (
	"context"
	"errors"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// GetByID retrieves a user by their unique identifier by delegating the call to the user repository.
func (s *service) GetByID(ctx context.Context, id int64) (model.User, error) {
	user, err := s.userCache.Get(ctx, id)
	if err == nil {
		logger.Info("got user from cache", logger.Int64("id", id))
		return user, nil
	}

	if !errors.Is(err, model.ErrNotFound) {
		logger.Warn("did not get user from cache",
			logger.String("error", err.Error()),
			logger.Int64("id", id),
		)
	}

	userDB, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			logger.Info("user not found", logger.Int64("id", id))
		} else {
			logger.Warn("did not get user from database", logger.String("error", err.Error()))
		}

		return model.User{}, err
	}

	if err = s.userCache.Set(ctx, userDB); err != nil {
		logger.Warn("not created user in cache",
			logger.String("error", err.Error()),
			logger.Int64("id", id),
		)
	}

	logger.Info("got user from database", logger.Int64("id", id))

	return userDB, nil
}
