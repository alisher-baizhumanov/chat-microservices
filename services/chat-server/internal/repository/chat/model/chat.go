package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChatCreate represents a chat creation data structure.
type ChatCreate struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at"`
}
