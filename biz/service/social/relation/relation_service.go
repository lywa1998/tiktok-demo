package relation

import (
	"context"
    "log"

	"github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/common/hlog"

    "tiktok_demo/biz/dal/db"
    common "tiktok_demo/biz/model/common"
    relation "tiktok_demo/biz/model/social/relation"
    user_service "tiktok_demo/biz/service/basic/user"
    "tiktok_demo/pkg/errno"
)

const (
    FOLLOW = 1
    UNFOLLOW = 2
)

type RelationService struct {
    ctx context.Context
    c *app.RequestContext
}

func NewRelationService(ctx context.Context, c *app.RequestContext) *RelationService {
    return &RelationService{ctx: ctx, c: c}
}

// FollowAction follow or unfollow by req
func (r *RelationService) FollowAction(req *relation.RelationActionRequest) (flag bool, err error) {
	_, err = db.CheckUserExistById(req.ToUserID)
	if err != nil {
		return false, err
	}
	if req.ActionType != FOLLOW && req.ActionType != UNFOLLOW {
		return false, errno.ParamErr
	}
	current_user_id, _ := r.c.Get("current_user_id")
	// Not allowed to pay attention to oneself
	if req.ToUserID == current_user_id.(int64) {
		return false, errno.ParamErr
	}
	new_follow_relation := &db.Follows{
		UserId:     req.ToUserID,
		FollowerId: current_user_id.(int64),
	}
	follow_exist, _ := db.QueryFollowExist(new_follow_relation.UserId, new_follow_relation.FollowerId)
	if req.ActionType == FOLLOW {
		if follow_exist {
			return false, errno.FollowRelationAlreadyExistErr
		}
		flag, err = db.AddNewFollow(new_follow_relation)
	} else {
		if !follow_exist {
			return false, errno.FollowRelationNotExistErr
		}
		flag, err = db.DeleteFollow(new_follow_relation)
	}
	return flag, err
}

func (r *RelationService) GetFollowList(req *relation.RelationFollowListRequest) (followerlist []*common.User, err error) {
	var followList []*common.User
	_, err = db.CheckUserExistById(req.UserID)
	if err != nil {
		return followList, err
	}

	current_user_id, exists := r.c.Get("current_user_id")
	if !exists {
		current_user_id = int64(0)
	}
	followIdList, err := db.GetFollowIdList(req.UserID)
	if err != nil {
		return followList, err
	}

	for _, follow := range followIdList {
		user_info, err := user_service.NewUserService(r.ctx, r.c).GetUserInfo(follow, current_user_id.(int64))
		if err != nil {
			continue
		}
		user := common.User{
			ID:              user_info.ID,
			Name:            user_info.Name,
			FollowCount:     user_info.FollowCount,
			FollowerCount:   user_info.FollowerCount,
			IsFollow:        user_info.IsFollow,
			Avatar:          user_info.Avatar,
			BackgroundImage: user_info.BackgroundImage,
			Signature:       user_info.Signature,
			TotalFavorited:  user_info.TotalFavorited,
			WorkCount:       user_info.WorkCount,
			FavoriteCount:   user_info.FavoriteCount,
		}
		followList = append(followList, &user)
	}
	return followList, nil
}

// GetFollowerList get follower list by the user id in req
func (r *RelationService) GetFollowerList(req *relation.RelationFollowerListRequest) ([]*common.User, error) {
	user_id := req.UserID
	var followerList []*common.User
	current_user_id, exists := r.c.Get("current_user_id")
	if !exists {
		current_user_id = int64(0)
	}

	followerIdList, err := db.GetFollowerIdList(user_id)
	if err != nil {
		return followerList, err
	}

	for _, follower := range followerIdList {
		user_info, err := user_service.NewUserService(r.ctx, r.c).GetUserInfo(follower, current_user_id.(int64))
		if err != nil {
			hlog.Error("func error: GetFollowerList -> GetUserInfo")
		}

		user := &common.User{
			ID:              user_info.ID,
			Name:            user_info.Name,
			FollowCount:     user_info.FollowCount,
			FollowerCount:   user_info.FollowerCount,
			IsFollow:        user_info.IsFollow,
			Avatar:          user_info.Avatar,
			BackgroundImage: user_info.BackgroundImage,
			Signature:       user_info.Signature,
			TotalFavorited:  user_info.TotalFavorited,
			WorkCount:       user_info.WorkCount,
			FavoriteCount:   user_info.FavoriteCount,
		}
		followerList = append(followerList, user)
	}
	return followerList, nil
}

// GetFriendList get friend list by the user id in req
func (r *RelationService) GetFriendList(req *relation.RelationFriendListRequest) ([]*relation.FriendUser, error) {
	user_id := req.UserID
	current_user_id, _ := r.c.Get("current_user_id")

	if current_user_id.(int64) != user_id {
		return nil, errno.FriendListNoPermissionErr
	}

	var friendList []*relation.FriendUser

	friendIdList, err := db.GetFriendIdList(user_id)
	if err != nil {
		return friendList, err
	}

	for _, id := range friendIdList {
		user_info, err := user_service.NewUserService(r.ctx, r.c).GetUserInfo(id, user_id)
		if err != nil {
			log.Printf("func error: GetFriendList -> GetUserInfo")
		}
		message, err := db.GetLatestMessageByIdPair(user_id, id)
		if err != nil {
			log.Printf("func error: GetFriendList -> GetLatestMessageByIdPair")
		}
		var msgType int64
		if message == nil { // No chat history
			msgType = 2
			message = &db.Messages{}
		} else if user_id == message.FromUserId {
			msgType = 1
		} else {
			msgType = 0
		}
		friendList = append(friendList, &relation.FriendUser{
			User: &common.User{
				ID:              user_info.ID,
				Name:            user_info.Name,
				FollowCount:     user_info.FollowCount,
				FollowerCount:   user_info.FollowerCount,
				IsFollow:        user_info.IsFollow,
				Avatar:          user_info.Avatar,
				BackgroundImage: user_info.BackgroundImage,
				Signature:       user_info.Signature,
				TotalFavorited:  user_info.TotalFavorited,
				WorkCount:       user_info.WorkCount,
				FavoriteCount:   user_info.FavoriteCount,
			},
			Message: message.Content,
			MsgType: msgType,
		})
	}

	return friendList, nil
}
