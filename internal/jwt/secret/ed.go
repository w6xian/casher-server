package secret

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

type EdGenerator struct {
}

func (ed *EdGenerator) Generate() (*OutSecret, error) {
	out := &OutSecret{}
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// x509格式封装
	x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey) // ed25519的这里privateKey不是一个指针
	if err != nil {
		log.Println(err)
		return nil, err
	}
	x509PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 设置pem编码的数据块
	privateBlock := &pem.Block{
		Type:  "PRIVATE_KEY",
		Bytes: x509PrivateKey,
	}
	privateKeyFile := KEY_PATH + "/ed/private.pem"
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
	publicKeyFile := KEY_PATH + "/ed/public.pem"
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
