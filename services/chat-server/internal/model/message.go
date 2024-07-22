package model

import "time"

// MessageSave represents a structure for saving a message.
type MessageSave struct {
	UserID int64
	ChatID string
	Text   string
}

// MessageCreate represents a structure for creating a new message.
type MessageCreate struct {
	ID string
	MessageSave
	CreatedAt time.Time
}
