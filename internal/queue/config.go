package queue

import "time"

// 自动扩缩容配置
type AutoScalerConfig struct {
	Interval      time.Duration
	HighThreshold int
	LowThreshold  int
	ScaleUpStep   int
	ScaleDownStep int
	Cooldown      time.Duration
}
