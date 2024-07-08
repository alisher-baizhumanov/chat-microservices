package grpc

import (
	"context"
	"log/slog"

	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"
)

type chatServer struct {
	desc.UnimplementedChatServiceV1Server
}

func (s *chatServer) Create(ctx context.Context, createIn *desc.CreateIn) (*desc.CreateOut, error) {
	slog.InfoContext(ctx, "created chat",
		slog.Any("usernames", createIn.GetUserIdList()),
		slog.String("name", createIn.GetName()),
	)

	return &desc.CreateOut{Id: gofakeit.Int64()}, nil
}

func (s *chatServer) Delete(ctx context.Context, deleteIn *desc.DeleteIn) (*emptypb.Empty, error) {
	slog.InfoContext(ctx, "deleted chat",
		slog.Int64("id", deleteIn.GetId()),
	)

	return nil, nil
}

func (s *chatServer) SendMessage(ctx context.Context, sendMessageIn *desc.SendMessageIn) (*emptypb.Empty, error) {
	message := sendMessageIn.GetMessage()

	slog.InfoContext(ctx, "send message",
		slog.String("form", message.GetForm()),
		slog.String("text", message.GetText()),
		slog.Time("timestamp", message.GetTimestamp().AsTime()),
	)

	return nil, nil
}
