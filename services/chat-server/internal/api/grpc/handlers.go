package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/api/grpc/converter"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
)

// ServerHandlers implements the desc.ChatServiceV1Server interface.
type ServerHandlers struct {
	desc.UnimplementedChatServiceV1Server

	chatService    service.ChatService
	messageService service.MessageService
}

// New creates a new instance of ServerHandlers with provided chatService and messageService.
func New(chatService service.ChatService, messageService service.MessageService) *ServerHandlers {
	return &ServerHandlers{
		chatService:    chatService,
		messageService: messageService,
	}
}

// CreateChat handles the gRPC request to create a new chat.
func (h *ServerHandlers) CreateChat(ctx context.Context, chat *desc.CreateChatIn) (*desc.CreateChatOut, error) {
	id, err := h.chatService.Save(
		ctx,
		converter.ChatProtoToModel(chat.GetChat()),
	)
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return &desc.CreateChatOut{Id: id}, nil
}

// DeleteChat handles the gRPC request to delete a chat by its ID.
func (h *ServerHandlers) DeleteChat(ctx context.Context, chat *desc.DeleteChatIn) (*emptypb.Empty, error) {
	err := h.chatService.Delete(ctx, chat.GetId())
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return nil, nil
}

// SendMessage handles the gRPC request to send a message.
func (h *ServerHandlers) SendMessage(ctx context.Context, message *desc.SendMessageIn) (*desc.SendMessageOut, error) {
	id, err := h.messageService.Send(
		ctx,
		converter.MessageProtoToModel(message.GetMessage()),
	)
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return &desc.SendMessageOut{Uuid: id}, nil
}
