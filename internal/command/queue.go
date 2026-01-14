package command

import (
	"casher-server/internal/config"
	"casher-server/internal/queue"
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
}

func NewQueueCommand(ctx context.Context, opt *config.Profile, lager *zap.Logger, cache *cache.Cache) *Command {
	return &Command{
		Context: ctx,
		Profile: opt,
		Lager:   lager,
		Cache:   cache,
	}
}

func (c *Command) HandleMessage(message *nsq.Message) error {
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
