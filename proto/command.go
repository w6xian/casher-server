package proto

import (
	"casher-server/internal/utils"
	"fmt"
)

type CmdReq struct {
	Id        string `json:"id"`
	TrackId   string `json:"track_id"`
	AppId     string `json:"app_id,omitempty"`
	ProxyId   int64  `json:"proxy_id,omitempty"`
	RoomId    int64  `json:"room_id,omitempty"`
	UserId    int64  `json:"user_id,omitempty"`
	Action    int    `json:"action"` // 操作类型
	AuthToken string `json:"auth_token"`
	Data      string `json:"data"` // 操作数据
	Lang      string `json:"lang"`
	Method    string `json:"method"` // service method name
	Ts        int64  `json:"ts"`
}

func (c *CmdReq) Bytes() []byte {
	return []byte(utils.JsonString(c))
}

func (c *CmdReq) Signal() error {
	// ts := time.Now().Unix()
	if c.AppId == "" {
		return fmt.Errorf("setSign appId is empty")
	}
	// appId + ts 签名 RsaEncrypt
	// code := fmt.Sprintf("%s:%d", c.AppId, ts)
	// sign, err := RsaEncrypt([]byte(code), []byte(LOGIN_PUBLIC_KEY))
	// if err != nil {
	// 	return err
	// }
	// c.Ts = ts
	// c.Sign = string(sign)
	return nil
}

func (c *CmdReq) CheckSignal() error {
	if c.AppId == "" {
		return fmt.Errorf("checkSign appId is empty")
	}
	// 校验 ts 是否为空
	if c.Ts == 0 {
		return fmt.Errorf("ts is empty")
	}
	// // 校验 sign 是否为空
	// if c.Sign == "" {
	// 	return fmt.Errorf("sign is empty")
	// }
	// // sign 解密 RsaDecrypt
	// code, err := RsaDecrypt([]byte(c.Sign), []byte(LOGIN_PRIVATE_KEY))
	// if err != nil {
	// 	return err
	// }
	// 校验 appId + ts 是否一致
	// expectedCode := fmt.Sprintf("%s:%d", c.AppId, c.Ts)
	// if string(code) != expectedCode {
	// 	return fmt.Errorf("invalid sign")
	// }
	return nil
}
