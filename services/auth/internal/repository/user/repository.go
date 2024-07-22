package user

import (
	db "github.com/alisher-baizhumanov/chat-microservices/pkg/client/postgres"
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

var _ def.UserRepository = (*Repository)(nil)

// Repository represents a storage for user data.
type Repository struct {
	client db.Client
}

// NewRepository creates and returns a new Repository instance.
func NewRepository(client db.Client) *Repository {
	return &Repository{client: client}
}
