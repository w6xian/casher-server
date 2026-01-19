package wsfuns

import (
	"casher-server/internal/store"
	"casher-server/internal/utils"
	"casher-server/internal/utils/id"
	"context"
	"fmt"
	"time"

	"github.com/w6xian/sloth"
	"github.com/w6xian/sloth/bucket"
	"github.com/w6xian/sloth/message"
	"github.com/w6xian/sloth/nrpc"
	"github.com/w6xian/tlv"
)

func (s *WsServerApi) Login(ctx context.Context, req string) ([]byte, error) {
	header := ctx.Value(sloth.HeaderKey).(message.Header)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	ctx, close := s.Start(ctx)
	defer close()
	appId := header.Get("app_id")
	sign := header.Get("sign")
	ts := header.Get("ts")
	norm := header.Get("norm")
	apiKey := header.Get("api_key")
	mchId := header.Get("mch_id")
	authInfo, err := s.Store.GetAuthInfoUseMA(ctx, appId, apiKey, mchId, sign, norm, utils.GetInt64(ts))
	if err != nil {
		authInfo = &store.AuthInfo{
			UserId:  0,
			ProxyId: 0,
			ShopId:  0,
		}
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
		RoomId: authInfo.ShopId,
		Token:  id.ShortID(),
	}
	if auth.UserId != ch.UserId() {
		svr.Bucket(ch.UserId()).DeleteChannel(ch)
	}
	// 确保bucket存在.userId对应，才能在调用时找到UserId
	lerr := svr.Bucket(auth.UserId).Put(auth.UserId, auth.RoomId, auth.Token, ch)
	if lerr != nil {
		return nil, lerr
	}
	return tlv.JsonEnpack(auth)
}
