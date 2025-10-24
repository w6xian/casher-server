package ed

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	ErrKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
	ErrNotRSAPublicKey     = errors.New("key is not a valid RSA public key")
)

type EDKey struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
	err        error
}

func NewEDKey() *EDKey {
	return &EDKey{}
}

func FromPrivatePemKey(priPem string) (*EDKey, error) {

	// 读取私钥文件
	privateKeyPEM, err := os.ReadFile(priPem)
	if err != nil {
		return nil, err
	}
	return FromPrivatePemKeyBytes(privateKeyPEM)
}

func FromPrivatePemKeyBytes(priKey []byte) (*EDKey, error) {
	if priKey == nil {
		return nil, errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	}
	if len(priKey) == 0 {
		return nil, errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	}
	k := &EDKey{}
	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(priKey); block == nil {
		return nil, errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	}

	// Parse the key
	var parsedKey interface{}
	var err error
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey ed25519.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(ed25519.PrivateKey); !ok {
		return nil, errors.New("invalid key: Key must be a PEM encoded ED key")
	}
	k.PrivateKey = pkey
	return k, nil
}

func FromPublicPemKey(pubPem string) (*EDKey, error) {

	// 读取私钥文件
	publicKeyPEM, err := os.ReadFile(pubPem)
	if err != nil {
		return nil, err
	}

	return FromPublicKeyBytes(publicKeyPEM)
}

func FromPublicKeyBytes(publicKeyPEM []byte) (*EDKey, error) {
	if publicKeyPEM == nil {
		return nil, errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	}
	if len(publicKeyPEM) == 0 {
		return nil, errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	}

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(publicKeyPEM); block == nil {
		return nil, errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	}

	// Parse the key
	var parsedKey interface{}
	var err error
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey ed25519.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(ed25519.PublicKey); !ok {
		return nil, errors.New("invalid key: Key must be a PEM encoded ED key")
	}
	k := &EDKey{
		PublicKey: pkey,
	}
	return k, nil
}

func DecodePublicPemKey(pubPem string) (ed25519.PublicKey, error) {
	// 读取私钥文件
	publicKeyPEM, err := os.ReadFile(pubPem)
	if err != nil {
		return nil, err
	}

	return DecodePublicPemBytes(publicKeyPEM)

}

func DecodePublicPemBytes(publicKeyPEM []byte) (ed25519.PublicKey, error) {

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(publicKeyPEM); block == nil {
		return nil, errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	}

	// Parse the key
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	var pkey ed25519.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(ed25519.PublicKey); !ok {
		return nil, errors.New("invalid key: Key must be a PEM encoded ED key")
	}
	return pkey, nil
}

func (k *EDKey) Valid() {}

func (k *EDKey) Keys() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		k.err = err
		return nil, nil, err
	}
	k.PublicKey = publicKey
	k.PrivateKey = privateKey
	k.err = err
	return publicKey, privateKey, err
}
func (k *EDKey) HasError() bool {
	return k.err != nil
}

func (k *EDKey) Save() error {
	p := "./"
	if k.HasError() {
		fmt.Println(k.err.Error())
		return k.err
	}
	// x509格式封装
	x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(k.PrivateKey) // ed25519的这里privateKey不是一个指针

	if err != nil {
		fmt.Println(err)
		return err
	}
	x509PublicKey, err := x509.MarshalPKIXPublicKey(k.PublicKey)
	if err != nil {
		log.Println(err)
		return err
	}

	// 设置pem编码的数据块
	privateBlock := &pem.Block{
		Type:  "PRIVATE_KEY",
		Bytes: x509PrivateKey,
	}
	privateKeyFile := p + "private.pem"
	privateFile, err := os.OpenFile(privateKeyFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer privateFile.Close()
	pem.Encode(privateFile, privateBlock) // pem编码

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509PublicKey,
	}
	publicKeyFile := p + "public.pem"
	publicFile, err := os.OpenFile(publicKeyFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer publicFile.Close()
	pem.Encode(publicFile, publicBlock)
	return nil
}

/**

func main() {

	k := ed.FromPemKey("./private.pem")

	// 生成JWT令牌
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	//
	//	  iss (issuer): 发行者
	//    exp (expiration time): 过期时间
	//    sub (subject): 主题
	//    aud (audience): 接收者
	//    iat (issued at): 发行时间
	//    nbf (not before): 生效时间
	//
	claims["sub"] = "1234567890"
	claims["token"] = uuid.New().String()
	claims["iat"] = 1516239022

	// 对JWT令牌进行签名
	signedToken, err := token.SignedString(k.PrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(signedToken)
	return

	publicKey, err := ed.DecodePublicPemKey("./public.pem")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	signedToken := "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MTYyMzkwMjIsInN1YiI6IjEyMzQ1Njc4OTAiLCJ0b2tlbiI6Ijg0NmI4ZTViLTA0ZTgtNDM2ZC1iOWE5LWVjOGYyNGVhMTViYyJ9.kSjeS4cGIXswcqTwArnQJNpTezXyOsW1-YcpMFri7Gted8UszBt3ugTNlBFhcdPUuoF61CKgcCos1LBoliuWAw"

	// 验证JWT令牌的有效性
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		for k, v := range claims {
			fmt.Println(k, v)
		}
	} else {
		fmt.Println(err)
	}
}
*/
