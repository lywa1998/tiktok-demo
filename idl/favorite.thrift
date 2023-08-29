namespace go interact.favorite

include 'common.thrift'

struct FavoriteActionRequest {
    1: string   token
    2: i64      video_id
    3: i32      action_type
}

struct FavoriteActionResponse {
    1: i32      status_code
    2: string   status_msg
}

struct FavoriteListRequest {
    1: i64      user_id
    2: string   token
}

struct FavoriteListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<common.Video> video_list
}

service FavoriteHandler {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest request) (api.post="/douyin/favorite/action/")

    FavoriteListResponse FavoriteList(1: FavoriteListRequest request) (api.get="/douyin/favorite/list/")
}
