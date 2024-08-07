package user

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/repository"
	def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
)

var _ def.UserService = (*Service)(nil)

// Service provides methods to interact with user data through the user repository.
type Service struct {
	userRepository repository.UserRepository
}

// NewService creates a new Service instance with the given user repository.
// It returns a pointer to the newly created Service.
func NewService(userRepository repository.UserRepository) *Service {
	return &Service{userRepository: userRepository}
}
