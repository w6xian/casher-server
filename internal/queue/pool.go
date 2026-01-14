package queue

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Actor 池，管理一组 worker Actor
type ActorPool struct {
	Name     string
	System   *ActorSystem
	Props    Props
	Min, Max int
	Workers  []*Child
	M        sync.Mutex
	Next     int
	Scaling  struct {
		cfg  AutoScalerConfig
		stop chan struct{}
		cool time.Time
	}
}

func NewPool(sys *ActorSystem, name string, min, max int) *ActorPool {
	if min < 1 {
		min = 1 // 最小池大小
	}
	if max < min {
		max = min // 最大池不能小于最小池
	}
	p := &ActorPool{Name: name, System: sys, Min: min, Max: max}
	return p
}

func (p *ActorPool) Start() {

	if p.Props.ActorFunc == nil {
		panic("ActorFunc 不能为空")
	}
	p.Resize(p.Min)
}

// 调整池大小
func (p *ActorPool) Resize(n int) {
	p.M.Lock()
	defer p.M.Unlock()
	if n < p.Min {
		n = p.Min
	}
	if n > p.Max {
		n = p.Max
	}
	cur := len(p.Workers)
	if n > cur {
		for i := cur; i < n; i++ {
			name := fmt.Sprintf("%s-worker-%d", p.Name, i)
			p.System.spawnChild(name, p.Props, nil)
			child := p.System.children[name]
			p.Workers = append(p.Workers, child)
		}
		fmt.Printf("[POOL] 扩容 %s: %d -> %d\n", p.Name, cur, n)
	} else if n < cur {
		p.Workers = p.Workers[:n]
		fmt.Printf("[POOL] 缩容 %s: %d -> %d\n", p.Name, cur, n)
	}
}

// 向池中的 worker 轮询发送消息
// internal/command/queue.go
func (p *ActorPool) Tell(msg []byte) {
	p.M.Lock()
	if len(p.Workers) == 0 {
		p.M.Unlock()
		fmt.Printf("[POOL] %s 没有 worker\n", p.Name)
		return
	}
	w := p.Workers[p.Next%len(p.Workers)]
	p.Next = (p.Next + 1) % len(p.Workers)
	p.M.Unlock()
	w.Ref.Tell(msg)
}
func (p *ActorPool) SetProps(props Props) {
	p.Props = props
}

// 获取池大小
func (p *ActorPool) Len() int { p.M.Lock(); defer p.M.Unlock(); return len(p.Workers) }

// 更新池大小指标
func (p *ActorPool) UpdatePoolMetrics() { ActorPoolSize.WithLabelValues(p.Name).Set(float64(p.Len())) }

// 启动自动扩缩容
func (p *ActorPool) StartAutoScaler(cfg AutoScalerConfig) {
	p.Start()
	p.Scaling.cfg = cfg
	p.Scaling.stop = make(chan struct{})
	p.Scaling.cool = time.Now()
	go func() {
		ticker := time.NewTicker(cfg.Interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				p.M.Lock()
				if time.Since(p.Scaling.cool) < cfg.Cooldown {
					p.M.Unlock()
					continue // 冷却期不扩缩容
				}
				total := 0
				for _, w := range p.Workers {
					total += w.Pending()
				}
				avg := 0
				if len(p.Workers) > 0 {
					avg = total / len(p.Workers)
				}
				p.M.Unlock()
				if avg > cfg.HighThreshold {
					p.Resize(p.Len() + cfg.ScaleUpStep)
					p.Scaling.cool = time.Now()
				}
				if avg < cfg.LowThreshold {
					p.Resize(p.Len() - cfg.ScaleDownStep)
					p.Scaling.cool = time.Now()
				}
			case <-p.Scaling.stop:
				fmt.Printf("[AUTOSCALE] %s 自动扩缩容已停止\n", p.Name)
				return
			}
		}
	}()
}

// 停止自动扩缩容
func (p *ActorPool) StopAutoScaler() {
	if p.Scaling.stop != nil {
		close(p.Scaling.stop)
	}
}
func (p *ActorPool) Shutdown(ctx context.Context) error { p.StopAutoScaler(); return nil }
