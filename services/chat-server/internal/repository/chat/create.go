package chat

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/chat/converter"
)

func (r *repository) Create(ctx context.Context, chatConverted model.ChatCreate, userIDList []int64) (string, error) {
	chat := converter.ChatCreateModelToData(chatConverted)

	res, err := r.collectionChat.InsertOne(ctx, chat)
	if err != nil {
		return "", fmt.Errorf("%w, message: %w", model.ErrDatabase, err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("%w, message: error asserting primitive.ObjectID", model.ErrGeneratingID)
	}

	slog.InfoContext(ctx, "created chat",
		slog.Any("user_id_list", userIDList),
	)

	return id.Hex(), nil
}
