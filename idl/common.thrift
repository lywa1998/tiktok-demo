namespace go common

struct User {
    1: i64 id               // user id
    2: string name          // user name
    3: i64 follow_count     // total number of people the user follows
    4: i64 follower_count   // total number of fans
    5: bool is_follow       // whether the currently logged-in user follows this user
    6: string avatar        // user avatar URL
    8: string background_image  // umage at the top of the user's personal page
    9: string signature     // user profile
    10: i64 total_favorited // number of liks for videos published by user
    11: i64 work_count      // number of videos published by user
    12: i64 favorite_count  // number of links by this user
}

struct Video {
    1: i64 id               // video id
    2: User author          // author information
    3: string play_url      // video playback URL
    4: string cover_url     // video cover URL
    5: i64 favorite_count   // total number of likes for the video
    6: i64 comment_count    // total number of comments on the video
    7: bool is_favorite     // true-Liked, false-did not like
    8: string title         // video title
}
