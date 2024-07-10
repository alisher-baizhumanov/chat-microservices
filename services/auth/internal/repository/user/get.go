package user

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// GetUser retrieves a user from the repository based on the provided user ID.
func (r *Repository) GetUser(ctx context.Context, id int64) (user *model.User, err error) {
	if id < 1 {
		return nil, model.ErrInvalidID
	}

	slog.InfoContext(ctx, "get user", slog.Int64("id", id))

	createdAt := gofakeit.Date()

	return &model.User{
		ID:        id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      model.UserRole,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}, nil
}
