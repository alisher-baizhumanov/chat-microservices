package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

// ValidateInterceptor is a gRPC unary server interceptor that validates the request
// object if it implements the validator interface. If the validation fails, the
// interceptor returns the error, otherwise it calls the next handler in the chain.
func ValidateInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if val, ok := req.(validator); ok {
		if err := val.Validate(); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
