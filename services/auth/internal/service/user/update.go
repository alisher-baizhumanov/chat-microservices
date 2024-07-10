package user

import (
	"context"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// UpdateUserFields updates specific fields of an existing user identified by their unique identifier.
// It delegates the update operation to the user repository and returns an error, if any.
func (s *Service) UpdateUserFields(ctx context.Context, id int64, userUpdate *model.UserUpdateOptions) (err error) {
	if userUpdate == nil {
		return model.ErrCanNotBeNil
	}

	return s.userRepository.UpdateUser(ctx, id, userUpdate)
}
