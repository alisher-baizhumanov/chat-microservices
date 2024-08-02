package app

import (
	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository"
	chatRepository "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository/chat"
	messageRepository "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/storage/repository/message"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
	chatService "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service/chat"
	messageService "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service/message"
)

type serviceProvider struct {
	mongoDatabase     mongo.Client
	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository
	chatService       service.ChatService
	messageService    service.MessageService
	gRPCServer        *grpc.ServerHandlers
}

func newServiceProvider(mongoClient mongo.Client) serviceProvider {
	return serviceProvider{mongoDatabase: mongoClient}
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

func (s *serviceProvider) getServerHandlers() *grpc.ServerHandlers {
	if s.gRPCServer == nil {
		s.gRPCServer = grpc.New(
			s.getChatService(),
			s.getMessageService(),
		)
	}

	return s.gRPCServer
}
