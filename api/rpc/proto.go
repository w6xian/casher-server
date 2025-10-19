package rpc

import (
	"casher-server/internal/config"
	"casher-server/internal/store"

	"go.uber.org/zap"
)

type Order struct {
	Profile *config.Profile
	Lager   *zap.Logger
	Store   *store.Store
}
