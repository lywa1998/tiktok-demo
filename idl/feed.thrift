namespace go basic.feed

include 'common.thrift'

struct FeedRequest {
    1: i64 latest_time
    2: string token
}

struct FeedResponse {
    1: i32 status_code
    2: string status_msg
    3: list<common.Video> video_list
    4: i64 next_time
}

service FeedHandler {
    FeedResponse Feed(1: FeedRequest request) (api.get="/douyin/feed/")
}
