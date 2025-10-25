package rpc

import (
	"fmt"
	"time"
)

// 实现IEncrypt
func setSign(req IEncrypt) error {
	appId := req.EncryptInfo()
	// 校验 appId 是否为空
	if appId == "" {
		return fmt.Errorf("server setSign appId is empty")
	}
	ts := time.Now().Unix()
	// appId + ts 签名 RsaEncrypt
	code := fmt.Sprintf("%s:%d", appId, ts)
	sign, err := RsaEncrypt([]byte(code), []byte(LOGIN_PUBLIC_KEY))
	if err != nil {
		return err
	}
	// 设置 sign
	sErr := req.SetSign(string(sign), ts)
	if sErr != nil {
		return sErr
	}
	return nil
}

// 实现IDecrypt
func checkSign(req IDecrypt, appId string) error {
	sign, ts := req.DecryptInfo()
	if appId == "" {
		return fmt.Errorf("server checkSign appId is empty")
	}
	// 校验 ts 是否为空
	if ts == 0 {
		return fmt.Errorf("ts is empty")
	}
	// 校验 sign 是否为空
	if sign == "" {
		return fmt.Errorf("sign is empty")
	}
	// sign 解密 RsaDecrypt
	code, err := RsaDecrypt([]byte(sign), []byte(LOGIN_PRIVATE_KEY))
	if err != nil {
		return err
	}
	// 校验 appId + ts 是否一致
	expectedCode := fmt.Sprintf("%s:%d", appId, ts)
	if string(code) != expectedCode {
		return fmt.Errorf("invalid sign")
	}
	return nil
}
