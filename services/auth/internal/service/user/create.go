package user

import (
	"context"
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// RegisterUser registers a new user in the system with the given registration details.
// It converts the UserRegister model to a UserCreate model, sets the default user role,
// and assigns the current time as the creation time. The function then delegates the
// creation operation to the user repository and returns the unique identifier of the
// newly registered user and an error, if any.
func (s *Service) RegisterUser(ctx context.Context, userRegister *model.UserRegister) (id int64, err error) {
	userCreate := &model.UserCreate{
		Name:           userRegister.Name,
		Email:          userRegister.Email,
		Role:           model.UserRole,
		HashedPassword: userRegister.Password,
		CreatedAt:      time.Now(),
	}

	return s.userRepository.CreateUser(ctx, userCreate)
}
