package grpc

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"
)

type ServerHandlers struct {
	desc.UnimplementedChatServiceV1Server
}

func (h *ServerHandlers) CreateChat(ctx context.Context, chat *desc.CreateChatIn) (*desc.CreateChatOut, error) {
	slog.InfoContext(ctx, "created chat",
		slog.Any("user_id_list", chat.GetChat().GetUserIdList()),
		slog.String("name", chat.GetChat().GetName()),
	)

	return &desc.CreateChatOut{Id: gofakeit.UUID()}, nil
}

func (h *ServerHandlers) DeleteChat(ctx context.Context, chat *desc.DeleteChatIn) (*emptypb.Empty, error) {
	slog.InfoContext(ctx, "deleted chat",
		slog.String("id", chat.GetId()),
	)

	return nil, nil
}

func (h *ServerHandlers) SendMessage(ctx context.Context, message *desc.SendMessageIn) (*desc.SendMessageOut, error) {
	slog.InfoContext(ctx, "send message",
		slog.Int64("user_id", message.GetMessage().GetUserId()),
		slog.String("text", message.GetMessage().GetText()),
		slog.String("chat_id", message.GetMessage().GetChatId()),
	)

	return nil, nil
}
