package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

// Recover is a gRPC interceptor that recovers from panics and logs the error details.
func Recover(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (result any, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("recovered from panic",
				logger.String("error", fmt.Sprintf("%v", r)),
				logger.String("method", info.FullMethod),
			)

			result = nil
			err = status.Error(codes.Internal, "")
		}
	}()

	return handler(ctx, req)
}
