package service

import (
	"context"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"tiktok_demo/biz/dal/db"
	"tiktok_demo/biz/model/basic/publish"
	"tiktok_demo/biz/model/common"
	"tiktok_demo/biz/mw/ffmpeg"
	"tiktok_demo/biz/mw/oss"
	feed_service "tiktok_demo/biz/service/basic/feed"
	"tiktok_demo/pkg/constants"
	"tiktok_demo/pkg/utils"
)

type PublishService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewPublishService create publish service
func NewPublishService(ctx context.Context, c *app.RequestContext) *PublishService {
	return &PublishService{ctx: ctx, c: c}
}

// PublishAction put file to OSS by the FileHeader in the req and store the bucket name and file name
func (s *PublishService) PublishAction(req *publish.PublishActionRequest) (err error) {
	v, _ := s.c.Get("current_user_id")
	title := s.c.PostForm("title")
	user_id := v.(int64)
	nowTime := time.Now()
	filename := utils.NewFileName(user_id, nowTime.Unix())
	req.Data.Filename = filename + path.Ext(req.Data.Filename)
	err = oss.PutToBucket(constants.OSSVideoBucketName, req.Data)
    if err != nil {
        hlog.CtxInfof(s.ctx, "err:"+err.Error())
        return
    }
	hlog.CtxInfof(s.ctx, "video upload success")
	PlayURL := constants.OSSVideoBucketName + "/" + req.Data.Filename
	buf, err := ffmpeg.GetSnapshot(utils.URLconvert(s.ctx, s.c, PlayURL))
	err = oss.PutToBucketByBuf(constants.OSSImgBucketName+filename+".png", buf)
	hlog.CtxInfof(s.ctx, "image upload success")
	if err != nil {
		hlog.CtxInfof(s.ctx, "err:"+err.Error())
        return
	}
	_, err = db.CreateVideo(&db.Video{
		AuthorID:    user_id,
		PlayURL:     PlayURL,
		CoverURL:    constants.OSSImgBucketName + "/" + filename + ".png",
		PublishTime: nowTime,
		Title:       title,
	})
	return err
}

// PublishList get the video list of user
func (s *PublishService) PublishList(req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp = &publish.PublishListResponse{}
	query_user_id := req.UserID
	current_user_id, exist := s.c.Get("current_user_id")
	if !exist {
		current_user_id = int64(0)
	}
	dbVideos, err := db.GetVideoByUserID(query_user_id)
	if err != nil {
		return resp, err
	}
	var videos []*common.Video

	f := feed_service.NewFeedService(s.ctx, s.c)
	err = f.CopyVideos(&videos, &dbVideos, current_user_id.(int64))
	if err != nil {
		return resp, err
	}
	for _, item := range videos {
		video := common.Video{
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
		resp.VideoList = append(resp.VideoList, &video)
	}
	return resp, nil
}
