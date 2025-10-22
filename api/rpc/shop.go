package rpc

import (
	"casher-server/internal/errors"
	"casher-server/internal/i18n"
	"casher-server/internal/lager"
	"casher-server/internal/store"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"golang.org/x/text/language"
)

// :: 启动方法
func (v *Shop) Start(ctx context.Context) (context.Context, func()) {
	ctx = lager.RequestLager(ctx, v.Lager)
	// 准备日志资料
	return ctx, func() {

	}
}

func (v *Shop) L(key string, def string, fields ...i18n.Field) string {
	if v.Language == "" {
		v.Language = language.Chinese.String()
	}
	l := len(fields)
	if l == 0 {
		return i18n.T(v.Language, key, def)
	}

	data := i18n.D{}
	for _, f := range fields {
		data[f.Key] = f.Value()
	}
	for k, v := range data {
		fmt.Printf("%s=%s\n", k, v)
	}
	return i18n.TWithData(v.Language, key, def, data)
}

func (v *Shop) Error(key string, def string, fields ...i18n.Field) *errors.ErrorL {
	return errors.FromLang(v)
}

//!! 基础方法完

func (c *Shop) Register(ctx context.Context, req *RegisterRequest, reply *RegisterReply) error {
	md5Sum := MD5(req.Name + req.Password)
	reply.Code = 0
	reply.AuthToken = string(md5Sum[:]) + "server Register"
	return nil
}

func (c *Shop) Login(ctx context.Context, req *LoginRequest, reply *LoginReply) error {
	md5Sum := MD5(req.Name + req.Password)
	reply.Code = 0
	reply.AuthToken = string(md5Sum[:]) + "server Login"
	return nil
}

func MD5(input string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(input))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

// SyncShopInfo 同步店铺信息
func (c *Shop) SyncShopInfo(ctx context.Context, req *store.ShopInfoReq, reply *store.ShopInfoReqReply) error {
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
	return c.Store.SyncShopInfo(ctx, req, reply)
}

// ShopLink 店铺链接
func (c *Shop) ShopLink(ctx context.Context, req *store.ShopLinkReq, reply *store.ShopLinkReqReply) error {
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
	return c.Store.ShopLink(ctx, req, reply)
}
