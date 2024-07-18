package converter

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	data "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/message/model"
)

func MessageCreateModelToData(message model.MessageCreate) (data.MessageCreate, error) {
	chatID, err := primitive.ObjectIDFromHex(message.ChatID)
	if err != nil {
		return data.MessageCreate{}, err
	}

	id := primitive.Binary{Subtype: 4, Data: []byte(message.ID)}

	return data.MessageCreate{
		ID:        id,
		UserID:    message.UserID,
		ChatID:    chatID,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}, nil
}
