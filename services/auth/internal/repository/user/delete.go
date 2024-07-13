package user

import (
	"context"
	"log/slog"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// DeleteUser removes a user from the repository based on the provided user ID.
func (r *Repository) DeleteUser(ctx context.Context, id int64) (err error) {
	if id < 1 {
		return model.ErrInvalidID
	}

	slog.InfoContext(ctx, "delete user", slog.Int64("id", id))

	return nil
}
