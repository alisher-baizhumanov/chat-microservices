package user

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

func (s *Service) UpdateUserFields(ctx context.Context, id int64, userUpdate *model.UserUpdateOptions) (err error) {
	return s.userRepository.UpdateUser(ctx, id, userUpdate)
}
