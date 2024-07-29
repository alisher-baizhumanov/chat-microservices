package user

import (
	"context"
	"log/slog"
)

// DeleteByID removes a user from the system identified by their unique identifier.
func (s *service) DeleteByID(ctx context.Context, id int64) error {
	if err := s.userRepository.DeleteUser(ctx, id); err != nil {
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
