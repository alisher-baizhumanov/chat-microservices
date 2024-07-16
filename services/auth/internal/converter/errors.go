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
	case errors.Is(err, model.ErrInvalidID):
		return status.Error(codes.InvalidArgument, message)
	case errors.Is(err, model.ErrInvalidSQLQuery):
		return status.Error(codes.InvalidArgument, message)
	case errors.Is(err, model.ErrNonUniqueEmail):
		return status.Error(codes.InvalidArgument, "this email is already taken")
	case errors.Is(err, model.ErrNonUniqueUsername):
		return status.Error(codes.InvalidArgument, "this name is already taken")
	case errors.Is(err, model.ErrNotFound):
		return status.Error(codes.NotFound, "")

	case errors.Is(err, model.ErrDatabase):
		return status.Error(codes.Internal, message)
	case errors.Is(err, model.ErrCanNotBeNil):
		return status.Error(codes.InvalidArgument, "pointer does not exist")
	default:
		return status.Errorf(codes.Internal, "unknown error; message=\"%s\"", message)
	}
}
