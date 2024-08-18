package interceptor

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

// Logger is a gRPC interceptor that logs the details of each RPC call.
func Logger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	resp, err := handler(ctx, req)

	end := time.Now()

	logger.Info("rpc call",
		logger.String("method", info.FullMethod),
		logger.Duration("duration", end.Sub(start)),
		logger.Int("request_size", len(fmt.Sprintf("%v", req))),
		logger.Int("response_size", len(fmt.Sprintf("%v", resp))),
	)

	return resp, err
}
