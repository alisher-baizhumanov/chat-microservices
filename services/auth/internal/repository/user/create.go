package user

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// CreateUser creates a new user in the repository with the provided user creation data.
func (r *Repository) CreateUser(ctx context.Context, user *model.UserCreate) (id int64, err error) {
	if user == nil {
		return 0, model.ErrCanNotBeNil
	}

	slog.InfoContext(ctx, "created user",
		slog.String("name", user.Name),
		slog.String("email", user.Email),
		slog.String("hashed_password", string(user.HashedPassword)),
	)

	return gofakeit.Int64(), nil
}