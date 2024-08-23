package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/grpc/interceptor"
)

// DefaultOptions is a set of default gRPC server options that include:
// - Insecure credentials (no TLS)
// - A unary interceptor that validates incoming requests
var (
	DefaultOptions = []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			interceptor.Recover,
			interceptor.Logging,
			interceptor.ValidateGRPC,
		),
	}
)
