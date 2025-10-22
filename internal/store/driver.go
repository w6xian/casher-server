package store

import (
	"context"

	"github.com/w6xian/sqlm"
)

type Driver interface {
	// store基础

	Close() error
	GetConnect(ctx context.Context) context.Context
	CloseConnect(ctx context.Context) error
	GetAction(ctx context.Context) *sqlm.Db
	GetLink(ctx context.Context) sqlm.ITable

	// 授权信息
	GetAuthByComId(link sqlm.ITable, comId int64) (*Auths, error)
	InsertAuth(link sqlm.ITable, auth *Auths) (int64, error)

	// 公司信息
	GetCRMCompanyById(link sqlm.ITable, id int64) (*CompInfo, error)
	GetCRMCompanyByOpenId(link sqlm.ITable, openId string) (*CompInfo, error)
	// 代理商信息
	GetProxyInfoById(link sqlm.ITable, id int64) (*ProxyInfo, error)
	CreateProxyLite(link sqlm.ITable, req *ShopInfoReq) (int64, error)
	// 店铺
	GetShopByAppId(link sqlm.ITable, appId string) (*ShopLinkReqReply, error)
	// 公司管理员
	GetComAdmin(link sqlm.ITable, proxyId int64, mobile string) (*Admin, error)
	//lagor
	GetMap(link sqlm.ITable, tableName string, pk string, value string) map[string]any
	GetMapById(link sqlm.ITable, tableName string, id int64) map[string]any
	SelectMap(link sqlm.ITable, tableName string, pk string, value string) []map[string]any
}
