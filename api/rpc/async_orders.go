package rpc

import (
	"casher-server/internal/store"
	"casher-server/proto"
	"context"
	"fmt"
)

// AsyncOrders 异步订单信息
func (c *Shop) AsyncOrders(ctx context.Context, req *store.AsyncRequest, reply *store.AsyncOrdersReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("AsyncOrders recover: %v\r\n", err)
		}
	}()
	//空方法
	ctx, stop := c.Start(ctx)
	defer stop()
	// 校验返回签名
	err := proto.CheckSign(req, req.AppId)
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
	// 2 调用数据库查询商品信息
	err = c.Store.AsyncOrders(ctx, req, reply)
	if err != nil {
		return err
	}
	// 校验返回签名
	err = proto.SetSign(reply, req.AppId)
	if err != nil {
		return err
	}
	return nil
}
