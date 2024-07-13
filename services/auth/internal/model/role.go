package model

// Role represents the role of a user in the system.
type Role int8

const (
	// NullRole indicates an undefined  role.
	NullRole Role = iota

	// UserRole indicates a standard user role.
	UserRole

	// AdminRole indicates an administrative user role.
	AdminRole
)

func (r Role) String() string {
	return [...]string{"null", "user", "admin"}[r]
}
