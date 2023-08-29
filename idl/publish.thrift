namespace go basic.publish

include 'common.thrift'

struct PublishActionRequest {
    1: string token
    2: binary data
    3: string title
}

struct PublishActionResponse {
    1: i32 status_code
    2: string status_msg
}

struct PublishListRequest {
    1: i64 user_id
    2: string token
}

struct PublishListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<common.Video> video_list
}

service PublishHanler {
    PublishListResponse PublishList(1: PublishListRequest request) (api.get="/douyin/publish/list/")

    PublishActionResponse PublishAction(1: PublishActionRequest request) (api.post="/douyin/publish/action/")
}
