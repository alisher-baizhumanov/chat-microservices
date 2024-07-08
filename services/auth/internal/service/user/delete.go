package user

import "context"

// DeleteByID removes a user from the system identified by their unique identifier.
func (s *Service) DeleteByID(ctx context.Context, id int64) (err error) {
	return s.userRepository.DeleteUser(ctx, id)
}
