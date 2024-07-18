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

// NewParticipantList creates a list of Participant from a list of user IDs
func NewParticipantList(userIDList []int64, chatID primitive.ObjectID, joinedAt time.Time) []Participant {
	participants := make([]Participant, len(userIDList))

	for i, userID := range userIDList {
		participants[i] = Participant{
			ChatID:   chatID,
			UserID:   userID,
			JoinedAt: joinedAt,
		}
	}

	return participants
}
