package model

import "time"

type ChatCreate struct {
	Name      string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
}
