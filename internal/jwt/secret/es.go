package secret

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

/*
为什么需要分类
在 ES 算法中，不同的签名方法（如 ES256、ES384、ES512）
对应不同的椭圆曲线和哈希函数。具体来说：
ES256 使用 P-256 曲线和 SHA-256 哈希函数。
ES384 使用 P-384 曲线和 SHA-384 哈希函数。
ES512 使用 P-521 曲线和 SHA-512 哈希函数。
这些不同的曲线和哈希函数提供了不同级别的安全性。
通过定义常量和类型 ESSigningMethodS，可以方便地管理和切换不同的签名方法，确保代码的灵活性和可扩展性
*/
const (
	ES256 ESSigningMethodS = "ES256"
	ES384 ESSigningMethodS = "ES384"
	ES512 ESSigningMethodS = "ES512"
)

type ESSigningMethodS string

type EsGenerator struct {
	SigningMethod ESSigningMethodS
}

func (es *EsGenerator) getCurve() elliptic.Curve {
	switch es.SigningMethod {
	case ES256:
		return elliptic.P256()
	case ES384:
		return elliptic.P384()
	case ES512:
		return elliptic.P521()
	default:
		return elliptic.P256()
	}
}

func (es *EsGenerator) Generate() (*OutSecret, error) {
	out := &OutSecret{}
	privateKey, err := ecdsa.GenerateKey(es.getCurve(), rand.Reader)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// x509格式封装
	x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey) // privateKey是一个指针
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// ，privateKey.PublicKey 是一个 rsa.PublicKey 类型的值（不是指针），
	// 因此你需要使用 & 来获取它的地址，以便将其传递给 x509.MarshalPKIXPublicKey 函数。
	// 这里接收的是 &privateKey.PublicKey，表示获取 privateKey.PublicKey 的地址。
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 设置pem编码的数据块
	privateBlock := &pem.Block{
		Type:  "PRIVATE_KEY",
		Bytes: x509PrivateKey,
	}
	privateKeyFile := KEY_PATH + "/es/private.pem"
	privateFile, err := os.OpenFile(privateKeyFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer privateFile.Close()
	pem.Encode(privateFile, privateBlock) // pem编码

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509PublicKey,
	}
	publicKeyFile := KEY_PATH + "/es/public.pem"
	publicFile, err := os.OpenFile(publicKeyFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer publicFile.Close()
	pem.Encode(publicFile, publicBlock)

	out.PrivateKeyFile = privateKeyFile
	out.PublicKeyFile = publicKeyFile

	return out, nil
}
