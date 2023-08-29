namespace go basic.user

include 'common.thrift'

struct UserRequest {
    1: i64 user_id
    2: string token
}

struct UserResponse {
    1: i32 status_code
    2: string status_msg
    3: common.User user
}

struct UserRegisterRequest {
    1: string name (api.query="name")
    2: string password (api.query="password")
}

struct UserRegisterResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct UserLoginRequest {
    1: string name (api.query="name")
    2: string password (api.query="password")
}

struct UserLoginResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

service User{
    UserResponse Uer(1: UserRequest request) (api.get="/douyin/user/")

    // When a new user registers, just provide a user name, password, and nickname, and the user name needs to be unique.
    // After successful creation, return the user id and permission token.
    UserRegisterResponse UserRegister(1: UserRegisterRequest request) (api.post="/douyin/user/register/")

    // Log in with username and password, and return user id and permission token after successful login.
    UserLoginResponse UserLogin(1: UserLoginRequest request) (api.post="/douyin/user/login/")
}
