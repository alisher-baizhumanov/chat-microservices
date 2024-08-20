package interceptor

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

// Metrics is a gRPC interceptor that logs the details of each RPC call.
func Metrics(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	resp, err := handler(ctx, req)

	var reqSize, respSize int
	if protoReq, ok := req.(proto.Message); ok {
		reqSize = proto.Size(protoReq)
	}

	if protoResp, ok := resp.(proto.Message); ok {
		respSize = proto.Size(protoResp)
	}

	logger.Info("rpc call",
		logger.String("method", info.FullMethod),
		logger.Duration("duration", time.Since(start)),
		logger.Int("request_size", reqSize),
		logger.Int("response_size", respSize),
	)

	return resp, err
}
