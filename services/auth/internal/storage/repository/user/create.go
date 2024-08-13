package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository/user/converter"
)

// CreateUser creates a new user in the repository with the provided user creation data.
func (r *repository) CreateUser(ctx context.Context, userConverted model.UserCreate) (int64, error) {
	user := converter.UserCreateModelToData(userConverted)

	sql, args, err := sq.Insert(tableUser).
		Columns(columnName, columnEmail, columnHashedPassword, columnRole, columnCreatedAt, columnUpdatedAt).
		Values(user.Name, user.Email, user.HashedPassword, user.Role, user.CreatedAt, user.CreatedAt).
		Suffix("RETURNING " + columnID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w, message: %w", model.ErrInvalidSQLQuery, err)
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: sql,
	}

	var id int64
	if err := r.client.DB().QueryRow(ctx, q, args...).Scan(&id); err != nil {
		return 0, convertUniqueDBErr(err)
	}

	return id, nil
}
