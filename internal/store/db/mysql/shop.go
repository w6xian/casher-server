package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) GetShopByAppId(link sqlm.ITable, appId string) (*store.ShopLinkReqReply, error) {
	shop := &store.ShopLinkReqReply{}
	authc, err := link.Table(store.TABLE_COM_SHOPS).Where("app_id = '%s'", appId).Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(shop)
	if err != nil {
		return nil, err
	}
	return shop, nil
}
