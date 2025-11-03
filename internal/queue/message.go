package queue

import "context"

// Action 枚举
type Action int

type Message struct {
	Context context.Context
	Id      string
	Action  Action `json:"action"`
	Data    any    `json:"data"`
	Tracker any    `json:"tracker"`
}
