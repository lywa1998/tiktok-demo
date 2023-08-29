package service

import (
	"context"
	"tiktok_demo/biz/dal/db"
	"tiktok_demo/biz/model/common"
	"tiktok_demo/pkg/constants"
	"tiktok_demo/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"

	favorite "tiktok_demo/biz/model/interact/favorite"
	feed_service "tiktok_demo/biz/service/basic/feed"
)

type FavoriteService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewFavoriteService create favorite service
func NewFavoriteService(ctx context.Context, c *app.RequestContext) *FavoriteService {
	return &FavoriteService{ctx: ctx, c: c}
}

// FavoriteAction like a video and return result
func (r *FavoriteService) FavoriteAction(req *favorite.FavoriteActionRequest) (flag bool, err error) {
	_, err = db.CheckVideoExistById(req.VideoID)
	if err != nil {
		return false, err
	}
	if req.ActionType != constants.FavoriteActionType && req.ActionType != constants.UnFavoriteActionType {
		return false, errno.ParamErr
	}
	current_user_id, _ := r.c.Get("current_user_id")

	new_favorite_relation := &db.Favorites{
		UserId:  current_user_id.(int64),
		VideoId: req.VideoID,
	}
	favorite_exist, _ := db.QueryFavoriteExist(new_favorite_relation.UserId, new_favorite_relation.VideoId)
	if req.ActionType == constants.FavoriteActionType {
		if favorite_exist {
			return false, errno.FavoriteRelationAlreadyExistErr
		}
		flag, err = db.AddNewFavorite(new_favorite_relation)
	} else {
		if !favorite_exist {
			return false, errno.FavoriteRelationNotExistErr
		}
		flag, err = db.DeleteFavorite(new_favorite_relation)
	}
	return flag, err
}

// GetFavoriteList query favorite list by the user id in the request
func (r *FavoriteService) GetFavoriteList(req *favorite.FavoriteListRequest) (favoritelist []*common.Video, err error) {
	query_user_id := req.UserID
	_, err = db.CheckUserExistById(query_user_id)

	if err != nil {
		return nil, err
	}
	current_user_id, _ := r.c.Get("current_user_id")

	video_id_list, err := db.GetFavoriteIdList(query_user_id)

	dbVideos, err := db.GetVideoListByVideoIDList(video_id_list)
	var videos []*common.Video
	f := feed_service.NewFeedService(r.ctx, r.c)
	err = f.CopyVideos(&videos, &dbVideos, current_user_id.(int64))
	for _, item := range videos {
		video := &common.Video{
			ID: item.ID,
			Author: &common.User{
				ID:              item.Author.ID,
				Name:            item.Author.Name,
				FollowCount:     item.Author.FollowCount,
				FollowerCount:   item.Author.FollowerCount,
				Avatar:          item.Author.Avatar,
				BackgroundImage: item.Author.BackgroundImage,
				Signature:       item.Author.Signature,
				TotalFavorited:  item.Author.TotalFavorited,
				WorkCount:       item.Author.WorkCount,
			},
			PlayURL:       item.PlayURL,
			CoverURL:      item.CoverURL,
			FavoriteCount: item.FavoriteCount,
			CommentCount:  item.CommentCount,
			IsFavorite:    item.IsFavorite,
			Title:         item.Title,
		}
		favoritelist = append(favoritelist, video)
	}
	return favoritelist, err
}
