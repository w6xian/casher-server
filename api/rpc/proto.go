package rpc

import (
	"casher-server/internal/config"
	"casher-server/internal/message"
	"casher-server/internal/store"

	"github.com/louis-xie-programmer/go-local-cache/cache"
	"go.uber.org/zap"
)

type Shop struct {
	Profile  *config.Profile
	Lager    *zap.Logger
	Store    *store.Store
	Actor    message.IMessager
	Language string
	Cache    *cache.Cache
}
type RegisterRequest struct {
	Name     string
	Password string
}

type LoginRequest struct {
	Name     string
	Password string
}

type RegisterReply struct {
	Code      int
	AuthToken string
}

type LoginReply struct {
	Code      int
	AuthToken string
}
