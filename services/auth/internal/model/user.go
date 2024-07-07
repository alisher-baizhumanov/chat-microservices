package model

import "time"

type User struct {
	Id        int64
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegister struct {
	Name            string
	Email           string
	Password        []byte
	PasswordConfirm []byte
}

type UserUpdateOptions struct {
	Role  *Role
	Name  *string
	Email *string
}

type UserCreate struct {
	Name           string
	Email          string
	Role           Role
	HashedPassword []byte
	CreatedAt      time.Time
}
