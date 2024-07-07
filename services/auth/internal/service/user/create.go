package user

import (
	"context"
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

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
