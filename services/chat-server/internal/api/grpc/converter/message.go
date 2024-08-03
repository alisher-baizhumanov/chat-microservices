package converter

import (
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

// MessageProtoToModel converts a protobuf MessageCreate to a MessageSave model.
func MessageProtoToModel(message *desc.MessageCreate) model.MessageSave {
	return model.MessageSave{
		UserID: message.GetUserId(),
		Text:   message.GetText(),
		ChatID: message.GetChatId(),
	}
}
