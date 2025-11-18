package rpc

import (
	"casher-server/internal/i18n"
	"casher-server/internal/store"
	"context"
	"fmt"
)

// GetBySn 获取公司信息
func (c *Shop) GetCompanyBySn(ctx context.Context, req *store.CompanySnReq, reply *store.CompanyReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("GetCompanyBySn recover: %v\r\n", err)
		}
	}()
	//空方法
	ctx, stop := c.Start(ctx)
	defer stop()
	// 1 获取日志资料
	lang := c.GetTracker(ctx, req)
	// 2 校验请求签名
	err := store.CheckSign(req, req.AppId)
	if err != nil {
		return lang.Error("get_company_by_sn_req_sign", "请求签名校验失败:{{.error}}", i18n.String("error", err.Error()))
	}
	// 1 初始化数据库连接
	ctx, close := c.Store.DbConnectWithClose(ctx)
	defer close()
	req.Tracker = lang
	// 4 校验请求参数
	vErr := req.Validate()
	if vErr != nil {
		return lang.Error("get_company_by_sn_validate", "请求参数校验失败:{{.error}}", i18n.String("error", vErr.Error()))
	}
	// 2 调用数据库查询商品信息
	err = c.Store.GetCluldCompaniesBySn(ctx, req, reply)
	if err != nil {
		return lang.Error("get_company_by_sn_reply", "查询公司信息失败:{{.error}}", i18n.String("error", err.Error()))
	}
	// 校验返回签名
	err = store.SetSign(reply, req.AppId)
	if err != nil {
		return lang.Error("get_company_by_sn_reply_sign", "设置签名失败:{{.error}}", i18n.String("error", err.Error()))
	}
	return nil
}

// GetByName 获取公司信息
func (c *Shop) GetCompanyByName(ctx context.Context, req *store.CompanyNameReq, reply *store.CompanyReply) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("GetCompanyByName recover: %v\r\n", err)
		}
	}()
	//空方法
	ctx, stop := c.Start(ctx)
	defer stop()
	// 1 获取日志资料
	lang := c.GetTracker(ctx, req)
	// 2 校验请求签名
	err := store.CheckSign(req, req.AppId)
	if err != nil {
		return lang.Error("get_company_by_name_req_sign", "请求签名校验失败:{{.error}}", i18n.String("error", err.Error()))
	}
	// 1 初始化数据库连接
	ctx, close := c.Store.DbConnectWithClose(ctx)
	defer close()
	req.Tracker = lang
	// 4 校验请求参数
	vErr := req.Validate()
	if vErr != nil {
		return lang.Error("get_company_by_name_validate", "请求参数校验失败:{{.error}}", i18n.String("error", vErr.Error()))
	}
	// 2 调用数据库查询商品信息
	err = c.Store.GetCluldCompaniesByName(ctx, req, reply)
	if err != nil {
		return lang.Error("get_company_by_name_reply", "查询公司信息失败:{{.error}}", i18n.String("error", err.Error()))
	}
	// 校验返回签名
	err = store.SetSign(reply, req.AppId)
	if err != nil {
		return lang.Error("get_company_by_name_reply_sign", "设置签名失败:{{.error}}", i18n.String("error", err.Error()))
	}
	return nil
}
