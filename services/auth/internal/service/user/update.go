package user

import (
	"context"
	"errors"
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// UpdateUserFields updates specific fields of an existing user identified by their unique identifier.
// It delegates the update operation to the user repository and returns an error, if any.
func (s *service) UpdateUserFields(ctx context.Context, id int64, userUpdate model.UserUpdateOptions) error {
	if err := s.userCache.Delete(ctx, id); err != nil {
		logger.Warn("not deleted user in cache",
			logger.String("error", err.Error()),
			logger.Int64("id", id),
		)
	}

	if err := s.userRepository.UpdateUser(ctx, id, userUpdate, time.Now()); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			logger.Info("user not found", logger.Int64("id", id))
		} else {
			logger.Warn("not updated user in database",
				logger.String("error", err.Error()),
				logger.Int64("id", id),
			)
		}

		return err
	}

	logger.Info("user updated", logger.Int64("id", id))
	return nil
}
