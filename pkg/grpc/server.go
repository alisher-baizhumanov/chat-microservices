package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server represents the gRPC server with its listener and server instance.
type Server struct {
	gRPCServer *grpc.Server
	listener   net.Listener
}

// Start runs the gRPC server in a separate goroutine to handle incoming requests.
func (s *Server) Start() {
	go func() {
		// We can ignore error because
		// Serve will return a non-nil error unless Stop or GracefulStop is called.
		_ = s.gRPCServer.Serve(s.listener)
	}()
}

// Stop gracefully stops the gRPC server, ensuring all ongoing requests are completed.
func (s *Server) Stop() {
	s.gRPCServer.GracefulStop()
}

// NewGRPCServer creates and returns a new Server instance listening on the specified port.
// It also registers the user service and reflection service to the gRPC server.
func NewGRPCServer(port int, serviceDesc *grpc.ServiceDesc, handlers any, opt ...grpc.ServerOption) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	if len(opt) == 0 {
		opt = DefaultOptions
	}

	gRPCServer := grpc.NewServer(opt...)
	reflection.Register(gRPCServer)
	gRPCServer.RegisterService(serviceDesc, handlers)

	return &Server{
		gRPCServer: gRPCServer,
		listener:   listener,
	}, nil
}
