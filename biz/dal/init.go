package dal

import (
    "tiktok_demo/biz/dal/db"
    "tiktok_demo/biz/mw/redis"
)

func Init() {
    db.Init() // mysql Init
    redis.Init() // redis Init
}
