package app

import (
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/api/grpc"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository"
	chatRepository "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/chat"
	messageRepository "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/repository/message"
	"github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service"
	chatService "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service/chat"
	messageService "github.com/alisher-baizhumanov/chat-microservices/services/chat-server/internal/service/message"
)

type serviceProvider struct {
	_mongoClient    any
	_chatRepo       repository.ChatRepository
	_messageRepo    repository.MessageRepository
	_chatService    service.ChatService
	_messageService service.MessageService
	_gRPCServer     *grpc.ServerHandlers
}

func newServiceProvider(mongoClient any) serviceProvider {
	return serviceProvider{_mongoClient: mongoClient}
}

func (s *serviceProvider) MongoClient() any {
	return nil
}

func (s *serviceProvider) ChatRepository() repository.ChatRepository {
	if s._chatRepo == nil {
		s._chatRepo = chatRepository.New(s.MongoClient())
	}

	return s._chatRepo
}

func (s *serviceProvider) MessageRepository() repository.MessageRepository {
	if s._messageRepo == nil {
		s._messageRepo = messageRepository.New(s.MongoClient())
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
