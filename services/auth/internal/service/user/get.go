package user

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// GetByID retrieves a user by their unique identifier by delegating the call to the user repository.
func (s *Service) GetByID(ctx context.Context, id int64) (user *model.User, err error) {
	return s.userRepository.GetUser(ctx, id)
}
