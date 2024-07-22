package converter

import (
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

// ChatProtoToModel converts a protobuf ChatCreate to a ChatSave model.
func ChatProtoToModel(chat *desc.ChatCreate) model.ChatSave {
	return model.ChatSave{
		Name:       chat.GetName(),
		UserIDList: chat.GetUserIdList(),
	}
}
