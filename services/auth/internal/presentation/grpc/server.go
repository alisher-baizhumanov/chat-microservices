package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/user-v1"
)

// Server represents the gRPC server with its listener and server instance.
type Server struct {
	gRPCServer *grpc.Server
	listener   net.Listener
}

// Start runs the gRPC server in a separate goroutine to handle incoming requests.
func (s *Server) Start() {
	go func() {
		_ = s.gRPCServer.Serve(s.listener)
	}()
}

// Stop gracefully stops the gRPC server, ensuring all ongoing requests are completed.
func (s *Server) Stop() {
	s.gRPCServer.GracefulStop()
}

// NewGRPCServer creates and returns a new Server instance listening on the specified port.
// It also registers the user service and reflection service to the gRPC server.
func NewGRPCServer(port int) (*Server, error) {
	var server Server

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server.listener = listener

	gRPCServer := grpc.NewServer()
	reflection.Register(gRPCServer)
	desc.RegisterUserServiceV1Server(gRPCServer, &userServer{})

	return &Server{
		gRPCServer: gRPCServer,
		listener:   listener,
	}, nil
}
