package user

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user/converter"
	data "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user/model"
)

// GetUser retrieves a user from the repository based on the provided user ID.
func (r *Repository) GetUser(ctx context.Context, id int64) (*model.User, error) {
	if id < 1 {
		return nil, model.ErrInvalidID
	}

	sql, args, err := sq.Select("name", "email", "role", "created_at", "updated_at").
		From("users").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w, message: %w", model.ErrInvalidSQLQuery, err)
	}

	user := data.User{ID: id}

	if err := r.pool.QueryRow(ctx, sql, args...).
		Scan(&user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrNotFound
		}

		return nil, fmt.Errorf("%w, message: %w, id: %d", model.ErrDatabase, err, id)
	}

	return converter.UserDataToModel(&user), nil
}
