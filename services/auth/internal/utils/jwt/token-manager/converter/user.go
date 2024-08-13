package converter

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
	jwtModel "github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/utils/jwt/token-manager/model"
)

func UserClaimsModelToJWT(userClaims model.UserClaims, expiresAt time.Time) jwtModel.UserJWTClaims {
	claims := jwtModel.UserJWTClaims{
		ID:   userClaims.ID,
		Role: int8(userClaims.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	return claims
}

func UserClaimsJWTToModel(userClaims jwtModel.UserJWTClaims) model.UserClaims {
	return model.UserClaims{
		ID:   userClaims.ID,
		Role: model.Role(userClaims.Role),
	}
}
