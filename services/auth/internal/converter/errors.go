package converter

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/model"
)

// ErrorModelToProto converts a given error from the internal model to a gRPC status error.
// It maps specific internal errors to corresponding gRPC error codes.
func ErrorModelToProto(err error) error {
	if err == nil {
		return nil
	}

	message := err.Error()
	switch {
	case errors.Is(err, model.ErrCanNotBeNil):
		return status.Error(codes.InvalidArgument, message)
	case errors.Is(err, model.ErrInvalidID):
		return status.Errorf(codes.InvalidArgument, message)
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
