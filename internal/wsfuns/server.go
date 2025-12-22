package wsfuns

import (
	"casher-server/internal/config"
	"casher-server/internal/store"
	"casher-server/internal/utils"
	"context"

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
	return string(utils.Serialize(map[string]string{"req": "server 1"})), nil
}

func (s *WsServerApi) Pong(ctx context.Context, req string) (struct{}, error) {
	return struct{}{}, nil
}
