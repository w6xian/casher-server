/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 18:18
 */
package connect

import (
	"casher-server/internal/config"
	"casher-server/internal/utils/id"
	"casher-server/proto"
	"context"
	"sync"
	"time"

	"fmt"
	_ "net/http/pprof"
	"runtime"

	"go.uber.org/zap"

	"modernc.org/mathutil"
)

var once sync.Once
var WsLogicObjc *WsLogic

type WsLogic struct {
	Server *Server
}

func InitWsLogicServer() *WsLogic {
	once.Do(func() {
		WsLogicObjc = &WsLogic{}
	})
	return WsLogicObjc
}

func (c *WsLogic) Channel(ctx context.Context, userId int64, action int, data string) {
	if c.Server == nil {
		return
	}
	b := c.Server.Bucket(userId)
	ch := b.Channel(userId)
	if ch == nil {
		return
	}
	cmd := proto.CmdReq{
		Id:     id.ShortID(),
		Ts:     time.Now().Unix(),
		Action: action,
		Data:   data,
	}
	msg := &proto.Msg{
		Body: cmd.Bytes(),
	}
	if err := ch.Push(ctx, msg); err != nil {
		fmt.Println("Connect layer Push() error", zap.Error(err))
	}
}

func (c *WsLogic) Room(ctx context.Context, roomId int64, action int, data string) {
	if c.Server == nil {
		return
	}
	room := c.Server.Room(roomId)
	if room == nil {
		return
	}
	if room.drop {
		return
	}
	fmt.Println("Room layer Push() roomId:", roomId)

	cmd := proto.CmdReq{
		Id:     id.ShortID(),
		Ts:     time.Now().Unix(),
		Action: action,
		Data:   data,
	}
	msg := &proto.Msg{
		Body: cmd.Bytes(),
	}
	room.Push(ctx, msg)
}

func (c *WsLogic) Broadcast(ctx context.Context, action int, data string) {
	if c.Server == nil {
		return
	}

	cmd := proto.CmdReq{
		Id:     id.ShortID(),
		Ts:     time.Now().Unix(),
		Action: action,
		Data:   data,
	}
	msg := &proto.Msg{
		Body: cmd.Bytes(),
	}
	if err := c.Server.Broadcast(ctx, msg); err != nil {
		return
	}
}

type Connect struct {
	ServerId string
	Lager    *zap.Logger
	Profile  *config.Profile
}

// server-bucket-channel
// each bucket has a channel to send message to client

func New(ctx context.Context, profile *config.Profile, lager *zap.Logger) *Connect {
	svr := new(Connect)
	svr.Profile = profile
	svr.Lager = lager
	return svr
}

func (c *Connect) Server(wsLogic *WsLogic) {
	// get Connect layer config
	connectConfig := c.Profile.Connect
	//set the maximum number of CPUs that can be executing
	runtime.GOMAXPROCS(connectConfig.ConnectBucket.CpuNum)

	bsNum := mathutil.Max(connectConfig.ConnectBucket.CpuNum, 1)
	bsNum = mathutil.Min(bsNum, runtime.NumCPU())
	//init Connect layer rpc server, logic client will call this
	bs := make([]*Bucket, bsNum)
	for i := 0; i < bsNum; i++ {
		bs[i] = NewBucket(BucketOptions{
			ChannelSize:   connectConfig.ConnectBucket.Channel,
			RoomSize:      connectConfig.ConnectBucket.Room,
			RoutineAmount: connectConfig.ConnectBucket.RoutineAmount,
			RoutineSize:   connectConfig.ConnectBucket.RoutineSize,
		})
	}
	operator := new(DefaultOperator)
	wsLogic.Server = NewServer(bs, operator, c.Profile, c.Lager)
	c.ServerId = fmt.Sprintf("%s-%s", "ws", id.ShortID())
	c.Lager.Info("Connect layer server id", zap.String("server_id", c.ServerId))
	//start Connect layer server handler persistent connection
	if err := c.InitWebsocket(wsLogic); err != nil {
		c.Lager.Panic("Connect layer InitWebsocket() error", zap.Error(err))
		panic("Connect layer InitWebsocket() error")
	}
}
