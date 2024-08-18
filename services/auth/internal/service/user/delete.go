package user

import (
	"context"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

// DeleteByID removes a user from the system identified by their unique identifier.
func (s *service) DeleteByID(ctx context.Context, id int64) error {
	if err := s.userCache.Delete(ctx, id); err != nil {
		logger.Warn("not deleted user in cache",
			logger.String("error", err.Error()),
			logger.Int64("id", id),
		)
	}

	if err := s.userRepository.DeleteUser(ctx, id); err != nil {
		logger.Warn("not deleted user in database",
			logger.String("error", err.Error()),
			logger.Int64("id", id),
		)

		return err
	}

	logger.Info("user deleted", logger.Int64("id", id))
	return nil
}
