package store

import (
	"casher-server/internal/jwt/secret"
	"casher-server/internal/timex"
	"encoding/base64"
	"fmt"
	"strings"
)

type IEncrypt interface {
	SetSign(sign string, ts int64) error
}
type IDecrypt interface {
	DecryptInfo() (sign string, ts int64)
}

const LOGIN_PUBLIC_KEY = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC6B8yicpX5alZUQTuGimkRy2R3
rwwirRC3OJkL3Z+uzlBxiJm0EtTNd7QMk15xBKwfmDvvmeZ/vKf58v+6LJXR40W+
0/PoaW613XVeHGx8seq53QLi65OPkwfnlVTGK1mrjMf+GqMIjsNaMtWSP4nOtOkD
Q+VMScfbSQOt1tpFHwIDAQAB
-----END PUBLIC KEY-----
`

const LOGIN_PRIVATE_KEY = `
-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALoHzKJylflqVlRB
O4aKaRHLZHevDCKtELc4mQvdn67OUHGImbQS1M13tAyTXnEErB+YO++Z5n+8p/ny
/7osldHjRb7T8+hpbrXddV4cbHyx6rndAuLrk4+TB+eVVMYrWauMx/4aowiOw1oy
1ZI/ic606QND5UxJx9tJA63W2kUfAgMBAAECgYAsEY7raKOYexVKTk2wmHc9bOY2
5/PC/yZ3kOPIiM68FUm/K3Hl05QvvEydsgdsVIQF1AVWxClzVxifwG3OB6PRtkPc
n/NaRpI7rDgpG2k9dsQ8uB8HqpSYH0CxFtNKoNiZIfGuEOOQls9j7jAknLIEN+L5
Eu+XW4c+UhLu0tNbKQJBAOJO8BYVCGxtv2j4lZrHTtPHDZvhyhLItmWOpwyjShwb
c5r7bvTQLCuI6kjB+stk5Un7FJykAn2BmPnespqGuPMCQQDScApavGRAiVf3d29r
qLGJWuegl2FGefjbRsFoWd5E2MZ4mqUrGz33a4fhb09IG1Ux45PUm1Uws8j1G4Bo
us4lAkAWso4X2OIwZc11zMDMdkLssKEnyjyHJ8RLaURN2y66pPIyUBdvzFUxxJii
1Xm+3o60nc8SasypI89g+Dn3j52LAkB4oG7fCkSxVclN+nGtFdsG8Ev8GypQmtRS
5aEyLumhL129fnAVYJ1JuaL/T63zmG9ilKCF67COpVAZaHVYE1TdAkEA0Vw6hK02
QrD9xVnnE5GkDFMyXi3BYDefwfVTS4IoOarG+IxMVjZn/uzgB6gldU9/AJEpMd/G
GSPDewU0n3+M3A==
-----END PRIVATE KEY-----
`

// 通过pubkey加密
func RsaEncrypt(plaintext []byte, publicKey []byte) ([]byte, error) {
	encrypted, err := secret.RsaEncrypt(plaintext, publicKey)
	if err != nil {
		return nil, err
	}
	return []byte(base64.StdEncoding.EncodeToString(encrypted)), nil
}

// 通过privatekey解密
func RsaDecrypt(encrypted []byte, privateKey []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(string(encrypted))
	if err != nil {
		return nil, err
	}
	decrypted, err := secret.RsaDecrypt(decodeBytes, privateKey)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

type APIReq struct {
	// 没有值，就不输出
	AppId string `json:"app_id,omitempty"`
	Sign  string `json:"sign,omitempty"`
	Ts    int64  `json:"ts,omitempty"`
}

// 实现 IDecrypt
func (req *APIReq) DecryptInfo() (string, int64) {
	return req.Sign, req.Ts
}

// 实现 IEncrypt
func (reply *APIReq) SetSign(sign string, ts int64) error {
	reply.Sign = sign
	reply.Ts = ts
	return nil
}

// 实现IEncrypt
func SetSign(req IEncrypt, cs ...string) error {
	// 校验 appId 是否为空
	if len(cs) == 0 {
		return fmt.Errorf("server setSign appId is empty")
	}
	ts := timex.UnixTime()
	// appId + ts 签名 RsaEncrypt
	code := fmt.Sprintf("%s:%d", strings.Join(cs, ""), ts)
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
func CheckSign(req IDecrypt, cs ...string) error {
	sign, ts := req.DecryptInfo()

	// 校验 ts 是否为空
	if ts == 0 {
		return fmt.Errorf("ts is empty")
	}
	// ts 校验是否过期（30秒）
	if ts < timex.Now().Unix()-30 {
		return fmt.Errorf("sign expired")
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
	expectedCode := fmt.Sprintf("%s:%d", strings.Join(cs, ""), ts)
	if string(code) != expectedCode {
		return fmt.Errorf("invalid sign")
	}
	return nil
}

// 实现IDecrypt
func CheckHeaderSign(sign, norm, appId, appKey, appSec string, ts int64) error {
	// 校验 norm 是否为空
	if norm == "" {
		return fmt.Errorf("norm is empty")
	}

	// 校验 ts 是否为空
	if ts == 0 {
		return fmt.Errorf("ts is empty")
	}
	// ts 校验是否过期（30秒）
	if ts < timex.Now().Unix()-30 {
		return fmt.Errorf("sign expired")
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
	expectedCode := fmt.Sprintf("%s:%s:%s:%s:%d", norm, appId, appKey, appSec, ts)

	if string(code) != expectedCode {
		return fmt.Errorf("invalid sign, expectedCode = %s, code = %s", expectedCode, string(code))
	}
	return nil
}

// 实现IEncrypt
// GetHeaderSign 获取请求头签名
// @return sign 签名
// @return norm 归一化字符串
// @return ts 时间戳
// @return err 错误信息
func GetHeaderSign(norm, appId, appKey, appSec string) ([]byte, string, int64, error) {
	// 校验 appId 是否为空
	if appId == "" {
		return nil, "", 0, fmt.Errorf("appId is empty")
	}
	// 校验 appKey 是否为空
	if appKey == "" {
		return nil, "", 0, fmt.Errorf("appKey is empty")
	}
	// 校验 appSec 是否为空
	if appSec == "" {
		return nil, "", 0, fmt.Errorf("appSec is empty")
	}
	// 校验 norm 是否为空
	if norm == "" {
		return nil, "", 0, fmt.Errorf("norm is empty")
	}
	ts := timex.UnixTime()
	// appId + ts 签名 RsaEncrypt
	code := fmt.Sprintf("%s:%s:%s:%s:%d", norm, appId, appKey, appSec, ts)
	sign, err := RsaEncrypt([]byte(code), []byte(LOGIN_PUBLIC_KEY))
	if err != nil {
		return nil, "", 0, err
	}
	return sign, norm, ts, nil
}
