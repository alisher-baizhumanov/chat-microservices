package chat

import "context"

func (s *service) Delete(ctx context.Context, id string) error {
	return s.chatRepo.DeleteChatByID(ctx, id)
}
