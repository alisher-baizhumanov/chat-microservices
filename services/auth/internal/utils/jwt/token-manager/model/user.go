package model

import "github.com/golang-jwt/jwt/v5"

// UserJWTClaims represents the claims in a JWT token for a user.
type UserJWTClaims struct {
	jwt.RegisteredClaims
	ID   int64 `json:"id"`
	Role int8  `json:"role"`
}
