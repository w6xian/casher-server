package v1

import (
	"casher-server/connect"
	"casher-server/internal/config"
	"casher-server/internal/muxhttp"
	"casher-server/internal/queue"
	"context"
	"net/http"
	"time"

	"github.com/louis-xie-programmer/go-local-cache/cache"
	"go.uber.org/zap"
)

type Api struct {
	Context context.Context
	Profile *config.Profile
	Lager   *zap.Logger
	Cache   *cache.Cache
	Actor   *queue.ActorPool
	WsLogic *connect.WsLogic
}

func NewApi(ctx context.Context, profile *config.Profile, lager *zap.Logger, cache *cache.Cache, actor *queue.ActorPool, wsLogic *connect.WsLogic) *Api {
	return &Api{
		Context: ctx,
		Profile: profile,
		Lager:   lager,
		Cache:   cache,
		Actor:   actor,
		WsLogic: wsLogic,
	}
}

func (v *Api) Start(req *http.Request) (*http.Request, func()) {

	// 准备日志资料
	return req, func() {
	}
}

type Request interface {
	Validate() error
}

// 记录参数
func (v *Api) PostV(req *http.Request, vars Request) error {
	pv := req.Context().Value(muxhttp.ContextKey("post_value")).(*muxhttp.HttpPostJson)
	return pv.Unmarshal(&vars)
}

func (v *Api) Get(key string, loader func(context.Context) (interface{}, time.Duration, error)) (any, error) {
	return v.Cache.GetOrLoadCtx(context.Background(), key, loader)
}

func (v *Api) Tell(msg queue.Message) {
	v.Actor.Tell(msg)
}
