package wsfuns

import (
	"casher-server/proto"
	"fmt"

	"github.com/w6xian/sloth/message"
)

func (s *WsServerApi) checkSign(header message.Header) error {

	// 校验 appId 是否为空
	if header.Get("app_id") == "" {
		return fmt.Errorf("server setSign appId is empty")
	}
	// ts
	if header.Get("ts") == "" {
		return fmt.Errorf("server setSign ts is empty")
	}
	// norm
	if header.Get("norm") == "" {
		return fmt.Errorf("server setSign norm is empty")
	}
	// sign
	if header.Get("sign") == "" {
		return fmt.Errorf("server setSign sign is empty")
	}

	// sign 解密 RsaDecrypt
	code, err := proto.RsaDecrypt([]byte(header.Get("sign")), []byte(proto.LOGIN_PRIVATE_KEY))
	if err != nil {
		return err
	}
	// 校验 appId + ts 是否一致
	expectedCode := fmt.Sprintf("%s:%s:%s", header.Get("norm"), header.Get("app_id"), header.Get("ts"))
	if string(code) != expectedCode {
		return fmt.Errorf("invalid sign")
	}
	return nil
}
