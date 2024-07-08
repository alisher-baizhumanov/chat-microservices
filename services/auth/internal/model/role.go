package model

// Role represents the role of a user in the system.
type Role int8

const (
	// NullRole indicates an undefined  role.
	NullRole Role = 0

	// UserRole indicates a standard user role.
	UserRole = 1

	// AdminRole indicates an administrative user role.
	AdminRole = 2
)

func (r Role) String() string {
	return [...]string{"null", "user", "admin"}[r]
}
