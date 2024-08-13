package auth

import (
	"context"
	"log/slog"
)

func (s *service) CheckAccess(ctx context.Context, path, accessToken string) error {
	claims, err := s.tokenManager.Verify(accessToken)
	if err != nil {
		return err
	}

	slog.InfoContext(ctx, "checking access",
		slog.String("path", path),
		slog.Int64("user_id", claims.ID),
		slog.String("role", claims.Role.String()),
	)

	return nil
}
