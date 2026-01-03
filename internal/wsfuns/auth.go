package wsfuns

import (
	"context"
	"fmt"
	"time"

	"github.com/w6xian/sloth"
	"github.com/w6xian/sloth/bucket"
	"github.com/w6xian/sloth/nrpc"
	"github.com/w6xian/tlv"
)

func (s *WsServerApi) Login(ctx context.Context, header *Header, mchId, apiKey string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	ctx, close := s.Start(ctx)
	defer close()

	authInfo, err := s.Store.GetAuthInfoUseMA(ctx, header.AppId, apiKey, mchId, header.Sign, header.Ts)
	if err != nil {
		return nil, err
	}

	ch, ok := ctx.Value(sloth.ChannelKey).(bucket.IChannel)
	if !ok {
		return nil, fmt.Errorf("channel not found")
	}
	svr, ok := ctx.Value(sloth.BucketKey).(nrpc.IBucket)
	if !ok {
		return nil, fmt.Errorf("bucket not found")
	}
	//根据data登录 解析出userId,roomId,token
	auth := nrpc.AuthInfo{
		UserId: authInfo.UserId,
		RoomId: authInfo.ProxyId,
		Token:  authInfo.ApiKey,
	}
	lerr := svr.Bucket(auth.UserId).Put(auth.UserId, auth.RoomId, auth.Token, ch)

	if lerr != nil {
		return nil, lerr
	}
	return tlv.JsonEnpack(auth)
}
