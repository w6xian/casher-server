package secret

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// HsGenerator HS HS256、HS384、HS512 属于对称加密算法。
// 加密和验证使用相同的密钥。
type HsGenerator struct {
	Length int
}

// Generate 生成一个指定长度随机密钥，将其编码为16进制字符串返回
// 生成过程，直接生成密钥
func (hs *HsGenerator) Generate() (*OutSecret, error) {
	out := &OutSecret{}
	length := 32
	if hs.Length > 0 {
		length = hs.Length
	}
	//rand.Seed(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	rand.Read(b)
	out.Secret = hex.EncodeToString(b)[:length]
	return out, nil
}
