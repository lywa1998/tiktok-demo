package redis

import (
	"time"

	"github.com/go-redis/redis/v7"

	"tiktok_demo/pkg/constants"
)

var (
    expireTime = time.Hour * 1
    rdbFollows *redis.Client
    rdbFavorite *redis.Client
)

func Init() {
    rdbFollows = redis.NewClient(&redis.Options{
        Addr: constants.RedisAddr,
        Password: constants.RedisPassword,
        DB: 0,
    })
    rdbFavorite = redis.NewClient(&redis.Options{
        Addr: constants.RedisAddr,
        Password: constants.RedisPassword,
        DB: 1,
    })
}
