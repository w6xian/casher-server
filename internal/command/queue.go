package command

import (
	"casher-server/api/wrpc"
	"casher-server/internal/config"
	"casher-server/internal/queue"
	"casher-server/internal/store"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/louis-xie-programmer/go-local-cache/cache"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

type Command struct {
	Context context.Context
	Profile *config.Profile
	Lager   *zap.Logger
	Cache   *cache.Cache
	WsProxy *wrpc.WSProxy
	Store   *store.Store
}

func NewQueueCommand(ctx context.Context, opt *config.Profile, store *store.Store, lager *zap.Logger, cache *cache.Cache, wsProxy *wrpc.WSProxy) *Command {
	return &Command{
		Context: ctx,
		Profile: opt,
		Lager:   lager,
		Cache:   cache,
		WsProxy: wsProxy,
		Store:   store,
	}
}

// curl -d "{\"id\":\"123456\",\"action\":1,\"data\":\"YWJj\"}" http://127.0.0.1:4151/pub?topic=CASH_SERVER_QUEUE
func (c *Command) HandleMessage(message *nsq.Message) error {
	msg := queue.ActorMessage{}
	err := json.Unmarshal(message.Body, &msg)
	if err != nil {
		fmt.Printf("[ERROR] 解析消息失败: %v\n", err)
		return nil
	}
	switch msg.Action {
	case 1: // 订单处理
		ctx, close := c.Store.DbApiConnectWithClose(context.Background())
		defer close()
		c.Store.NoticeNewOrder(ctx, &store.WsReq{
			AppId:   "610800923266441381",
			UserId:  10,
			Tracker: store.NewAnonimousTracker(ctx),
		})
	case 2:
	case 0xFF:
		fmt.Printf("[%s] 处理 %s\n", msg.Id, msg.Data)
	}
	return nil
}

func (c *Command) ActorFunc(qctx queue.Context, data []byte) {
	msg := queue.Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Printf("[ERROR] 解析消息失败: %v\n", err)
		return
	}
	switch msg.Action {
	case 1:
		time.Sleep(2 * time.Second)
		fmt.Printf("[%s] 处理 %s\n", qctx.Self.Path(), msg.Data)

	case 2:
		if msg.Data == "fail" {
			panic("模拟故障")
		}
	case 0xFF:
		fmt.Printf("[%s] 处理 %s\n", qctx.Self.Path(), msg.Data)
	}
}
