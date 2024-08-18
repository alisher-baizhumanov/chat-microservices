package app

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc/http-gateway"
	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/chat-v1"
	grpcHandler "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/config"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
	chatService "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service/chat"
	messageService "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service/message"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository"
	chatRepository "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository/chat"
	messageRepository "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository/message"
)

type serviceProvider struct {
	mongoDatabase     mongo.Client
	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository
	chatService       service.ChatService
	messageService    service.MessageService
	gRPCHandlers      *grpcHandler.ServerHandlers
	cfg               *config.Config
}

func newServiceProvider(mongoClient mongo.Client, cfg *config.Config) serviceProvider {
	return serviceProvider{mongoDatabase: mongoClient, cfg: cfg}
}

func (s *serviceProvider) getConfig() *config.Config {
	return s.cfg
}

func (s *serviceProvider) getMongoDatabase() mongo.Client {
	return s.mongoDatabase
}

func (s *serviceProvider) getChatRepository() repository.ChatRepository {
	if s.chatRepository == nil {
		db := s.getMongoDatabase()

		s.chatRepository = chatRepository.New(
			db.Collection(chatRepository.CollectionChat),
			db.Collection(chatRepository.CollectionParticipants),
		)
	}

	return s.chatRepository
}

func (s *serviceProvider) getMessageRepository() repository.MessageRepository {
	if s.messageRepository == nil {
		db := s.getMongoDatabase()

		s.messageRepository = messageRepository.New(
			db.Collection(messageRepository.CollectionMessages),
		)
	}

	return s.messageRepository
}

func (s *serviceProvider) getChatService() service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.New(
			s.getChatRepository(),
		)
	}

	return s.chatService
}

func (s *serviceProvider) getMessageService() service.MessageService {
	if s.messageService == nil {
		s.messageService = messageService.New(
			s.getMessageRepository(),
		)
	}

	return s.messageService
}

func (s *serviceProvider) getGRPCHandlers() *grpcHandler.ServerHandlers {
	if s.gRPCHandlers == nil {
		s.gRPCHandlers = grpcHandler.New(
			s.getChatService(),
			s.getMessageService(),
		)
	}

	return s.gRPCHandlers
}

func (s *serviceProvider) getGRPCServer() (*grpc.Server, error) {
	return grpc.NewGRPCServer(
		s.getConfig().GRPCServerPort,
		[]grpc.Service{
			{
				ServiceDesc: &desc.ChatServiceV1_ServiceDesc,
				Handler:     s.getGRPCHandlers(),
			},
		},
	)
}

func (s *serviceProvider) getHTTPServer(ctx context.Context) (*http.Server, error) {
	mux := runtime.NewServeMux()

	if err := desc.RegisterChatServiceV1HandlerFromEndpoint(
		ctx,
		mux,
		s.getConfig().GRPCAddress(),
		s.getConfig().GRPCClientDialOptions(),
	); err != nil {
		return nil, err
	}

	return http.NewHTTPServer(
		s.getConfig().HTTPServerPort,
		mux,
	), nil
}
