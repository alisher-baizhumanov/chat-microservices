package user

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user/converter"
)

// UpdateUser updates a user's information in the repository based on the provided user ID and update options.
func (r *Repository) UpdateUser(ctx context.Context, id int64, optionsConverted *model.UserUpdateOptions) error {
	if id < 1 {
		return model.ErrInvalidID
	}

	if optionsConverted == nil {
		return model.ErrCanNotBeNil
	}

	options := converter.UserUpdateOptionModelToData(optionsConverted)

	builder := sq.Update(tableNameUser).
		PlaceholderFormat(sq.Dollar).
		Set(fieldUpdatedAt, time.Now())

	if options.Name != nil {
		builder = builder.Set(fieldName, *options.Name)
	}

	if options.Email != nil {
		builder = builder.Set(fieldEmail, *options.Email)
	}

	if options.Role != nil {
		builder = builder.Set(fieldRole, *options.Role)
	}

	sql, args, err := builder.Where(sq.Eq{fieldID: id}).ToSql()
	if err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrInvalidSQLQuery, err)
	}

	res, err := r.pool.Exec(ctx, sql, args...)
	if err != nil {
		return convertUniqueDBErr(err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%w, id: %d", model.ErrNotFound, id)
	}

	return nil
}
