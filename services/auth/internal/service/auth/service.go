package auth

import (
	def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/hasher"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/jwt"
)

// service provides methods to interact with user data through the user repository.
type service struct {
	userRepository repository.UserRepository
	hasher         hasher.PasswordHasher
	tokenManager   jwt.TokenManager
}

// New creates a new Service instance with the given user repository.
// It returns a pointer to the newly created Service.
func New(userRepository repository.UserRepository, hasher hasher.PasswordHasher, manager jwt.TokenManager) def.AuthService {
	return &service{userRepository: userRepository, hasher: hasher, tokenManager: manager}
}
