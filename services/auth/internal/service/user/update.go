package user

import (
	"context"
	"log/slog"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// UpdateUserFields updates specific fields of an existing user identified by their unique identifier.
// It delegates the update operation to the user repository and returns an error, if any.
func (s *service) UpdateUserFields(ctx context.Context, id int64, userUpdate model.UserUpdateOptions) error {
	if err := s.userRepository.UpdateUser(ctx, id, userUpdate); err != nil {
		return err
	}

	if err := s.userCache.Delete(ctx, id); err != nil {
		slog.ErrorContext(ctx, "not deleted user in cache",
			slog.String("error", err.Error()),
			slog.Int64("id", id),
		)
	}

	return nil
}
