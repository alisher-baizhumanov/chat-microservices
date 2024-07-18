package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MessageCreate represents a structure for creating a new message.
type MessageCreate struct {
	ID        primitive.Binary   `bson:"_id"`
	UserID    int64              `bson:"userId"`
	ChatID    primitive.ObjectID `bson:"chatId"`
	Text      string             `bson:"text"`
	CreatedAt time.Time          `bson:"createdAt"`
}
