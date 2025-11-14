package rpc

import (
	"casher-server/internal/store"
	"casher-server/proto"
	"context"
	"fmt"
)

// GetPrdBySn 获取商品信息
func (c *Shop) ReqPrdBySn(ctx context.Context, req *store.PrdSnReq, reply *store.PublicProductReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ReqPrdBySn recover: %v\r\n", err)
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
	err = c.Store.GetPublicProductBySn(ctx, req, reply)
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
