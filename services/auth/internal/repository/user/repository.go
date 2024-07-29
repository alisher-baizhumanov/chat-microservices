package user

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"
)

const (
	tableUser = "users"

	columnID             = "id"
	columnName           = "nickname"
	columnEmail          = "email"
	columndRole          = "role"
	columnCreatedAt      = "created_at"
	columnUpdatedAt      = "updated_at"
	columnHashedPassword = "hashed_password"

	constraintFieldUserName = tableUser + "_" + columnName + "_key"
	constraintFieldEmail    = tableUser + "_" + columnEmail + "_key"

	postgresUniqueErrorCode = "23505"
)

// repository represents a storage for user data.
type repository struct {
	client db.Client
}

// NewRepository creates and returns a new Repository instance.
func NewRepository(client db.Client) def.UserRepository {
	return &repository{client: client}
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
