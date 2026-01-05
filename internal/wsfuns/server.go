package wsfuns

import (
	"casher-server/internal/config"
	"casher-server/internal/store"
	"casher-server/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/w6xian/sloth"
	"github.com/w6xian/sloth/message"
	"github.com/w6xian/tlv"
	"go.uber.org/zap"
)

type WsServerApi struct {
	Profile  *config.Profile
	Lager    *zap.Logger
	Store    *store.Store
	Language string
}

type Header struct {
	TrackId string `json:"track_id"`
	AppId   string `json:"app_id"`
	Lang    string `json:"lang"`
	Ts      int64  `json:"ts"`
	Sign    string `json:"sign"`
}

func NewWsServerApi(profile *config.Profile, lager *zap.Logger, store *store.Store, language string) *WsServerApi {
	return &WsServerApi{
		Profile:  profile,
		Lager:    lager,
		Store:    store,
		Language: language,
	}
}
func (s *WsServerApi) Test(ctx context.Context, req string) (string, error) {
	return string(utils.Serialize(map[string]string{"req": "server 1", "resp": time.Now().Format("2006-01-02 15:04:05")})), nil
}

func (s *WsServerApi) Pong(ctx context.Context, req string) (struct{}, error) {
	return struct{}{}, nil
}

func (s *WsServerApi) ProductSn(ctx context.Context, sn string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	ctx, close := s.Start(ctx)
	defer close()

	header := ctx.Value(sloth.HeaderKey).(message.Header)
	if header == nil {
		return nil, fmt.Errorf("header is nil")
	}
	// 校验 sign
	err := s.checkSign(header)
	if err != nil {
		return nil, err
	}
	tracker, err := s.TrackerFromHeader(ctx, header)
	if err != nil {
		return nil, err
	}
	resp, err := s.Store.GetPublicProductBySnV2(ctx, tracker, sn)
	if err != nil {
		return nil, err
	}
	return tlv.JsonEnpack(resp)
}

func (s *WsServerApi) ReplaceProduct(ctx context.Context, req []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	ctx, close := s.Start(ctx)
	defer close()
	header := ctx.Value(sloth.HeaderKey).(message.Header)
	if header == nil {
		return nil, fmt.Errorf("header is nil")
	}
	// 校验 sign
	err := s.checkSign(header)
	if err != nil {
		return nil, err
	}

	tracker, err := s.TrackerFromHeader(ctx, header)
	if err != nil {
		return nil, err
	}
	var reqBody *store.ProductModel
	// req, err = tlv.JsonUnpack(req)
	// if err != nil {
	// 	return nil, err
	// }
	err = json.Unmarshal(req, &reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := s.Store.ReplacePublicProduct(ctx, tracker, reqBody)
	if err != nil {
		return nil, err
	}
	return tlv.JsonEnpack(resp)
}
