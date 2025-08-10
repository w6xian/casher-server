package is

import (
	"os"
	"reflect"
	"regexp"
	"strings"

	"casher-server/internal/utils/array"
)

func Nil(x any) bool {
	if x == nil {
		return true
	}

	return reflect.ValueOf(x).IsNil()
}

func ServiceArgs() bool {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		return array.InArray(arg, []string{"install", "uninstall"})
	}
	return false
}

func ExistFile(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func StartUseService(info string) bool {
	info = strings.ToLower(info)
	return info == "services.exe"
}

func StartUseExec(info string) bool {
	info = strings.ToLower(info)
	return info == "explorer.exe"
}

func StartUseCommond(info string) bool {
	info = strings.ToLower(info)
	return info == "cmd.exe"
}
func StartUseOthers(info string) bool {
	info = strings.ToLower(info)
	return array.InArray(info, []string{"services.exe", "explorer.exe", "cmd.exe"})
}

// IsEmptyString 检查字符串是否为空
func EmptyString(s string) bool {
	return s == ""
}

// IsMobile 检查字符串是否为手机号
func Mobile(m string) bool {
	if len(m) != 11 {
		return false
	}
	// 正则表达式匹配中国手机号：
	// ^1[3-9]\d{9}$
	// 解释：
	// ^1        : 以1开头
	// [3-9]     : 第二位为3-9（排除1和2，对应最新号段规则）
	// \d{9}$    : 后面跟9位数字
	mobileRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return mobileRegex.MatchString(m)
}
