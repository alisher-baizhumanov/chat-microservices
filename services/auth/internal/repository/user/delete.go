package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// DeleteUser removes a user from the repository based on the provided user ID.
func (r *Repository) DeleteUser(ctx context.Context, id int64) error {
	if id < 1 {
		return model.ErrInvalidID
	}

	sql, args, err := sq.Delete(tableNameUser).
		Where(sq.Eq{fieldID: id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrInvalidSQLQuery, err)
	}

	if _, err := r.pool.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	return nil
}
