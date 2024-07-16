package model

import "time"

// ChatSave represents a structure for saving chat information.
type ChatSave struct {
	Name       string
	UserIDList []int64
}

// ChatCreate represents a structure for creating a new chat.
type ChatCreate struct {
	Name      string
	CreatedAt time.Time
}
