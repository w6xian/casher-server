/**
 * Created by lock
 * Date: 2019-10-06
 * Time: 22:46
 */
package rpc

import (
	"casher-server/internal/config"
	"casher-server/internal/store"
	"context"
	"sync"

	"github.com/smallnest/rpcx/server"
	"go.uber.org/zap"
)

var once sync.Once

func InitLogicRpcServer(ctx context.Context, profile *config.Profile, logger *zap.Logger, store *store.Store) {
	once.Do(func() {
		order := new(Order)
		order.Profile = profile
		order.Lager = logger
		order.Store = store

		s := server.NewServer()
		s.DisableHTTPGateway = true
		s.RegisterName("micro-order", order, "")
		s.Serve("tcp", profile.Server.RpcAddr)
	})
}
