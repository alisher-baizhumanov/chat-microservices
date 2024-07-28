package model

import "time"

// User represents a user info in the cache.
type User struct {
	Name      string    `redis:"name"`
	Email     string    `redis:"email"`
	Role      int8      `redis:"role"`
	CreatedAt time.Time `redis:"createdAt"`
	UpdatedAt time.Time `redis:"updatedAt"`
}
