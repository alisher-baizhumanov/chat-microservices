package user

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository/user/converter"
)

// CreateUser creates a new user in the repository with the provided user creation data.
func (r *Repository) CreateUser(ctx context.Context, userConverted *model.UserCreate) (int64, error) {
	if userConverted == nil {
		return 0, model.ErrCanNotBeNil
	}

	user := converter.UserCreateModelToData(userConverted)

	sql, args, err := sq.Insert(tableNameUser).
		Columns(fieldName, fieldEmail, fieldRole, fieldCreatedAt, fieldUpdatedAt).
		Values(user.Name, user.Email, user.HashedPassword, user.Role, user.CreatedAt, user.CreatedAt).
		Suffix("RETURNING " + fieldID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("%w, message: %w", model.ErrInvalidSQLQuery, err)
	}

	var id int64
	if err := r.pool.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return 0, convertUniqueDBErr(err)
	}

	return id, nil
}

func convertUniqueDBErr(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == postgresUniqueErrorCode {
			switch pgErr.ConstraintName {
			case constraintFieldUserName:
				return fmt.Errorf("%w, message: %s", model.ErrNonUniqueUsername, pgErr.Message)
			case constraintFieldEmail:
				return fmt.Errorf("%w, message: %s", model.ErrNonUniqueEmail, pgErr.Message)
			}
		}
	}

	return fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
}
