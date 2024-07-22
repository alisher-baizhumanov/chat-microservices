package chat

import (
	"context"
	"time"

	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/model"
)

func (s *service) Save(ctx context.Context, chatSave model.ChatSave) (string, error) {
	chatCreate := model.ChatCreate{
		Name:      chatSave.Name,
		CreatedAt: time.Now(),
	}

	return s.chatRepo.Create(ctx, chatCreate, chatSave.UserIDList)
}
