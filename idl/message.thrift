namespace go social.message

struct Message {
    1: i64 id
    2: i64 to_user_id
    3: i64 from_user_id
    4: string content
    5: i64 create_time
}

struct MessageChatRequest {
    1: string token
    2: i64 to_user_id
    3: i64 pre_msg_time
}

struct MessageChatResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Message> message_list
}

struct MessageActionRequest {
    1: string token
    2: i64 to_user_id
    3: i32 action_type
    4: string content
}

struct MessageActionResponse {
    1: i32 status_code
    2: string status_msg
}

service MessageHandler {
    MessageChatResponse MessageChat(1: MessageChatRequest request) (api.get="/douyin/message/chat/")

    MessageActionResponse MessageAction(1: MessageActionRequest request) (api.post="/douyin/message/action/")
}
