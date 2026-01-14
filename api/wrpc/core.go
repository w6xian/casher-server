package wrpc

import (
	"casher-server/internal/timex"
	"casher-server/internal/utils/id"
	"casher-server/proto"
	"context"
	"fmt"
	"strconv"

	"github.com/w6xian/sloth"
	"github.com/w6xian/sloth/message"
	"github.com/w6xian/sqlm"
)

type WSReq struct {
	ShopId    int64  `json:"shop_id"`
	HandlerId int64  `json:"handler_id"`
	TrackId   string `json:"track_id"`
}

func NewWSReq(shopId, handlerId int64, trackId string) *WSReq {
	return &WSReq{
		ShopId:    shopId,
		HandlerId: handlerId,
		TrackId:   trackId,
	}
}

type IDriver interface {
	GetLink(ctx context.Context) sqlm.ITable
}

type WSProxy struct {
	Server       *sloth.ClientRpc
	Client       *sloth.ServerRpc
	ServerHeader message.Header
	ClientHeader message.Header
	Driver       IDriver
}

func (p *WSProxy) CallClient(ctx context.Context, userId int64, method string, params ...any) ([]byte, error) {
	header := p.ServerHeader.Clone()
	// sign, norm, ts, err := p.Sign(header)
	// if err != nil {
	// 	return nil, err
	// }
	// header.Set("ts", ts)
	// header.Set("sign", string(sign))
	// header.Set("norm", norm)
	header.Set("track_id", id.ShortID())
	if p.ServerHeader != nil {
		return p.Server.CallWithHeader(ctx, header, userId, method, params...)
	}
	return p.Server.Call(ctx, userId, method, params...)
}

func (p *WSProxy) CallServer(ctx context.Context, method string, params ...any) ([]byte, error) {
	header := p.ClientHeader.Clone()
	sign, norm, ts, err := p.Sign(header)
	if err != nil {
		return nil, err
	}
	header.Set("ts", ts)
	header.Set("sign", string(sign))
	header.Set("norm", norm)
	header.Set("track_id", id.ShortID())
	return p.Client.CallWithHeader(ctx, header, method, params...)
}

func (p *WSProxy) CallServerWithHeader(ctx context.Context, header message.Header, method string, params ...any) ([]byte, error) {
	// 合并 ClientHeader
	for k, v := range p.ClientHeader {
		if header.Get(k) == "" {
			header.Set(k, v)
		}
	}
	sign, norm, ts, err := p.Sign(header)
	if err != nil {
		return nil, err
	}
	header.Set("ts", ts)
	header.Set("sign", string(sign))
	header.Set("norm", norm)
	header.Set("track_id", id.ShortID())
	return p.Client.CallWithHeader(ctx, header, method, params...)
}

func (p *WSProxy) Broadcast(ctx context.Context, action int, data string) error {
	p.Server.Broadcast(ctx, action, data)
	return nil
}

func (p *WSProxy) Channel(ctx context.Context, userId int64, action int, data string) error {
	p.Server.Channel(ctx, userId, action, data)
	return nil
}

func (p *WSProxy) Room(ctx context.Context, roomId int64, action int, data string) error {
	p.Server.Room(ctx, roomId, action, data)
	return nil
}

// Sign 获取请求头签名
// @return sign 签名
// @return norm 归一化字符串
// @return ts 时间戳  strconv.FormatInt(ts, 10)
// @return err 错误信息
func (p *WSProxy) Sign(header message.Header) ([]byte, string, string, error) {
	ts := timex.UnixTime()
	appId := header.Get("app_id")
	appSecret := header.Get("api_secret")
	appKey := header.Get("api_key")
	// 校验 appId 是否为空
	if appId == "" {
		return nil, "", "", fmt.Errorf("server Sign appId is empty")
	}
	// 校验 appKey 是否为空
	if appKey == "" {
		return nil, "", "", fmt.Errorf("server Sign appKey is empty")
	}
	// 校验 appSecret 是否为空
	if appSecret == "" {
		return nil, "", "", fmt.Errorf("server Sign appSecret is empty")
	}
	norm := id.ShortID()
	// appId + ts 签名 RsaEncrypt
	code := fmt.Sprintf("%s:%s:%s:%s:%d", norm, appId, appKey, appSecret, ts)
	sign, err := proto.RsaEncrypt([]byte(code), []byte(proto.LOGIN_PUBLIC_KEY))
	if err != nil {
		return nil, "", "", err
	}
	tsStr := strconv.FormatInt(ts, 10)
	return sign, norm, tsStr, nil
}

func (p *WSProxy) SetClientHeader(auth map[string]string) {
	for k, v := range auth {
		p.ClientHeader.Set(k, v)
	}
}

func (p *WSProxy) Logined(userId int64, roomId int64) {
	p.Client.RoomId = roomId
	p.Client.UserId = userId
}

func (p *WSProxy) IsLogin(ctx context.Context) bool {
	return p.Client.UserId != 0
}
