package rpc

import (
	"casher-server/internal/config"
	"casher-server/internal/lager"
	"casher-server/internal/store"
	"context"

	"go.uber.org/zap"
)

type Shop struct {
	Profile *config.Profile
	Lager   *zap.Logger
	Store   *store.Store
}

func (v *Shop) Start(ctx context.Context) (context.Context, func()) {
	ctx = lager.RequestLager(ctx, v.Lager)
	// 准备日志资料
	return ctx, func() {

	}
}
