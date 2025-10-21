package rpc

import (
	"casher-server/internal/store"
	"context"
	"fmt"
)

// ReadProductSuppUseCode 读取产品供应商使用码
func (c *Shop) ReadProductSuppUseCode(ctx context.Context, req *store.SuppCodeRequest, reply *store.SuppCodeReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ReadProductSuppUseCode recover: %v\r\n", err)
		}
	}()
	//空方法
	ctx, stop := c.Start(ctx)
	defer stop()
	// 1 初始化数据库连接
	ctx, close := c.Store.DbConnectWithClose(ctx)
	defer close()
	return c.Store.ReadProductSuppUseCode(ctx, req, reply)
}
