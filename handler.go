package main

import (
	"casher-server/internal/config"
	"casher-server/internal/store"
	"context"
	"fmt"
	"time"

	"github.com/w6xian/sloth/bucket"
	"github.com/w6xian/sloth/nrpc/wsocket"
	"go.uber.org/zap"
)

type Handler struct {
	Profile  *config.Profile
	Lager    *zap.Logger
	Store    *store.Store
	Language string
}

func NewHandler(profile *config.Profile, logger *zap.Logger, store *store.Store, language string) *Handler {
	return &Handler{
		Profile:  profile,
		Lager:    logger,
		Store:    store,
		Language: language,
	}
}

// OnClose implements wsocket.IServerHandleMessage.
func (h *Handler) OnClose(ctx context.Context, s *wsocket.WsServer, ch bucket.IChannel) error {
	fmt.Println("OnClose")
	return nil
}

// OnError implements wsocket.IServerHandleMessage.
func (h *Handler) OnError(ctx context.Context, s *wsocket.WsServer, ch bucket.IChannel, err error) error {
	fmt.Println("OnError:", err)
	return nil
}

// OnOpen implements wsocket.IServerHandleMessage.
func (h *Handler) OnOpen(ctx context.Context, s *wsocket.WsServer, ch bucket.IChannel) error {
	fmt.Println("OnOpen")
	return nil
}

func (h *Handler) OnData(ctx context.Context, s *wsocket.WsServer, ch bucket.IChannel, msgType int, message []byte) error {
	fmt.Println("OnData:", msgType, string(message), ch.UserId())
	return nil
}

type HelloService struct {
	Id int64 `json:"id"`
}

func (h *HelloService) Test(ctx context.Context, data []byte) (any, error) {
	h.Id = h.Id + 1
	fmt.Println("Test args:", string(data))
	if h.Id%5 == 1 {
		return nil, fmt.Errorf("error %d", h.Id)
	}
	return map[string]string{"req": "server 1", "time": time.Now().Format("2006-01-02 15:04:05")}, nil
}
func (h *HelloService) Login(ctx context.Context, data []byte) (any, error) {
	return map[string]string{"user_id": "2", "time": time.Now().Format("2006-01-02 15:04:05")}, nil
}
