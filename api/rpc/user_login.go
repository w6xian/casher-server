package rpc

import (
	"casher-server/internal/store"
	"context"
	"fmt"
)

func (c *Shop) AuthLogin(ctx context.Context, req *store.LoginRequest, reply *store.LoginReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("AsyncUsersExtra recover: %v\r\n", err)
		}
	}()
	//空方法
	ctx, stop := c.Start(ctx)
	defer stop()
	// 1 初始化数据库连接
	ctx, close := c.Store.DbConnectWithClose(ctx)
	defer close()

	// 通过api_key 校验商户是否存在
	authInfo, err := c.Store.GetAuthInfo(ctx, req.MchId, req.ApiKey)
	if err != nil {
		return err
	}
	// 校验返回签名
	err = checkSign(req, authInfo.MchId, authInfo.ApiKey, authInfo.ApiSecret)
	if err != nil {
		return err
	}

	// 3 获取日志资料
	tracker := c.GetTracker(ctx, req)
	req.Tracker = tracker
	vErr := req.Validate()
	if vErr != nil {
		return vErr
	}
	// 2 调用数据库查询会员信息
	err = c.Store.Login(ctx, req, reply)
	if err != nil {
		return err
	}
	reply.OpenId = authInfo.OpenId
	reply.UserId = authInfo.UserId
	reply.RoomId = authInfo.ShopId
	// 校验返回签名
	err = setSign(reply, authInfo.MchId, authInfo.ApiKey, authInfo.ApiSecret)
	if err != nil {
		return err
	}
	return nil
}
