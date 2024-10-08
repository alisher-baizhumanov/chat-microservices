package user

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/user/converter"
	data "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/user/model"
)

// GetUser retrieves a user from the repository based on the provided user ID.
func (r *repository) GetUser(ctx context.Context, id int64) (model.User, error) {
	if id < 1 {
		return model.User{}, model.ErrInvalidID
	}

	sql, args, err := sq.Select(columnName, columnEmail, columndRole, columnCreatedAt, columnUpdatedAt).
		From(tableUser).
		Where(sq.Eq{columnID: id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return model.User{}, fmt.Errorf("%w, message: %w", model.ErrInvalidSQLQuery, err)
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: sql,
	}

	user := data.User{ID: id}

	if err := r.client.DB().ScanOne(ctx, &user, q, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, model.ErrNotFound
		}

		return model.User{}, fmt.Errorf("%w, message: %w, id: %d", model.ErrDatabase, err, id)
	}

	return converter.UserDataToModel(user), nil
}
