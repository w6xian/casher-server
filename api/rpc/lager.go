package rpc

import (
	"casher-server/internal/lager"
	"context"
)

func (v *Order) getLager(ctx context.Context) *lager.LogReq {
	return lager.FromContext(ctx)
}

// query查询
func (v *Order) Query(ctx context.Context, title string) {
	logReq := v.getLager(ctx)
	logReq.OperationTitle = title
	logReq.OperationType = "query"
}

// 操作
func (v *Order) Event(ctx context.Context, title string) {
	logReq := v.getLager(ctx)
	logReq.OperationTitle = title
	logReq.OperationType = "event"
}

// 查看
func (v *Order) View(ctx context.Context, title string) {
	logReq := v.getLager(ctx)
	logReq.OperationTitle = title
	logReq.OperationType = "view"
}
