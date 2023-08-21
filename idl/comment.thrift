namespace go interact.comment

include 'common.thrift'

struct Comment {
    1: i64      id          // video comment id
    2: common.User     user        // comment user information
    3: string   content     // comment
    4: string   create_date // comment publication data, format mm-dd
}

struct CommentActionRequest {
    1: string   token           // user authentication token
    2: i64      video_id        // video id
    3: i32      action_type     // 1: post a comment, 2: delete a comment
    4: string   comment_text    // comment content filled in by users, when action_type = 1
    5: i64      comment_id      // id of the comment to delete, when action_type = 2
}

struct CommentActionResponse {
    1: i32      status_code
    2: string   status_msg
    3: Comment  comment
}

struct CommentListRequest {
    1: i32      token
    2: i64      video_id
}

struct CommentListResponse {
    1: i32      status_code
    2: string   status_msg
    3: list<Comment>  comment_list
}

service CommentHandler {
    CommentActionResponse CommentAction(1: CommentActionRequest request) (api.post="/douyin/comment/action/")

    CommentListResponse CommentList(1: CommentListRequest request) (api.get="/douyin/comment/list/")
}
