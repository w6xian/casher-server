package store

import (
	"casher-server/internal/i18n"
	"casher-server/internal/lager"
	"casher-server/internal/utils"
	"context"
	"fmt"
	"strings"
)

type Req struct {
	Lang    string         `json:"lang"`
	AppId   string         `json:"app_id"`
	TrackId string         `json:"track_id"`
	Tracker *lager.Tracker `json:"-"`
}

func (req *Req) GetTrackInfo(ctx context.Context) (string, string, string) {
	return req.AppId, req.TrackId, req.Lang
}

type ShopInfoReq struct {
	Avatar     string `json:"avatar"`
	Name       string `json:"name"`
	AreaCode   string `json:"area_code"`
	AreaPath   string `json:"area_path"`
	AreaStreet string `json:"area_street"`
	Address    string `json:"address"`
	Street     string `json:"street"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	GeoHash    string `json:"geo_hash"`
	Chief      int64  `json:"chief"`
	ChiefName  string `json:"chief_name"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	Type       int64  `json:"type"`
	Mark       string `json:"mark"`
}

type ShopInfoReqReply struct {
	Id      int64  `json:"id"`
	ProxyId int64  `json:"proxy_id"`
	OpenId  string `json:"open_id"`
	ComId   int64  `json:"com_id"`
	StoreId int64  `json:"store_id"`
}

type ShopLinkReq struct {
	Req
	Code     string `json:"code"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Sign     string `json:"sign"`
}

type ShopLinkReqReply struct {
	Id         int64  `json:"id"`
	ProxyId    int64  `json:"proxy_id"`
	OpenId     string `json:"open_id"`
	AppId      string `json:"app_id"`
	ComId      int64  `json:"com_id"`
	StoreId    int64  `json:"store_id"`
	Avatar     string `json:"avatar"`
	Sn         string `json:"sn"`
	Name       string `json:"name"`
	AreaCode   string `json:"area_code"`
	AreaPath   string `json:"area_path"`
	AreaStreet string `json:"area_street"`
	Address    string `json:"address"`
	Longitude  string `json:"longitude"`
	Latitude   string `json:"latitude"`
	GeoHash    string `json:"geo_hash"`
	Chief      int64  `json:"chief"`
	ChiefName  string `json:"chief_name"`
	Mobile     string `json:"mobile"`
	Mark       string `json:"mark"`
	Status     int64  `json:"status"`
}

func (s *Store) SyncShopInfo(ctx context.Context, req *ShopInfoReq, reply *ShopInfoReqReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("SyncShopInfo", "SyncShopInfo", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	fmt.Println("SyncShopInfo link=", link)
	// 2.1 数据驱动
	db := s.GetDriver()
	fmt.Print(db)
	// 2.2 创建公司信息
	proxy, err := db.CreateProxyLite(link, req)
	if err != nil {
		return err
	}
	fmt.Println("SyncShopInfo proxy=", proxy)
	// 创建客户
	// 创建仓库
	return nil
}

func (s *Store) ShopLink(ctx context.Context, req *ShopLinkReq, reply *ShopLinkReqReply) error {
	lang := req.Tracker
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("ShopLink", "ShopLink", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	fmt.Println("ShopLink link=", link)
	// 2.1 数据驱动
	db := s.GetDriver()
	fmt.Print(db)
	// 2.2 通过appId获取店铺信息
	shop, err := db.GetShopByAppId(link, req.Code)
	if err != nil {
		return err
	}
	fmt.Println("ShopLink shop=", shop)
	// 登录信息
	proxyId := shop.ProxyId
	admin, err := db.GetComAdmin(link, proxyId, req.Mobile)
	if err != nil {
		return lang.Error("shop_link_get_admin", "获取店铺管理员失败:{{.error}}", i18n.String("error", err.Error()))
	}
	if admin.FailTimes >= 3 {
		return lang.Error("shop_link_fail_times", "店铺链接失败次数超过次限制")
	}
	// 对比密码
	// 请用PHP中的password_verify函数对比密码，给出Golang的实现
	fmt.Println("ShopLink req.Password=", req.Password)
	fmt.Println("ShopLink admin.Password=", admin.Password)
	fmt.Println("ShopLink utils.MD5(utils.MD5(req.Password))=", utils.MD5(utils.MD5(strings.TrimSpace(req.Password))))
	fmt.Println("ShopLink utils.MD5(utils.MD5(strings.TrimSpace(req.Password)))=", utils.VerifyPassword(utils.MD5(utils.MD5(strings.TrimSpace(req.Password))), admin.Password))
	if !utils.VerifyPassword(utils.MD5(utils.MD5(strings.TrimSpace(req.Password))), admin.Password) {
		return lang.Error("shop_link_password", "店铺链接密码错误")
	}
	// 登录成功
	reply.ProxyId = proxyId
	reply.OpenId = shop.OpenId
	reply.AppId = shop.AppId
	reply.ComId = shop.ComId
	reply.StoreId = shop.StoreId
	reply.Avatar = shop.Avatar
	reply.Sn = shop.Sn
	reply.Name = shop.Name
	reply.AreaCode = shop.AreaCode
	reply.AreaPath = shop.AreaPath
	reply.AreaStreet = shop.AreaStreet
	reply.Address = shop.Address
	reply.Longitude = shop.Longitude
	reply.Latitude = shop.Latitude
	reply.GeoHash = shop.GeoHash
	reply.Chief = shop.Chief
	reply.ChiefName = shop.ChiefName
	reply.Mobile = shop.Mobile
	reply.Mark = shop.Mark
	reply.Status = shop.Status
	return nil
}
