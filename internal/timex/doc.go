package timex

import (
	"time"
)

// 时间相关函数

var local *time.Location

func InitLocation(location string) {
	if local == nil {
		local, _ = time.LoadLocation(location)
	}
	time.Local = local
}

func getLocal() *time.Location {
	if local == nil {
		panic("timex: local is nil,pls init location(timex.InitLocation())")
	}
	return local
}

func UnixTime() int64 {
	t := time.Now().In(getLocal())
	return t.Unix()
}

func Now() time.Time {
	return time.Now().In(getLocal())
}
