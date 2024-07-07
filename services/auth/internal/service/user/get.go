package user

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func (s *Service) GetById(ctx context.Context, id int64) (user *model.User, err error) {
	return s.userRepository.GetUser(ctx, id)
}
