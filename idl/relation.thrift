namespace go social.relation

include 'common.thrift'

struct RelationActionRequest {
    1: string token
    2: i64 to_user_id
    3: i32 action_type
}

struct RelationActionResponse {
    1: i32 status_code
    2:string status_msg
}

struct RelationFollowListRequest {
    1: i64 user_id
    2: string token
}

struct RelationFollowListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<common.User> user_list
}

struct RelationFollowerListRequest {
    1: i64 user_id
    2: string token
}

struct RelationFollowerListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<common.User> user_list
}

struct FriendUser {
    1: common.User user
    2: string message
    3: i64 msg_type
}

struct RelationFriendListRequest {
    1: i64 user_id
    2: string token
}

struct RelationFriendListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<FriendUser> user_list
}

service RelationHandler {
    RelationActionResponse RelationAction(1: RelationActionRequest request) (api.post="/douyin/relation/action/")

    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest request) (api.get="/douyin/relation/follow/list/")

    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest request) (api.get="/douyin/relation/follower/list/")

    RelationFriendListResponse RelationFriendList(1: RelationFriendListRequest request) (api.get="/douyin/relation/friend/list/")
}
