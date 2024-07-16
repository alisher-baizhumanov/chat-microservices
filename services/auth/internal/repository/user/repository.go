package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"

	def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"
)

const (
	tableNameUser = "users"

	fieldID        = "id"
	fieldName      = "name"
	fieldEmail     = "email"
	fieldRole      = "role"
	fieldCreatedAt = "created_at"
	fieldUpdatedAt = "updated_at"

	constraintFieldUserName = tableNameUser + "_" + fieldName + "_key"
	constraintFieldEmail    = tableNameUser + "_" + fieldEmail + "_key"

	postgresUniqueErrorCode = "23505"
)

var _ def.UserRepository = (*Repository)(nil)

// Repository represents a storage for user data.
type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository creates and returns a new Repository instance.
func NewRepository(ctx context.Context, connectionString string) (*Repository, error) {
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Repository{
		pool: pool,
	}, nil
}

// Stop gracefully stops the database connection pool.
func (r *Repository) Stop() {
	r.pool.Close()
}
