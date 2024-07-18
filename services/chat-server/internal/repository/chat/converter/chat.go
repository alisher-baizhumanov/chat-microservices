package converter

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	data "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/chat/model"
)

// ChatCreateModelToData converts a model.ChatCreate to a data.ChatCreate.
// model.ChatCreate does not contain data.ChatCreate ID
func ChatCreateModelToData(chat model.ChatCreate) data.ChatCreate {
	return data.ChatCreate{
		Name:      chat.Name,
		CreatedAt: chat.CreatedAt,
	}
}
