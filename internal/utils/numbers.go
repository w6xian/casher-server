package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func JoinInt64(arr []int64, s string) string {
	str := strings.Trim(fmt.Sprint(arr), "[ ]")
	return strings.Replace(str, " ", s, -1)
}

func Decimal(num float64, f int) float64 {
	fs := fmt.Sprintf("%%.%df", f)
	fmt.Println(fs)
	num, _ = strconv.ParseFloat(fmt.Sprintf(fs, num), 64)
	return num
}

func IsZero[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](tar T, def T) T {
	if tar > 0 || tar < 0 {
		return tar
	}
	return def
}
