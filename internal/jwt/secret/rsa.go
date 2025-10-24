package secret

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
)

type RsGenerator struct {
}

// Generate 生成RSA密钥对
// 步骤：
// 生成RSA私钥和公钥。
// 将私钥和公钥转换为X509格式。
// 将私钥和公钥封装为PEM格式并保存到指定路径。
// 返回包含私钥和公钥文件路径的对象。
func (rs *RsGenerator) Generate() (*OutSecret, error) {
	out := &OutSecret{}
	var err error
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	// rsa.GenerateKey 用于生成指定位数rsa密钥对 rand.Reader提供加密安全随机数，1024位数
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// x509格式封装
	// x509是一种公私钥证书标准格式，定义公钥证书结构
	// 用x509封装，保证一致性
	x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	// ，privateKey.PublicKey 是一个 rsa.PublicKey 类型的值（不是指针），
	// 因此你需要使用 & 来获取它的地址，以便将其传递给 x509.MarshalPKIXPublicKey 函数。
	// 这里接收的是 &privateKey.PublicKey，表示获取 privateKey.PublicKey 的地址。
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// 表示 PEM 编码的数据块。block用于保存加密密钥、证书等二进制数据。
	privateBlock := &pem.Block{
		Type:  "PRIVATE KEY",  // 数据块类型，例如"PRIVATE KEY" 或 "PUBLIC KEY"
		Bytes: x509PrivateKey, // 用x509封装的，实际的二进制数据
	}
	// 生成对应文件

	privateKeyFile := KEY_PATH + "/rsa/private.pem"
	privateFile, err := os.OpenFile(privateKeyFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer privateFile.Close()
	pem.Encode(privateFile, privateBlock) // pem.Encode将block数据块写入到io.Writer接口，这里就是文件

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509PublicKey,
	}
	publicKeyFile := KEY_PATH + "/rsa/public.pem"
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

// 公钥加密
func RsaEncrypt(data, keyBytes []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), ciphertext)
	if err != nil {
		return nil, err
	}
	return data, nil
}
