package lager

import (
	"casher-server/internal/i18n"
	"context"
	"errors"
	"fmt"

	"golang.org/x/text/language"
)

type ITracker interface {
	// appid, trackid, lang
	GetTrackInfo(ctx context.Context) (string, string, string)
	GetOpenId(ctx context.Context) string
}

type Tracker struct {
	AppId   string `json:"app_id"`
	TrackId string `json:"track_id"`
	Lang    string `json:"lang"`
	ProxyId int64  `json:"proxy_id"`
	ComId   int64  `json:"com_id"`
	StoreId int64  `json:"store_id"`
	ShopId  int64  `json:"shop_id"`
}

// Error 错误日志
func (t *Tracker) Error(msgId, def string, fields ...i18n.Field) error {
	return errors.New(t.L(msgId, def, fields...))
}

func (t *Tracker) L(msgId, def string, fields ...i18n.Field) string {
	if t.Lang == "" {
		t.Lang = language.Chinese.String()
	}
	l := len(fields)
	if l == 0 {
		return i18n.T(t.Lang, msgId, def)
	}

	data := i18n.D{}
	for _, f := range fields {
		data[f.Key] = f.Value()
	}
	for k, v := range data {
		fmt.Printf("%s=%s\n", k, v)
	}
	return i18n.TWithData(t.Lang, msgId, def, data)
}
