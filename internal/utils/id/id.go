package id

import (
	"errors"
	"fmt"
	"math"

	"casher-server/internal/crypto/aes"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

func NextId(svr int64) (int64, error) {
	node, err := snowflake.NewNode(svr)
	if err != nil {
		return 0, err
	}

	// Generate a snowflake ID.
	snowId := node.Generate()
	return snowId.Int64(), nil
}

// 对称加密IP和端口，当做clientId
func EnId(raw, key []byte) string {
	str, err := aes.Base64AESCBCEncrypt(raw, key)
	if err != nil {
		panic(err)
	}
	return str
}

// 对称加密IP和端口，当做clientId
func DeId(clientId string, key []byte) []byte {
	str, err := aes.Base64AESCBCDecrypt(clientId, key)
	if err != nil {
		panic(err)
	}
	return str
}

func AESKey(key string) []byte {
	switch len(key) {
	case 16, 24, 32:
		return []byte(key)
	default:
		panic(errors.New("aes 加密密码长度需要：16，24，32"))
	}
}

// 生成uuid
func GetUuid() string {
	return uuid.NewString()
}

var unitArr = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func ByteFmt(size int64) string {
	if size == 0 {
		return "unknown"
	}
	fs := float64(size)
	p := int(math.Log(fs) / math.Log(1024))
	val := fs / math.Pow(1024, float64(p))
	_, frac := math.Modf(val)
	if frac > 0 {
		return fmt.Sprintf("%.1f%s", math.Floor(val*10)/10, unitArr[p])
	} else {
		return fmt.Sprintf("%d%s", int(val), unitArr[p])
	}
}
