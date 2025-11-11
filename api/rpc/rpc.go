/**
 * Created by lock
 * Date: 2019-10-06
 * Time: 22:46
 */
package rpc

import (
	"casher-server/internal/config"
	"casher-server/internal/queue"
	"casher-server/internal/store"
	"context"
	"sync"

	"github.com/smallnest/rpcx/server"
	"go.uber.org/zap"
)

var once sync.Once

func InitLogicRpcServer(ctx context.Context, profile *config.Profile, logger *zap.Logger, store *store.Store, actor *queue.ActorPool) {
	once.Do(func() {
		shop := new(Shop)
		shop.Profile = profile
		shop.Lager = logger
		shop.Store = store
		shop.Actor = actor
		shop.Language = profile.Apps.Language

		s := server.NewServer()
		s.DisableHTTPGateway = true
		s.RegisterName("micro-shop", shop, "")
		s.Serve("tcp", profile.Server.RpcAddr)
	})
}
