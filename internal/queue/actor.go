package queue

import (
	"fmt"
	"time"
)

// Actor 重启策略类型
type RestartStrategy int

// Actor 重启策略枚举
const (
	RestartNever     RestartStrategy = iota // 永不重启
	RestartAlways                           // 总是重启
	RestartOnFailure                        // 失败时重启
)

// Actor 属性配置
// ActorFunc: 处理消息的函数
// Mailbox: 邮箱容量
// Strategy: 重启策略
// MaxRestarts: 最大重启次数
// Window: 重启计数窗口
type Props struct {
	ActorFunc   func(ctx Context, msg Message)
	Mailbox     int
	Strategy    RestartStrategy
	MaxRestarts int
	Window      time.Duration
}

// Actor 引用，包含路径、消息通道、关闭信号
type ActorRef struct {
	path string
	send chan<- Message
	done <-chan struct{}
}

// 向 Actor 发送消息
func (r ActorRef) Tell(msg Message) {
	select {
	case <-r.done:
		fmt.Printf("[WARN] actor %s 正在关闭，丢弃消息 %#v\n", r.path, msg)
	case r.send <- msg:
	}
}

// 获取 Actor 路径
func (r ActorRef) Path() string { return r.path }

// Actor 上下文，包含自身引用、状态、系统指针、停止信号
type Context struct {
	Self  ActorRef
	state map[string]interface{}
	sys   *ActorSystem
	stop  chan struct{}
}

// 创建子 Actor
func (c *Context) SpawnChild(name string, props Props) ActorRef {
	return c.sys.spawnChild(name, props, &c.Self)
}

// 停止 Actor
func (c *Context) Stop() { close(c.stop) }

// ---------------- Actor System ----------------
