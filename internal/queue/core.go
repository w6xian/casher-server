package queue

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

// Actor 系统，管理所有子 Actor
type ActorSystem struct {
	mu       sync.Mutex
	children map[string]*Child
	wg       sync.WaitGroup
}

// 子 Actor 结构体
// name: 名称
// props: 属性
// mailbox: 消息队列
// ref: Actor 引用
// stop: 停止信号
// closing: 关闭信号
// restarts: 重启时间记录
type Child struct {
	Name     string
	Props    Props
	Mailbox  chan []byte
	Ref      ActorRef
	Stop     chan struct{}
	Closing  chan struct{}
	Mu       sync.Mutex
	Restarts []time.Time
}

// 获取待处理消息数
func (c *Child) Pending() int { return len(c.Mailbox) }

// 更新队列长度指标
func (c *Child) UpdateMetrics() {
	ActorQueueLength.WithLabelValues(c.Name).Set(float64(len(c.Mailbox)))
}

// 增加重启次数指标
func (c *Child) IncrementRestartMetric() { ActorRestarts.WithLabelValues(c.Name).Inc() }

// 创建 Actor 系统
func NewSystem() *ActorSystem { return &ActorSystem{children: make(map[string]*Child)} }

// 创建 Actor
func (s *ActorSystem) Spawn(name string, props Props) ActorRef { return s.spawnChild(name, props, nil) }

// 创建子 Actor，支持父子关系
func (s *ActorSystem) spawnChild(name string, props Props, parent *ActorRef) ActorRef {
	if props.Mailbox <= 0 {
		props.Mailbox = 16 // 默认邮箱容量
	}
	if props.Window == 0 {
		props.Window = 5 * time.Second // 默认重启窗口
	}
	c := &Child{Name: name, Props: props, Mailbox: make(chan []byte, props.Mailbox), Stop: make(chan struct{}), Closing: make(chan struct{})}
	c.Ref = ActorRef{path: name, send: c.Mailbox, done: c.Closing}

	// 加入系统管理
	s.mu.Lock()
	if _, exists := s.children[name]; exists {
		s.mu.Unlock()
		panic("actor name already exists: " + name)
	}
	s.children[name] = c
	s.mu.Unlock()

	// 启动 Actor
	s.runChild(c)
	return c.Ref
}

// 启动子 Actor 的主循环，处理消息、异常与重启
func (s *ActorSystem) runChild(c *Child) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			// 每次重启都新建 Context
			ctx := Context{Self: c.Ref, state: make(map[string]interface{}), sys: s, stop: make(chan struct{})}
			exit := make(chan interface{})
			go func() {
				defer func() {
					if r := recover(); r != nil {
						exit <- r // 捕获异常
						return
					}
					exit <- nil // 正常退出
				}()
				for {
					select {
					case msg := <-c.Mailbox:
						c.Props.ActorFunc(ctx, msg) // 处理消息
					case <-ctx.stop:
						return // 主动停止
					case <-c.Stop:
						return // 强制停止
					case <-c.Closing:
						// 关闭时清空队列
					drain:
						for {
							select {
							case msg := <-c.Mailbox:
								c.Props.ActorFunc(ctx, msg)
							default:
								break drain
							}
						}
						return
					}
				}
			}()

			res := <-exit
			if res == nil {
				// 正常退出，移除 Actor
				s.mu.Lock()
				delete(s.children, c.Name)
				s.mu.Unlock()
				return
			}

			// 异常退出，打印错误和堆栈
			fmt.Printf("[ERROR] actor %s 崩溃: %v\nStack:\n%s\n", c.Name, res, string(debug.Stack()))
			c.IncrementRestartMetric()
			c.Mu.Lock()
			now := time.Now()
			c.Restarts = append(c.Restarts, now)
			cut := now.Add(-c.Props.Window)
			j := 0
			for i := len(c.Restarts) - 1; i >= 0; i-- {
				if c.Restarts[i].After(cut) {
					j = i
				}
			}
			c.Restarts = c.Restarts[j:]
			count := len(c.Restarts)
			c.Mu.Unlock()

			// 判断是否重启
			shouldRestart := false
			switch c.Props.Strategy {
			case RestartAlways:
				shouldRestart = true
			case RestartOnFailure:
				if count <= c.Props.MaxRestarts || c.Props.MaxRestarts == 0 {
					shouldRestart = true
				}
			}

			if !shouldRestart {
				fmt.Printf("[SUPERVISOR] 不重启 %s\n", c.Name)
				s.mu.Lock()
				delete(s.children, c.Name)
				s.mu.Unlock()
				return
			}

			// 重启退避时间，防止频繁重启
			backoff := time.Millisecond * 100 * time.Duration(count)
			if backoff > 5*time.Second {
				backoff = 5 * time.Second
			}
			fmt.Printf("[SUPERVISOR] %s %v 后重启 (count=%d)\n", c.Name, backoff, count)
			time.Sleep(backoff)
		}
	}()
}

// 优雅关闭 Actor 系统
func (s *ActorSystem) Shutdown(ctx context.Context) error {
	s.mu.Lock()
	for _, c := range s.children {
		select {
		case <-c.Closing:
		default:
			close(c.Closing)
		}
	}
	s.mu.Unlock()

	done := make(chan struct{})
	go func() { s.wg.Wait(); close(done) }()

	select {
	case <-done:
		return nil // 正常关闭
	case <-ctx.Done():
		// 超时后强制停止所有 Actor
		s.mu.Lock()
		for _, c := range s.children {
			select {
			case <-c.Stop:
			default:
				close(c.Stop)
			}
		}
		s.mu.Unlock()
		return ctx.Err()
	}
}
