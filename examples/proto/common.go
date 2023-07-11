package proto

const (
	ServerUrl = "ws://127.0.0.1:8000/rpc"
)

type UserLoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type UserLoginResp struct {
	UserId      int32  `json:"user_id"`
	UserName    string `json:"user_name"`
	Age         int32  `json:"age"`
	HomeAddress string `json:"home_address"`
	Token       string `json:"token"`
}

type UserLogoutReq struct {
	UserId int32  `json:"user_id"`
	Token  string `json:"token"`
}

type UserLogoutResp struct {
	OK bool `json:"ok"`
}

type SendNotifyReq struct {
	UserId  int32  `json:"user_id"`
	Message string `json:"message"`
}

type UserEvent int

const (
	UserEvent_Login  UserEvent = 1
	UserEvent_Logout UserEvent = 2
)

type UserNotify struct {
	UserId    int32     `json:"user_id"`
	UserEvent UserEvent `json:"user_event"`
}
