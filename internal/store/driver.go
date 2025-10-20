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
}
