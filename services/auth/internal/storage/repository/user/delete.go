package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// DeleteUser removes a user from the repository based on the provided user ID.
func (r *repository) DeleteUser(ctx context.Context, id int64) error {
	if id < 1 {
		return model.ErrInvalidID
	}

	sql, args, err := sq.Delete(tableUser).
		Where(sq.Eq{columnID: id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrInvalidSQLQuery, err)
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: sql,
	}

	if _, err := r.client.DB().Exec(ctx, q, args...); err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	return nil
}
