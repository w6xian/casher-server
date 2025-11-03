/**
 * Created by lock
 * Date: 2019-08-10
 * Time: 18:38
 */
package proto

type LoginRequest struct {
	Name     string
	Password string
}

type LoginResponse struct {
	Code      int
	AuthToken string
}

type GetUserInfoRequest struct {
	UserId int
}

type GetUserInfoResponse struct {
	Code     int
	UserId   int
	UserName string
}

type RegisterRequest struct {
	Name     string
	Password string
}

type RegisterReply struct {
	Code      int
	AuthToken string
}

type LogoutRequest struct {
	AuthToken string
}

type LogoutResponse struct {
	Code int
}

type CheckAuthRequest struct {
	AuthToken string
}

type CheckAuthResponse struct {
	Code     int
	UserId   int64
	UserName string
}

type ConnectRequest struct {
	AuthToken string `json:"auth_token"`
	RoomId    int64  `json:"room_id"`
	ServerId  string `json:"server_id"`
}

type ConnectReply struct {
	UserId int64
}

type DisConnectRequest struct {
	RoomId int64 `json:"room_id"`
	UserId int64 `json:"user_id"`
}

type DisConnectReply struct {
	Has bool
}

type Send struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	FromUserId   int64  `json:"from_user_id"`
	FromUserName string `json:"from_user_name"`
	ToUserId     int64  `json:"to_user_id"`
	ToUserName   string `json:"to_user_name"`
	RoomId       int64  `json:"room_id"`
	Op           int    `json:"op"`
	CreateTime   string `json:"create_time"`
}

type SendTcp struct {
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
	FromUserId   int64  `json:"from_user_id"`
	FromUserName string `json:"from_user_name"`
	ToUserId     int64  `json:"to_user_id"`
	ToUserName   string `json:"to_user_name"`
	RoomId       int64  `json:"room_id"`
	Op           int    `json:"op"`
	CreateTime   string `json:"create_time"`
	AuthToken    string `json:"auth_token"` //仅tcp时使用，发送msg时带上
}
