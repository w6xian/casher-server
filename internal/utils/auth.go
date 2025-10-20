package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Orz3Decode(token string) string {
	if strings.HasPrefix(token, "orz:L/") || strings.HasPrefix(token, "L/") {
		s := strings.Split(token, "/")
		if len(s) == 3 {
			session := strings.TrimRight(s[1], "G")
			// ip, _ := strconv.ParseInt(strings.Split(s[2], ";")[0], 10, 64)
			// ipv4 := util.InetNtoA(ip)
			return session
		}
	}
	return token
}

func Orz3Encode(token string, ip string) string {
	return fmt.Sprintf("L/%s;no-cache", parseOrz3(token, ip))
}

func parseOrz3(token string, ip string) string {
	intIp := IPv4ToInt(ip)
	// !!! 取后8位，这里用服务器加密，其他软件要注意这里
	l8 := token[8:]
	k := GetInt(string(l8)) & 0xFFFF
	return fmt.Sprintf("%sG/%d", token, intIp^uint32(k))
}

func IsOrz(token string) bool {
	if strings.HasPrefix(token, "orz:L/") || strings.HasPrefix(token, "L/") {
		return true
	}
	return false
}

func GenerateSignature(apiKey, apiSecret string, timestamp int64, code, token string) string {
	message := fmt.Sprintf("%s%d%s%s", apiKey, timestamp, code, token)
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(message))
	signature := hex.EncodeToString(mac.Sum(nil))
	return signature
}

// CalcSign 计算签名：传入结构体，按Tag字母升序排序，值转换为字符串后，按Tag:Value格式拼接，多个属性之间用分号分隔，算出md5再与key再md5得到sign
func CalcSign(obj interface{}, key string) string {
	// 使用反射获取结构体信息
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// 如果是指针，获取其元素
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	// 确保是结构体类型
	if t.Kind() != reflect.Struct {
		return ""
	}
	// 存储tag和值的映射
	tagValues := make(map[string]string)
	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// 获取json tag
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		// 处理tag中的选项（如omitempty）
		tagName := strings.Split(tag, ",")[0]
		if tagName == "sign" {
			continue // 跳过sign字段本身
		}
		// 获取字段值
		fieldValue := v.Field(i)
		// 检查字段是否可导出且有效
		if fieldValue.CanInterface() {
			// 将值转换为字符串
			valueStr := fmt.Sprintf("%v", fieldValue.Interface())
			tagValues[tagName] = valueStr
		}
	}
	// 获取所有tag并排序
	var tags []string
	for tag := range tagValues {
		tags = append(tags, tag)
	}
	sort.Strings(tags)
	// 按顺序拼接tag:value
	var parts []string
	for _, tag := range tags {
		parts = append(parts, fmt.Sprintf("%s:%s", tag, tagValues[tag]))
	}
	// 用分号连接所有部分
	strToSign := strings.Join(parts, ";")
	// 计算第一次MD5
	h1 := md5.Sum([]byte(strToSign))
	md5Result := hex.EncodeToString(h1[:])
	// 与key拼接后计算第二次MD5
	h2 := md5.Sum([]byte(md5Result + key))
	finalSign := hex.EncodeToString(h2[:])
	return finalSign
}
