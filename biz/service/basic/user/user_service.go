package service

import (
	"context"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"

	"tiktok_demo/biz/dal/db"
	"tiktok_demo/biz/model/basic/user"
	"tiktok_demo/biz/model/common"
	"tiktok_demo/pkg/constants"
	"tiktok_demo/pkg/errno"
	"tiktok_demo/pkg/utils"
)

type UserService struct {
    ctx context.Context
    c *app.RequestContext
}

// NewUserService create user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
    return &UserService{ ctx: ctx, c: c }
}

// UserRegister register user return user id
func (s *UserService) UserRegister(req *user.UserRegisterRequest) (user_id int64, err error) {
    user, err := db.QueryUser(req.Name)
    if err != nil {
        return 0, err
    }
    if *user != (db.User{}) {
        return 0, errno.UserAlreadyExistErr
    }

    password, err := utils.Crypt(req.Password)
    user_id, err = db.CreateUser(&db.User{
        UserName: req.Name,
        Password: password,
        Avatar: constants.DefaultAva,
        BackgroundImage: constants.DefaultBackground,
        Signature: constants.DefaultSign,
    })
    return user_id, nil
}

// UserInfo the function of user api
func (s *UserService) UserInfo(req *user.UserRequest) (*common.User, error) {
    query_user_id := req.UserID
    current_user_id, exists := s.c.Get("current_user_id")
    if !exists {
        current_user_id = 0
    }
    return s.GetUserInfo(query_user_id, current_user_id.(int64))
}

func (s *UserService) GetUserInfo(query_user_id, current_user_id int64) (*common.User, error) {
    u := common.NewUser()
    errChan := make(chan error, 7)
    defer close(errChan)
    var wg sync.WaitGroup
    wg.Add(7)

    // Query u.Name, u.Avatar, u.BackgroundImage and u.Signature from database
    go func() {
        defer wg.Done()
        dbUser, err := db.QueryUserById(query_user_id)
        if err != nil {
            errChan <- err
        } else {
            u.Name = dbUser.UserName
            u.Avatar = utils.URLconvert(s.ctx, s.c, dbUser.Avatar)
            u.BackgroundImage = utils.URLconvert(s.ctx, s.c, dbUser.BackgroundImage)
            u.Signature = dbUser.Signature
        }
    }()
    
    // Query u.WorkCount from database
    go func ()  {
        defer wg.Done()
        WorkCount, err := db.GetWorkCount(query_user_id)
        if err != nil {
            errChan <- err
        } else {
            u.WorkCount = WorkCount
        }
    }()

    // Query u.FollowCount from database
    go func ()  {
        defer wg.Done()
        FollowCount, err := db.GetFollowCount(query_user_id)
        if err != nil {
            errChan <- err
        } else {
            u.FollowCount = FollowCount
        }
    }()

    // Query u.FollwerCount from database
    go func ()  {
        defer wg.Done()
        FollwerCount, err := db.GetFollowerCount(query_user_id)
        if err != nil {
            errChan <- err
        } else {
            u.FollowerCount = FollwerCount
        }
    }()

    // Query u.IsFollow from database
    go func ()  {
        defer wg.Done()
        if current_user_id != 0 {
            IsFollow, err := db.QueryFollowExist(current_user_id, query_user_id)
            if err != nil {
                errChan <- err
            } else {
                u.IsFollow = IsFollow
            }
        } else {
            u.IsFollow = false
        }
    }()

    // Query u.FavoriteCount from database
    go func ()  {
        defer wg.Done()
        FavoriteCount, err := db.GetFavoriteCountByUserID(query_user_id)
        if err != nil {
            errChan <- err
        } else {
            u.FavoriteCount = FavoriteCount
        }
    }()

    // Query u.TotalFavorited from database
    go func ()  {
        defer wg.Done()
        TotalFavorited, err := db.QueryTotalFavoritedByAuthorID(query_user_id)
        if err != nil {
            errChan <- err
        } else {
            u.TotalFavorited = TotalFavorited
        }
    }()

    wg.Wait()
    select {
    case result := <-errChan:
        return &common.User{}, result
    default:
    }
    u.ID = query_user_id
    return u, nil
}
