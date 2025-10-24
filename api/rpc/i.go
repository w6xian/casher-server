package rpc

import (
	"casher-server/internal/jwt/secret"
	"encoding/base64"
)

type IEncrypt interface {
	EncryptInfo() (appId string)
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
