package user

import def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"

var _ def.UserRepository = (*Repository)(nil)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}
