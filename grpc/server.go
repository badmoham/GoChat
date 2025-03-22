package grpc

import (
	"context"

	"GoChat/proto/chat"
	"GoChat/services"
)

type ChatServer struct {
	chat.UnimplementedChatServiceServer
}

func NewChatServer() *ChatServer {
	return &ChatServer{}
}

func (s *ChatServer) ListMessages(ctx context.Context, req *chat.ChatRequest) (*chat.MessageList, error) {
	requesterUserID, err := services.DecryptJWT(req.JwtToken)
	if err != nil {
		return nil, err
	}
	messages, err := services.GetUserChatRoomMessages(requesterUserID, uint(req.ChatRoomId), int(req.Limit), int((req.Page-1)*req.Limit))
	if err != nil {
		return nil, err
	}

	protoMessages := make([]*chat.Message, 0, len(messages))
	for _, m := range messages {
		protoMessages = append(protoMessages, &chat.Message{
			Id:         uint64(m.ID),
			Type:       uint64(m.Type),
			Text:       m.Text,
			UserId:     uint64(m.UserID),
			ChatRoomId: uint64(m.ChatRoomID),
			CreatedAt:  m.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:  m.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}
	return &chat.MessageList{Messages: protoMessages}, nil
}
