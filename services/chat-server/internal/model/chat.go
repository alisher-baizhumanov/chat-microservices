package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChatSave represents a structure for saving chat information.
type ChatSave struct {
	Name       string
	UserIDList []int64
}

// Chat represents a structure for creating a new chat.
type Chat struct {
	ID        primitive.ObjectID
	Name      string
	CreatedAt time.Time
}
