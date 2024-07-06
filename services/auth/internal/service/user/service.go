package user

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"
	def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
)

var _ def.UserService = (*Service)(nil)

type Service struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *Service {
	return &Service{userRepository: userRepository}
}
