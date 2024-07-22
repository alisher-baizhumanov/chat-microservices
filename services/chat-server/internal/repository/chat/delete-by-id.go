package chat

import (
	"context"
	"log/slog"
)

func (r *repository) DeleteByID(ctx context.Context, id string) error {
	slog.InfoContext(ctx, "deleted chat",
		slog.String("id", id),
	)

	return nil
}
