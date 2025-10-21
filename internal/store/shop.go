package store

import (
	"casher-server/internal/lager"
	"context"
	"fmt"
)

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
