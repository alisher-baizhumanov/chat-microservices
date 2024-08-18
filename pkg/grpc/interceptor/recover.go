package interceptor

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

// Recover is a gRPC interceptor that recovers from panics and logs the error details.
func Recover(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (result any, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("recovered from panic",
				slog.String("error", fmt.Sprintf("%v", r)),
				slog.String("method", info.FullMethod),
			)
		}

		//result = nil
		//err = status.Error(codes.Internal, "internal server error")
	}()

	return handler(ctx, req)
}
