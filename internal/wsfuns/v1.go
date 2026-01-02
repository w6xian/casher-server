package wsfuns

import (
	"casher-server/internal/lager"
	"context"

	"github.com/w6xian/sloth"
	"github.com/w6xian/sloth/nrpc/wsocket"
)

// :: 启动方法
func (v1 *WsServerApi) Start(ctx context.Context) (context.Context, func()) {
	ctx = lager.RequestLager(ctx, v1.Lager)
	ctx, close := v1.Store.DbConnectWithClose(ctx)
	// 准备日志资料
	return ctx, func() {
		close()
	}
}
func (v1 *WsServerApi) GetTracker(ctx context.Context, header *Header) (*lager.Tracker, error) {

	ch := ctx.Value(sloth.ChannelKey).(*wsocket.WsChannelServer)
	auth, err := ch.GetAuthInfo()
	if err != nil {
		return nil, err
	}
	return &lager.Tracker{
		AppId:   header.AppId,
		TrackId: header.TrackId,
		Lang:    header.Lang,
		ProxyId: auth.UserId,
		ComId:   auth.RoomId,
		StoreId: 0,
		ShopId:  0,
		UserId:  auth.UserId,
	}, nil
}

func (v1 *WsServerApi) AnonimousTracker(ctx context.Context, header *Header) (*lager.Tracker, error) {
	return &lager.Tracker{
		AppId:   header.AppId,
		TrackId: header.TrackId,
		Lang:    header.Lang,
		ProxyId: 0,
		ComId:   2,
		StoreId: 0,
		ShopId:  0,
	}, nil
}
