package model

import "github.com/golang-jwt/jwt/v5"

type UserJWTClaims struct {
	jwt.RegisteredClaims
	ID   int64 `json:"id"`
	Role int8  `json:"role"`
}
