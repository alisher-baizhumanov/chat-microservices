package converter

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	data "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository/chat/model"
)

// ParticipantModelToData converts a model.Participant to a data.Participant.
func ParticipantModelToData(p model.Participant) data.Participant {
	return data.Participant{
		ChatID:   p.ChatID,
		UserID:   p.UserID,
		JoinedAt: p.JoinedAt,
	}
}
