package chat

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

func (r *repository) Create(ctx context.Context, chat model.ChatCreate, userIDList []int64) (string, error) {
	slog.InfoContext(ctx, "created chat",
		slog.Any("user_id_list", userIDList),
		slog.String("name", chat.Name),
		slog.Time("created_at", chat.CreatedAt),
	)

	return gofakeit.UUID(), nil
}
