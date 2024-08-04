package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// DefaultOptions is a set of default gRPC server options that include:
// - Insecure credentials (no TLS)
// - A unary interceptor that validates incoming requests
var (
	DefaultOptions = []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(ValidateInterceptor),
	}
)
