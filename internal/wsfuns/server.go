package wsfuns

import (
	"casher-server/internal/config"
	"casher-server/internal/store"
	"casher-server/internal/utils"
	"context"
	"time"

	"go.uber.org/zap"
)

type WsServerApi struct {
	Profile  *config.Profile
	Lager    *zap.Logger
	Store    *store.Store
	Language string
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
