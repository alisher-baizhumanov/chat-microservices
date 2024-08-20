package grpc

import "google.golang.org/grpc"

// Service represents a gRPC service with its description and handler.
type Service struct {
	ServiceDesc *grpc.ServiceDesc
	Handler     any
}
