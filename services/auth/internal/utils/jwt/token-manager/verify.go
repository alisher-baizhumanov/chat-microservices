package token_manager

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/jwt/token-manager/converter"
	jwtModel "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/jwt/token-manager/model"
)

func (j *jwtTokenManager) Verify(accessToken string) (model.UserClaims, error) {
	claims, err := parseToken(accessToken, j.secretAccessKey)
	if err != nil {
		return model.UserClaims{}, err
	}

	return converter.UserClaimsJWTToModel(*claims), nil
}

func parseToken(tokenString string, secretKey []byte) (*jwtModel.UserJWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtModel.UserJWTClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("%w, message: %w", model.ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(*jwtModel.UserJWTClaims)
	if !ok {
		return nil, fmt.Errorf("%w, message: can not be conberted to custom", model.ErrInvalidToken)
	}

	return claims, nil
}
