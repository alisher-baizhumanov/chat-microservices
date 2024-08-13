package jwt

import "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"

type TokenManager interface {
	Verify(accessToken string) (model.UserClaims, error)
	GenerateAccessToken(refreshToken string) (string, error)
	RegenerateRefreshToken(refreshToken string) (string, error)
	GenerateRefreshToken(claims model.UserClaims) (string, error)
}
