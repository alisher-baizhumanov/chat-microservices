package converter

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	data "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/chat/model"
)

// ChatCreateModelToData converts a model.ChatCreate to a data.ChatCreate.
func ChatCreateModelToData(chat model.Chat) data.ChatCreate {
	return data.ChatCreate{
		ID:        chat.ID,
		Name:      chat.Name,
		CreatedAt: chat.CreatedAt,
	}
}
