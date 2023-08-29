package service

import (
	"context"
    "log"
	"tiktok_demo/biz/dal/db"
	"tiktok_demo/biz/model/common"
	"tiktok_demo/biz/model/interact/comment"
	"tiktok_demo/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"

	user_service "tiktok_demo/biz/service/basic/user"
)

type CommentService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewCommentService create comment service
func NewCommentService(ctx context.Context, c *app.RequestContext) *CommentService {
	return &CommentService{ctx: ctx, c: c}
}

// AddNewComment add a comment and return comment if success
func (c *CommentService) AddNewComment(req *comment.CommentActionRequest) (*comment.Comment, error) {
	current_user_id, _ := c.c.Get("current_user_id")
	video_id := req.VideoID
	action_type := req.ActionType
	comment_text := req.CommentText
	comment_id := req.CommentID
	comment := &comment.Comment{}
	// 发表评论
	if action_type == 1 {
		db_comment := &db.Comment{
			UserId:      current_user_id.(int64),
			VideoId:     video_id,
			CommentText: comment_text,
		}
		err := db.AddNewComment(db_comment)
		if err != nil {
			return comment, err
		}
		comment.ID = db_comment.ID
		comment.CreateDate = db_comment.CreatedAt.Format("01-02")
		comment.Content = db_comment.CommentText
		comment.User, err = c.getUserInfoById(current_user_id.(int64), current_user_id.(int64))
		if err != nil {
			return comment, err
		}
		return comment, nil
	} else {
		err := db.DeleteCommentById(comment_id)
		if err != nil {
			return comment, err
		}
		return comment, nil
	}
}

// getUserInfoById get common.User by user id via user service
func (c *CommentService) getUserInfoById(current_user_id, user_id int64) (*common.User, error) {
	u, err := user_service.NewUserService(c.ctx, c.c).GetUserInfo(user_id, current_user_id)
	var comment_user *common.User
	if err != nil {
		return comment_user, err
	}
	comment_user = &common.User{
		ID:              u.ID,
		Name:            u.Name,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
	return comment_user, nil
}

func (c *CommentService) CommentList(req *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	resp := &comment.CommentListResponse{}
	video_id := req.VideoID

	// get current_user_id
	current_user_id, _ := c.c.Get("current_user_id")

	dbcomments, err := db.GetCommentListByVideoID(video_id)
	if err != nil {
		return resp, err
	}
	var comments []*comment.Comment
	err = c.copyComment(&comments, &dbcomments, current_user_id.(int64))
	if err != nil {
		return resp, err
	}
	resp.CommentList = comments
	resp.StatusMsg = errno.SuccessMsg
	resp.StatusCode = errno.SuccessCode
	return resp, nil
}

// copyComment convert comment list from db to model
func (c *CommentService) copyComment(result *[]*comment.Comment, data *[]*db.Comment, current_user_id int64) error {
	for _, item := range *data {
		comment := c.createComment(item, current_user_id)
		*result = append(*result, comment)
	}
	return nil
}

// createComment convert single comment from db to model
func (c *CommentService) createComment(data *db.Comment, userId int64) *comment.Comment {
	comment := &comment.Comment{
		ID:         data.ID,
		Content:    data.CommentText,
		CreateDate: data.CreatedAt.Format("01-02"),
	}

	user_info, err := c.getUserInfoById(userId, data.UserId)
	if err != nil {
		log.Printf("func error")
	}
	comment.User = user_info
	return comment
}
