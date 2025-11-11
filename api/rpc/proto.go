package rpc

import (
	"casher-server/internal/config"
	"casher-server/internal/queue"
	"casher-server/internal/store"

	"go.uber.org/zap"
)

type Shop struct {
	Profile  *config.Profile
	Lager    *zap.Logger
	Store    *store.Store
	Actor    *queue.ActorPool
	Language string
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
