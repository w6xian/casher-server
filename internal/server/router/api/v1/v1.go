package v1

import (
	"casher-server/api/wrpc"
	"casher-server/internal/config"
	"casher-server/internal/muxhttp"
	"casher-server/internal/queue"
	"casher-server/internal/store"
	"casher-server/internal/utils/def"
	"casher-server/internal/utils/id"
	"context"
	"fmt"
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
	WsProxy *wrpc.WSProxy
	Store   *store.Store
}

func NewApi(ctx context.Context, storeInstance *store.Store, profile *config.Profile, lager *zap.Logger, cache *cache.Cache, actor *queue.ActorPool, wsProxy *wrpc.WSProxy) *Api {
	return &Api{
		Context: ctx,
		Profile: profile,
		Lager:   lager,
		Cache:   cache,
		Actor:   actor,
		WsProxy: wsProxy,
		Store:   storeInstance,
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

func (v *Api) DbConnectWithClose(ctx context.Context) (context.Context, func()) {
	return v.Store.GetConnect(ctx), func() {
		v.Store.CloseConnect(ctx)
	}
}
func (v *Api) AnonymousTracker(req *http.Request) (*store.Tracker, error) {
	tracker := store.NewAnonimousTracker(req.Context())
	tracker.TrackId = muxhttp.GetRequestId(req)
	tracker.Language = muxhttp.GetLanguage(req)
	tracker.MachineNo = v.Profile.Machine.Id
	tracker.MachineId = v.Profile.Machine.Code
	return tracker, nil
}
func (v *Api) GetTracker(req *http.Request, checkLock bool) (*store.Tracker, error) {
	if !checkLock {
		return v.AnonymousTracker(req)
	}
	nextId, err := id.NextId(def.GetNumber(v.Profile.Machine.Id, 1))
	if err != nil {
		return nil, err
	}
	tracker := store.NewTracker()
	tracker.TrackId = fmt.Sprintf("%d", nextId)
	tracker.MachineNo = v.Profile.Machine.Id
	tracker.MachineId = v.Profile.Machine.Code
	tracker.Language = muxhttp.GetLanguage(req)
	return tracker, nil
}
