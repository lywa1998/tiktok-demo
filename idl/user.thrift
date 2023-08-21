namespace go basic.user

include 'common.thrift'

struct UserReq {
    1: i64 user_id
    2: string token
}

struct UserResp {
    1: i32 status_code
    2: string status_msg
    3: common.User user
}

struct UserRegisterReq {
    1: string name (api.query="name")
    2: string password (api.query="password")
}

struct UserRegisterResp {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct UserLoginReq {
    1: string name (api.query="name")
    2: string password (api.query="password")
}

struct UserLoginResp {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

service UserHandler {
    UserResp Uer(1: UserReq request) (api.get="/douyin/user/")

    // When a new user registers, just provide a user name, password, and nickname, and the user name needs to be unique.
    // After successful creation, return the user id and permission token.
    UserRegisterResp UserRegister(1: UserRegisterReq request) (api.post="/douyin/user/register/")

    // Log in with username and password, and return user id and permission token after successful login.
    UserLoginResp UserLogin(1: UserLoginReq request) (api.post="/douyin/user/login/")
}
