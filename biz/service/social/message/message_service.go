package message

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"tiktok_demo/biz/dal/db"
	"tiktok_demo/biz/model/social/message"
	"tiktok_demo/pkg/errno"
	"tiktok_demo/pkg/utils"
)

type MessageService struct {
    ctx context.Context
    c *app.RequestContext
}

// NewMessageService create message service
func NewMessageService(ctx context.Context, c *app.RequestContext) *MessageService {
    return &MessageService{ctx: ctx, c: c}
}

// GetMessageChat get chat records
func (s *MessageService) GetMessageChat(req *message.MessageChatRequest) ([]*message.Message, error)  {
    messages := make([]*message.Message, 0)
    from_user_id, _ := s.c.Get("current_user_id")
    to_user_id := req.ToUserID
    pre_msg_time := req.PreMsgTime
    db_messages, err := db.GetMessageByIdPair(from_user_id.(int64), to_user_id, utils.MillTimeStampToTime(pre_msg_time))
    if err != nil {
        return messages, err
    }
    for _, db_message := range db_messages {
        messages = append(messages, &message.Message{
            ID: db_message.ID,
            ToUserID: db_message.ToUserId,
            FromUserID: db_message.FromUserId,
            Content: db_message.Content,
            CreateTime: db_message.CreatedAt.UnixNano() / 1000000,
        })
    }
    return messages, nil
}

// MessageAction add a message
func (s *MessageService) MessageAction(req *message.MessageActionRequest) error {
    from_user_id, _ := s.c.Get("current_user_id")
    to_user_id := req.ToUserID
    content := req.Content

    ok, err := db.AddNewMessage(&db.Messages{
        FromUserId: from_user_id.(int64),
        ToUserId: to_user_id,
        Content: content,
    })
    if err != nil {
        return err
    }
    if !ok {
        return errno.MessageAddFailedErr
    }
    return nil
}
