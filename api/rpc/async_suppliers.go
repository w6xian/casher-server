package rpc

import (
	"casher-server/internal/store"
	"context"
	"fmt"
)

// AsyncSuppliers 异步供应商信息
func (c *Shop) AsyncSuppliers(ctx context.Context, req *store.AsyncRequest, reply *store.AsyncSuppliersReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("AsyncSuppliers recover: %v\r\n", err)
		}
	}()
	//空方法
	ctx, stop := c.Start(ctx)
	defer stop()
	// 校验返回签名
	err := checkSign(req, req.AppId)
	if err != nil {
		return err
	}
	// 1 初始化数据库连接
	ctx, close := c.Store.DbConnectWithClose(ctx)
	defer close()
	// 3 获取日志资料
	tracker := c.GetTracker(ctx, req)
	req.Tracker = tracker
	vErr := req.Validate()
	if vErr != nil {
		return vErr
	}
	// 2 调用数据库查询供应商信息
	err = c.Store.AsyncSuppliers(ctx, req, reply)
	if err != nil {
		return err
	}
	// 校验返回签名
	err = setSign(reply, req.AppId)
	if err != nil {
		return err
	}
	return nil
}
