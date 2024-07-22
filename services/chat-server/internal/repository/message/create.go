package message

import (
	"context"
	"log/slog"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

func (r *repository) CreateMessage(ctx context.Context, message model.MessageCreate) error {
	slog.InfoContext(ctx, "send message",
		slog.Int64("user_id", message.UserID),
		slog.String("text", message.Text),
		slog.String("chat_id", message.ChatID),
		slog.Time("created_at", message.CreatedAt),
		slog.String("id", message.ID),
	)

	return nil
}
