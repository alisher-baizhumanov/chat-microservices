package model

import (
	"time"
)

// User represents a user info in the system with basic details and timestamps.
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserRegister represents the data required to register a new user.
type UserRegister struct {
	Name            string
	Email           string
	Password        []byte
	PasswordConfirm []byte
}

// UserUpdateOptions represents the options available for updating a user's details.
type UserUpdateOptions struct {
	Role  *Role
	Name  *string
	Email *string
}

// UserCreate represents the data required to create a new user.
type UserCreate struct {
	Name           string
	Email          string
	Role           Role
	HashedPassword []byte
	CreatedAt      time.Time
}

type UserClaims struct {
	ID   int64
	Role Role
}
