package user

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"

	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/user/converter"
	data "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/user/model"
)

// UpdateUser updates a user's information in the repository based on the provided user ID and update options.
func (r *repository) UpdateUser(ctx context.Context, id int64, optionsConverted model.UserUpdateOptions, updatedAt time.Time) error {
	if id < 1 {
		return model.ErrInvalidID
	}

	options := converter.UserUpdateOptionModelToData(optionsConverted)

	sql, args, err := buildSQLQuery(&options, id, updatedAt)
	if err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrInvalidSQLQuery, err)
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: sql,
	}

	res, err := r.client.DB().Exec(ctx, q, args...)
	if err != nil {
		return convertUniqueDBErr(err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%w, id: %d", model.ErrNotFound, id)
	}

	return nil
}

func buildSQLQuery(options *data.UserUpdateOptions, id int64, updatedAt time.Time) (sql string, args []any, err error) {
	builder := sq.Update(tableUser).
		PlaceholderFormat(sq.Dollar).
		Set(columnUpdatedAt, updatedAt)

	if options.Name != nil {
		builder = builder.Set(columnName, *options.Name)
	}

	if options.Email != nil {
		builder = builder.Set(columnEmail, *options.Email)
	}

	if options.Role != nil {
		builder = builder.Set(columndRole, *options.Role)
	}

	return builder.Where(sq.Eq{columnID: id}).ToSql()
}
