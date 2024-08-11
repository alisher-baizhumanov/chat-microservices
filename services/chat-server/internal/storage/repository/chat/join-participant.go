package chat

import (
	"context"
	"fmt"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository/chat/converter"
)

// JoinParticipants adds multiple participants to the database.
func (r *repository) JoinParticipants(ctx context.Context, participantsConverted []model.Participant) error {
	participants := make([]any, len(participantsConverted))

	for i, participantConverted := range participantsConverted {
		participants[i] = converter.ParticipantModelToData(participantConverted)
	}

	if err := r.collectionParticipants.InsertMany(ctx, "participant.InitChat", participants); err != nil {
		return fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	return nil
}
