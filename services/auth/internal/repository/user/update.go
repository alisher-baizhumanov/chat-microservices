package user

import (
	"context"
	"log/slog"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// UpdateUser updates a user's information in the repository based on the provided user ID and update options.
func (r *Repository) UpdateUser(ctx context.Context, id int64, options *model.UserUpdateOptions) (err error) {
	if id < 1 {
		return model.ErrInvalidID
	}

	if options == nil {
		return model.ErrCanNotBeNil
	}

	log := slog.With("id", id)

	if options.Role != nil {
		log = log.With(slog.String("role", (*options.Role).String()))
	}

	if options.Name != nil {
		log = log.With(slog.String("name", *options.Name))
	}

	if options.Email != nil {
		log = log.With(slog.String("email", *options.Email))
	}

	log.InfoContext(ctx, "update user")
	return nil
}
