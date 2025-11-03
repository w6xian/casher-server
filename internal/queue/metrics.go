package queue

import "github.com/prometheus/client_golang/prometheus"

// ---------------- Metrics ----------------
// Actor 队列长度指标
var (
	ActorQueueLength = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "actor_queue_length", Help: "每个 actor 当前待处理消息数"},
		[]string{"actor"},
	)
	// Actor 池大小指标
	ActorPoolSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "actor_pool_size", Help: "当前 actor 池大小"},
		[]string{"pool"},
	)
	// Actor 重启次数指标
	ActorRestarts = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "actor_restarts_total", Help: "actor 重启次数"},
		[]string{"actor"},
	)
)
