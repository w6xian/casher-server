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
	// 通过高品条形码查询商品信息
	GetPublicProductBySn(link sqlm.ITable, sn string) (*ProductModel, error)
	// 新增商品
	InsertPublicProduct(link sqlm.ITable, prd *ProductModel) (int64, error)
	InsertPublicProductVersion(link sqlm.ITable, prd *ProductModel) (int64, error)
	// GetPublicCompanyBySn 通过高品条形码查询公司信息
	GetPublicCompanyBySn(link sqlm.ITable, sn string) (*CompanyModel, error)
	// GetPublicCompanyByName 通过高品公司名称查询公司信息
	GetPublicCompanyByName(link sqlm.ITable, name string) (*CompanyModel, error)
	// GetPublicCompanyById 根据公司ID查询公司
	GetPublicCompanyById(link sqlm.ITable, id int64) (*CompanyModel, error)
	// QueryPublicCompanyBySn 根据公司编码查询公司
	QueryPublicCompanyBySn(link sqlm.ITable, sn string) ([]*CompanyLiteModel, error)
	// QueryPublicCompanyByName 根据公司名称查询公司

	//订单
	QueryOrders(link sqlm.ITable, req *AsyncRequest) (*AsyncOrdersReply, error)
	SelectOrderItems(link sqlm.ITable, orderId int64) ([]*OrderLiteItem, error)
	//商品
	QueryProducts(link sqlm.ITable, req *AsyncRequest) (*AsyncProductsReply, error)
	QueryProductsExtra(link sqlm.ITable, req *AsyncRequest) (*AsyncProductsExtraReply, error)
	QueryProductUpdate(link sqlm.ITable, req *IdRequest) (*ProductLite, error)
	AsyncUpdateProduct(link sqlm.ITable, req *UpdateRequest, kv map[string]any) (int64, error)
	GetProductByUnionId(link sqlm.ITable, proxyId int64, shopId int64, unionId string) (*ProductLite, error)

	//用户
	QueryUsers(link sqlm.ITable, req *AsyncRequest) (*AsyncUsersReply, error)
	QueryUsersExtra(link sqlm.ITable, req *AsyncRequest) (*AsyncUsersExtraReply, error)
	// supplier
	QuerySuppliers(link sqlm.ITable, req *AsyncRequest) (*AsyncSuppliersReply, error)

	// 授权信息
	GetAuthInfo(link sqlm.ITable, mchId, apiKey string) (*AuthInfo, error)
}
