package user

import (
	def "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/cache"
	repository "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/storage/repository"
)

// service provides methods to interact with user data through the user repository.
type service struct {
	userRepository repository.UserRepository
	userCache      cache.UserCache
}

// New creates a new Service instance with the given user repository.
// It returns a pointer to the newly created Service.
func New(userRepository repository.UserRepository, userCache cache.UserCache) def.UserService {
	return &service{userRepository: userRepository, userCache: userCache}
}
