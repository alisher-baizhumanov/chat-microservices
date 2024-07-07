package user

import "context"

func (s *Service) DeleteById(ctx context.Context, id int64) (err error) {
	return s.userRepository.DeleteUser(ctx, id)
}
