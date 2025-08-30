package rpc

import (
	"casher-server/internal/config"

	"go.uber.org/zap"
)

type Order struct {
	Profile *config.Profile
	Lager   *zap.Logger
}
