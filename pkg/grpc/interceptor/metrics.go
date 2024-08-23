package interceptor

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/alisher-baizhumanov/chat-microservices/pkg/logger"
)

// Logging is a gRPC interceptor that logs the details of each RPC call.
func Logging(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	resp, err := handler(ctx, req)

	duration := time.Since(start)
	var reqSize, respSize int
	if protoReq, ok := req.(proto.Message); ok {
		reqSize = proto.Size(protoReq)
	}

	if protoResp, ok := resp.(proto.Message); ok {
		respSize = proto.Size(protoResp)
	}

	code := errorCode(err).String()

	logging(info.FullMethod, reqSize, respSize, duration, code)

	return resp, err
}

func errorCode(err error) codes.Code {
	if err == nil {
		return codes.OK
	}

	errStatus, ok := status.FromError(err)
	if !ok {
		return codes.Unknown
	}

	return errStatus.Code()
}

func logging(method string, reqSize, respSize int, duration time.Duration, code string) {
	logger.Info("rpc call",
		logger.String("method", method),
		logger.Duration("duration", duration),
		logger.Int("request_size", reqSize),
		logger.Int("response_size", respSize),
		logger.String("code", code),
	)
}
