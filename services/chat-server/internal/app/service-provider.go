package app

import (
	"github.com/alisher-baizhumanov/chat-microservices/pkg/client/mongo"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
	chatRepository "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/chat"
	messageRepository "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/message"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
	chatService "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service/chat"
	messageService "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service/message"
)

type serviceProvider struct {
	_mongoDatabase  mongo.Client
	_chatRepo       repository.ChatRepository
	_messageRepo    repository.MessageRepository
	_chatService    service.ChatService
	_messageService service.MessageService
	_gRPCServer     *grpc.ServerHandlers
}

func newServiceProvider(mongoClient mongo.Client) serviceProvider {
	return serviceProvider{_mongoDatabase: mongoClient}
}

func (s *serviceProvider) MongoDatabase() mongo.Client {
	return s._mongoDatabase
}

func (s *serviceProvider) ChatRepository() repository.ChatRepository {
	if s._chatRepo == nil {
		db := s.MongoDatabase()

		s._chatRepo = chatRepository.New(
			db.Collection(chatRepository.CollectionChat),
			db.Collection(chatRepository.CollectionParticipants),
		)
	}

	return s._chatRepo
}

func (s *serviceProvider) MessageRepository() repository.MessageRepository {
	if s._messageRepo == nil {
		db := s.MongoDatabase()

		s._messageRepo = messageRepository.New(
			db.Collection(messageRepository.CollectionMessages),
		)
	}

	return s._messageRepo
}

func (s *serviceProvider) ChatService() service.ChatService {
	if s._chatService == nil {
		s._chatService = chatService.New(
			s.ChatRepository(),
		)
	}

	return s._chatService
}

func (s *serviceProvider) MessageService() service.MessageService {
	if s._messageService == nil {
		s._messageService = messageService.New(
			s.MessageRepository(),
		)
	}

	return s._messageService
}

func (s *serviceProvider) ServerHandlers() *grpc.ServerHandlers {
	if s._gRPCServer == nil {
		s._gRPCServer = grpc.New(
			s.ChatService(),
			s.MessageService(),
		)
	}

	return s._gRPCServer
}
