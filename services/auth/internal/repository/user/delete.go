package user

import (
	"context"
	"log/slog"
)

func (r *Repository) DeleteUser(ctx context.Context, id int64) (err error) {
	slog.InfoContext(ctx, "delete user", slog.Int64("id", id))

	return nil
}
