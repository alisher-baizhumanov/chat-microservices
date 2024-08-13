package token_manager

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	jwtInterface "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/jwt"
)

type jwtTokenManager struct {
	signingMethod    jwt.SigningMethod
	secretAccessKey  []byte
	secretRefreshKey []byte
	ttlAccessKey     time.Duration
	ttlRefreshKey    time.Duration
}

func New(secretAccessKey, secretRefreshKey []byte, ttlAccessKey, ttlRefreshKey time.Duration) jwtInterface.TokenManager {
	return &jwtTokenManager{
		signingMethod:    jwt.SigningMethodHS256,
		secretAccessKey:  secretAccessKey,
		secretRefreshKey: secretRefreshKey,
		ttlAccessKey:     ttlAccessKey,
		ttlRefreshKey:    ttlRefreshKey,
	}
}
