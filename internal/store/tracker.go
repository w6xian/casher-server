package store

import (
	"casher-server/internal/errors"
	"casher-server/internal/i18n"
	"context"

	"golang.org/x/text/language"
)

type Tracker struct {
	ProxyId int64 `json:"proxy_id"`
	// ShopId 店铺Id
	ShopId    int64  `json:"shop_id"`
	ShopName  string `json:"shop_name"`
	ComId     int64  `json:"com_id"`
	StoreId   int64  `json:"store_id"`
	MachineId string `json:"machine_id"`
	AppId     string `json:"app_id"`
	OpenId    string `json:"open_id"`
	// 序号
	MachineNo   int64  `json:"machine_no"`
	TrackId     string `json:"track_id"`
	HandlerId   int64  `json:"handler_id"`
	HandlerName string `json:"handler_name"`
	Language    string `json:"language"`
	JwtToken    string `json:"jwt_token"`
}

func (t *Tracker) L(key string, def string, fields ...i18n.Field) string {
	if t.Language == "" {
		t.Language = language.Chinese.String()
	}
	l := len(fields)
	if l == 0 {
		return i18n.T(t.Language, key, def)
	}

	data := i18n.D{}
	for _, f := range fields {
		data[f.Key] = f.Value()
	}
	return i18n.TWithData(t.Language, key, def, data)
}

func (t *Tracker) Lang() language.Tag {
	if t.Language == "" {
		t.Language = language.Chinese.String()
	}
	return language.Make(t.Language)
}

func (t *Tracker) Error(k string, err string, fields ...i18n.Field) error {
	return errors.New(t.L("msg_error", err, fields...))
}

func NewAnonimousTracker(ctx context.Context) *Tracker {
	return &Tracker{
		MachineNo: 1,
	}
}

func NewTracker() *Tracker {
	return &Tracker{
		MachineNo: 1,
	}
}
