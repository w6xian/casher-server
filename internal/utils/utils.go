package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

func RandBytes(size int) []byte {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = byte(rand.Int31n(128))
	}
	return buf
}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func Copy(dst, src interface{}) {
	aj, _ := json.Marshal(src)
	_ = json.Unmarshal(aj, dst)
}

func GetInt64(val interface{}) int64 {
	switch value := val.(type) {
	case string:
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0
		}
		return val
	case int64:
		return value
	case int:
		return int64(value)
	case float64:
		return int64(value)
	default:
		fmt.Println("GetInt64=", value)
	}
	return 0
}

func GetInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return val
}

func ParseInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func Uint64(value string) uint64 {
	if d, err := strconv.ParseUint(value, 10, 64); err == nil {
		return d
	}
	return 0

}

func GetFloat64(value string) float64 {
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return val
}

func GetString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch ft := value.(type) {
	case float64:
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		key = strconv.Itoa(ft)
	case uint:
		key = strconv.Itoa(int(ft))
	case int8:
		key = strconv.Itoa(int(ft))
	case uint8:
		key = strconv.Itoa(int(ft))
	case int16:
		key = strconv.Itoa(int(ft))
	case uint16:
		key = strconv.Itoa(int(ft))
	case int32:
		key = strconv.Itoa(int(ft))
	case uint32:
		key = strconv.Itoa(int(ft))
	case int64:
		key = strconv.FormatInt(ft, 10)
	case uint64:
		key = strconv.FormatUint(ft, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func Find[T comparable](tar T, tars []T) int {
	for i, v := range tars {
		if tar == v {
			return i
		}
	}
	return -1
}

func IsEmptyString(tar string, def string) string {
	if len(tar) <= 0 {
		return def
	}
	return def
}

// GetIntranetIp 获取本机内网IP
func GetIntranetIp() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}

	return ""
}
