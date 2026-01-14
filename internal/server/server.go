package server

import (
	"casher-server/api/wrpc"
	"casher-server/internal/command"
	"casher-server/internal/config"
	"casher-server/internal/muxhttp/mw"
	"casher-server/internal/queue"
	"casher-server/internal/server/router"
	v1 "casher-server/internal/server/router/api/v1"
	"casher-server/internal/server/router/frontend"
	"casher-server/internal/store"
	"casher-server/internal/wsfuns"
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	// "github.com/patrickmn/go-cache"
	"github.com/louis-xie-programmer/go-local-cache/cache"
	"github.com/pkg/errors"
	"github.com/w6xian/sloth"
	"github.com/w6xian/sloth/nrpc/wsocket"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// TCP服务
type Server struct {
	Profile           *config.Profile
	runnerCancelFuncs []context.CancelFunc
	Router            *mux.Router
	Store             *store.Store
	Lager             *zap.Logger
	Cache             *cache.Cache
	Actor             *queue.ActorPool
	WsProxy           *wrpc.WSProxy
}

func NewServer(ctx context.Context, opt *config.Profile, store *store.Store, lager *zap.Logger, cache *cache.Cache, actorPool *queue.ActorPool, wsProxy *wrpc.WSProxy) (*Server, error) {
	s := &Server{
		Profile: opt,
		Store:   store,
		Lager:   lager,
		Cache:   cache,
		Actor:   actorPool,
		WsProxy: wsProxy,
	}
	s.Router = mux.NewRouter()
	s.Router.Use(mw.CORSMethodMiddleware(opt.Server.Origins))
	// Initialize profiler
	s.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Lager.Info("404 page not found", zap.String("path", r.URL.Path))
		http.Redirect(w, r, "/setting/", http.StatusFound)
	})
	// s.profiler.StartMemoryMonitor(ctx)
	return s, nil
}

func (s *Server) Start(ctx context.Context, queueCmd *command.Command) error {
	var address, network string
	address = s.Profile.Server.Address
	network = s.Profile.Server.Network
	listener, err := net.Listen(network, address)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	go func() {
		r := mux.NewRouter()
		ln, err := net.Listen("tcp", s.Profile.Server.WsAddr)
		if err != nil {
			panic(err)
		}
		httpsvr := &http.Server{
			Handler: r,
		}

		handler := NewHandler(s.Profile, s.Lager, s.Store, s.Profile.Apps.Language)
		wsServerApi := wsfuns.NewWsServerApi(s.Profile, s.Lager, s.Store, s.Profile.Apps.Language)
		newConnect := sloth.ServerConn(s.WsProxy.Server)
		newConnect.RegisterRpc("v1", wsServerApi, "")
		newConnect.ListenOption(
			wsocket.WithRouter(r),
			wsocket.WithServerHandle(handler),
		)
		httpsvr.Serve(ln)
	}()

	// http服务
	go func() {
		// httpListener := muxServer.Match(cmux.HTTP1Fast())
		// 注册UI
		frontend.NewFrontendService("setting").Serve(ctx, s.Router)
		// 注册文件管理
		prefix := "/statics/"
		s.Router.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(http.Dir("./statics"))))
		// 注册方法
		api := v1.NewApi(ctx, s.Store, s.Profile, s.Lager, s.Cache, s.Actor, s.WsProxy)
		router.Register(ctx, s.Router, api)
		// 绑定路由到Http
		http.Handle("/", s.Router)
		http.Serve(listener, nil)
	}()

	fmt.Println("-------------http server started", listener.Addr().String())
	s.StartBackgroundRunners(ctx)
	fmt.Println("server started", listener.Addr().String())
	return nil
}

func (s *Server) Shutdown(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	slog.Info("server shutting down")

	for _, cancelFunc := range s.runnerCancelFuncs {
		if cancelFunc != nil {
			cancelFunc()
		}
	}

	if err := s.Store.Close(); err != nil {
		slog.Error("failed to close database", slog.String("error", err.Error()))
	}
	select {
	case <-ctx.Done():
		slog.Info("server shutdown timeout", slog.String("error", ctx.Err().Error()))
	default:
		slog.Info("server shutdown completed")
	}
}

func (s *Server) StartBackgroundRunners(ctx context.Context) {
	// Create a separate context for each background runner
	// This allows us to control cancellation for each runner independently
	_, s3Cancel := context.WithCancel(ctx)

	// Store the cancel function so we can properly shut down runners
	s.runnerCancelFuncs = append(s.runnerCancelFuncs, s3Cancel)

}
