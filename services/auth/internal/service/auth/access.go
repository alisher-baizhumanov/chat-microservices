package auth

import "context"

func (s *service) GetAccessToken(_ context.Context, refreshToken string) (string, error) {
	return s.tokenManager.GenerateAccessToken(refreshToken)
}
