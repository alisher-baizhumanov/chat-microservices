package user

import def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"

var _ def.UserRepository = (*Repository)(nil)

// Repository represents a storage for user data.
type Repository struct {
}

// NewRepository creates and returns a new Repository instance.
func NewRepository() *Repository {
	return &Repository{}
}
