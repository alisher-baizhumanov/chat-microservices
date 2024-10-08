package model

import "time"

// User represents a user info in the system with basic details and timestamps.
type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"nickname"`
	Email     string    `db:"email"`
	Role      int8      `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// UserUpdateOptions represents the options available for updating a user's details.
type UserUpdateOptions struct {
	Role  *int8
	Name  *string
	Email *string
}

// UserCreate represents the data required to create a new user.
type UserCreate struct {
	Name           string
	Email          string
	Role           int8
	HashedPassword []byte
	CreatedAt      time.Time
}
