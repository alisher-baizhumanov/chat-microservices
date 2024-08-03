package model

// User represents a user info in the cache.
type User struct {
	Name      string `redis:"name"`
	Email     string `redis:"email"`
	Role      int8   `redis:"role"`
	CreatedAt string `redis:"createdAt"`
	UpdatedAt string `redis:"updatedAt"`
}
