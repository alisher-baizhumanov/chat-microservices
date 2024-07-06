package model

type Role int8

const (
	NullRole  Role = 0
	UserRole       = 1
	AdminRole      = 2
)

func (r Role) String() string {
	return [...]string{"null", "user", "admin"}[r]
}
