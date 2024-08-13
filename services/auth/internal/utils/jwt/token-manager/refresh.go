package token_manager

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/jwt/token-manager/converter"
)

func (j *jwtTokenManager) GenerateRefreshToken(userClaims model.UserClaims) (string, error) {
	expiresAt := time.Now().Add(j.ttlRefreshKey)

	claims := converter.UserClaimsModelToJWT(userClaims, expiresAt)

	return generateToken(claims, j.secretRefreshKey, j.signingMethod)
}

func (j *jwtTokenManager) RegenerateRefreshToken(refreshToken string) (string, error) {
	claims, err := parseToken(refreshToken, j.secretRefreshKey)
	if err != nil {
		return "", err
	}

	return j.GenerateRefreshToken(converter.UserClaimsJWTToModel(*claims))
}

func generateToken(claims jwt.Claims, secretKey []byte, method jwt.SigningMethod) (string, error) {
	token := jwt.NewWithClaims(method, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("generating token: %w", err)
	}

	return tokenString, nil
}
