package wsfuns

import (
	"casher-server/proto"
	"fmt"
	"strings"

	"github.com/w6xian/sloth/message"
)

func (s *WsServerApi) checkSign(header message.Header) error {
	appId := header.Get("app_id")
	sign := header.Get("sign")
	ts := header.Get("ts")
	norm := header.Get("norm")
	apiKey := header.Get("api_key")
	// 校验 appId 是否为空
	if appId == "" {
		return fmt.Errorf("server checkSign appId is empty")
	}
	// ts
	if ts == "" {
		return fmt.Errorf("server setSign ts is empty")
	}
	// norm
	if norm == "" {
		return fmt.Errorf("server setSign norm is empty")
	}
	// sign
	if sign == "" {
		return fmt.Errorf("server setSign sign is empty")
	}

	// sign 解密 RsaDecrypt
	code, err := proto.RsaDecrypt([]byte(sign), []byte(proto.LOGIN_PRIVATE_KEY))
	if err != nil {
		return err
	}
	// 校验 appId + ts 是否一致
	prev := fmt.Sprintf("%s:%s:%s", norm, appId, apiKey)
	next := fmt.Sprintf(":%s", ts)

	if !strings.HasPrefix(string(code), prev) || !strings.HasSuffix(string(code), next) {
		return fmt.Errorf("invalid sign, expectedCode = %s:%s, code = %s", prev, next, string(code))
	}
	fmt.Println("code = ", string(code), prev, next)
	return nil
}
