package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Participant struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ChatID   primitive.ObjectID `bson:"chatId"`
	UserID   int                `bson:"userId"`
	JoinedAt time.Time          `bson:"joinedAt"`
}
