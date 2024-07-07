package user

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func (r *Repository) GetUser(ctx context.Context, id int64) (user *model.User, err error) {
	slog.InfoContext(ctx, "get user", slog.Int64("id", id))

	createdAt := gofakeit.Date()

	return &model.User{
		Id:        id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      model.UserRole,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}, nil
}
