package constants

const (
    MySQLDefaultDSN = "douyin:douyin123@tcp(127.0.0.1:18000)/douyin?charset=utf8&parseTime=True&loc=Local"

    OSSEndPoint = "oss-cn-shenzhen.aliyuncs.com"
    OSSAccessKeyID = "LTAI5tM5k7UEXB7nsFJ6Cenz"
    OSSAccessKeySecret= "FKFY4MIVvpfRvXagnPz8WDcQTjfdEF"

    RedisAddr = "localhost:18003"
    RedisPassword= "douyin123"
)

const (
    UserTableName      = "users"
	FollowsTableName   = "follows"
	VideosTableName    = "videos"
	MessageTableName   = "messages"
	FavoritesTableName = "likes"
	CommentTableName   = "comments"

	VideoFeedCount       = 30
	FavoriteActionType   = 1
	UnFavoriteActionType = 2

    OSSVideoBucketName = "videobucket"
    OSSImgBucketName = "imagebucket"

	DefaultSign       = "TikTok"
	DefaultAva        = "avatar/test1.jpg"
	DefaultBackground = "background/test1.png"
)
