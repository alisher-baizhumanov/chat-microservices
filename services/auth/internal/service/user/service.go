package user

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
	def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
)

// service provides methods to interact with user data through the user repository.
type service struct {
	userRepository repository.UserRepository
	userCache      cache.UserCache
}

// NewService creates a new Service instance with the given user repository.
// It returns a pointer to the newly created Service.
func NewService(userRepository repository.UserRepository, userCache cache.UserCache) def.UserService {
	return &service{userRepository: userRepository, userCache: userCache}
}
