package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alisher-baizhumanov/chat-microservices/protos/generated/auth-v1"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/api/grpc/converter"
	"github.com/alisher-baizhumanov/chat-microservices/services/auth/internal/service"
)

// AuthHandlers implements the gRPC server for auth-related operations.
type AuthHandlers struct {
	desc.UnimplementedAuthServiceV1Server

	authService service.AuthService
}

// NewAuthHandlers creates and returns a new UserHandlers instance.
func NewAuthHandlers(authService service.AuthService) *AuthHandlers {
	return &AuthHandlers{authService: authService}
}

// Login handles the gRPC request for user login.
func (a AuthHandlers) Login(ctx context.Context, in *desc.LoginIn) (*desc.LoginOut, error) {
	refreshToken, err := a.authService.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return &desc.LoginOut{RefreshToken: refreshToken}, nil
}

// GetRefreshToken handles the gRPC request for getting a new refresh token.
func (a AuthHandlers) GetRefreshToken(ctx context.Context, in *desc.GetRefreshTokenIn) (*desc.GetRefreshTokenOut, error) {
	refreshToken, err := a.authService.GetRefreshToken(ctx, in.GetRefreshToken())
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return &desc.GetRefreshTokenOut{RefreshToken: refreshToken}, nil
}

// GetAccessToken handles the gRPC request for getting a new access token.
func (a AuthHandlers) GetAccessToken(ctx context.Context, in *desc.GetAccessTokenIn) (*desc.GetAccessTokenOut, error) {
	accessToken, err := a.authService.GetAccessToken(ctx, in.GetRefreshToken())
	if err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return &desc.GetAccessTokenOut{AccessToken: accessToken}, nil
}

// Check handles the gRPC request for checking access to a specific endpoint.
func (a AuthHandlers) Check(ctx context.Context, in *desc.CheckIn) (*emptypb.Empty, error) {
	if err := a.authService.CheckAccess(ctx, in.GetEndpointAddress(), in.GetAccessToken()); err != nil {
		return nil, converter.ErrorModelToProto(err)
	}

	return nil, nil
}
