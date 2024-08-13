package token_manager

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (j *jwtTokenManager) GenerateAccessToken(refreshToken string) (string, error) {
	claims, err := parseToken(refreshToken, j.secretRefreshKey)
	if err != nil {
		return "", err
	}

	expiresAt := time.Now().Add(j.ttlRefreshKey)

	claims.ExpiresAt = &jwt.NumericDate{Time: expiresAt}

	return generateToken(claims, j.secretAccessKey, j.signingMethod)
}
