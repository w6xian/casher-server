package rpc

import (
	"casher-server/internal/lager"
	"context"
)

func (v *Shop) getLager(ctx context.Context) *lager.LogReq {
	return lager.FromContext(ctx)
}

// query查询
func (v *Shop) Query(ctx context.Context, title string) {
	logReq := v.getLager(ctx)
	logReq.OperationTitle = title
	logReq.OperationType = "query"
}

// 操作
func (v *Shop) Event(ctx context.Context, title string) {
	logReq := v.getLager(ctx)
	logReq.OperationTitle = title
	logReq.OperationType = "event"
}

// 查看
func (v *Shop) View(ctx context.Context, title string) {
	logReq := v.getLager(ctx)
	logReq.OperationTitle = title
	logReq.OperationType = "view"
}
