package rpc

import (
	"casher-server/internal/store"
	"context"
	"fmt"
)

// AsyncProducts 异步商品信息
func (c *Shop) AsyncProducts(ctx context.Context, req *store.AsyncRequest, reply *store.AsyncProductsReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("AsyncProducts recover: %v\r\n", err)
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
	// 2 调用数据库查询商品信息
	err = c.Store.AsyncProducts(ctx, req, reply)
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

// AsyncProductsExtra 异步商品信息
func (c *Shop) AsyncProductsExtra(ctx context.Context, req *store.AsyncRequest, reply *store.AsyncProductsExtraReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("AsyncProductsExtra recover: %v\r\n", err)
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
	// 2 调用数据库查询商品信息
	err = c.Store.AsyncProductsExtra(ctx, req, reply)
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
