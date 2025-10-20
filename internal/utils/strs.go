package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToCamelCase(input string) string {
	titleSpace := cases.Title(language.Dutch).String(strings.Replace(input, "_", " ", -1))
	camel := strings.Replace(titleSpace, " ", "", -1)
	return strings.ToUpper(camel[:1]) + camel[1:]
}

func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}
func Base64Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}

func MD5(input string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(input))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func IsEmptyUseDefault(value string, defaultValue string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return defaultValue
}
