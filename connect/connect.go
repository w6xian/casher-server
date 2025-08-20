/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 18:18
 */
package connect

import (
	"casher-server/internal/config"
	"casher-server/internal/utils/id"
	"context"

	"fmt"
	_ "net/http/pprof"
	"runtime"

	"go.uber.org/zap"
)

var DefaultServer *Server

type Connect struct {
	ServerId string
	Lager    *zap.Logger
	Profile  *config.Profile
}

func New(ctx context.Context, profile *config.Profile, lager *zap.Logger) *Connect {

	svr := new(Connect)
	svr.Profile = profile
	svr.Lager = lager
	return svr
}

func (c *Connect) Run() {
	// get Connect layer config
	connectConfig := config.Conf.Connect

	//set the maximum number of CPUs that can be executing
	runtime.GOMAXPROCS(connectConfig.ConnectBucket.CpuNum)

	//init Connect layer rpc server, logic client will call this
	Buckets := make([]*Bucket, connectConfig.ConnectBucket.CpuNum)
	for i := 0; i < connectConfig.ConnectBucket.CpuNum; i++ {
		Buckets[i] = NewBucket(BucketOptions{
			ChannelSize:   connectConfig.ConnectBucket.Channel,
			RoomSize:      connectConfig.ConnectBucket.Room,
			RoutineAmount: connectConfig.ConnectBucket.RoutineAmount,
			RoutineSize:   connectConfig.ConnectBucket.RoutineSize,
		})
	}
	operator := new(DefaultOperator)
	DefaultServer = NewServer(Buckets, operator, c.Profile, c.Lager)
	c.ServerId = fmt.Sprintf("%s-%s", "ws", id.ShortID())
	c.Lager.Info("Connect layer server id", zap.String("server_id", c.ServerId))

	//start Connect layer server handler persistent connection
	if err := c.InitWebsocket(); err != nil {
		c.Lager.Panic("Connect layer InitWebsocket() error", zap.Error(err))
		panic("Connect layer InitWebsocket() error")
	}
}
