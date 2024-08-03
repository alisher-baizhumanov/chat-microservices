package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Participant represents a participant in a chat
type Participant struct {
	ChatID   primitive.ObjectID `bson:"chatId"`
	UserID   int64              `bson:"userId"`
	JoinedAt time.Time          `bson:"joinedAt"`
}
