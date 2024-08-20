package auth

import "context"

func (s *service) GetRefreshToken(_ context.Context, refreshToken string) (string, error) {
	return s.tokenManager.RegenerateRefreshToken(refreshToken)
}
