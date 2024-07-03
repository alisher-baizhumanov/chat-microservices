package grpc

import (
	"context"
	"log/slog"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"
)

type chatServer struct {
	desc.UnimplementedChatServiceV1Server
}

func (s *chatServer) Create(ctx context.Context, createIn *desc.CreateIn) (*desc.CreateOut, error) {
	usernames := createIn.Usernames

	slog.InfoContext(ctx, "created chat", slog.Any("usernames", usernames))

	return &desc.CreateOut{Id: gofakeit.Int64()}, nil
}

func (s *chatServer) Delete(ctx context.Context, deleteIn *desc.DeleteIn) (*emptypb.Empty, error) {
	id := deleteIn.Id

	slog.InfoContext(ctx, "deleted chat", slog.Int64("id", id))

	return nil, nil
}

func (s *chatServer) SendMessage(ctx context.Context, sendMessageIn *desc.SendMessageIn) (*emptypb.Empty, error) {
	message := sendMessageIn.Message

	timeStamp := time.Unix(message.Timestamp.Seconds, 0)

	slog.InfoContext(ctx, "send message",
		slog.String("form", message.Form),
		slog.String("text", message.Text),
		slog.Time("timestamp", timeStamp),
	)

	return nil, nil
}
