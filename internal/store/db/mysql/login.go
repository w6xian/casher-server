package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (d *DB) GetAuthInfo(link sqlm.ITable, mchId, apiKey string) (*store.AuthInfo, error) {
	authInfo := &store.AuthInfo{}
	authc, err := link.Table(store.TABLE_COM_SHOPS_AUTHS).
		Where("mch_id = '%s'", mchId).
		And("api_key = '%s'", apiKey).
		Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(authInfo)
	if err != nil {
		return nil, err
	}
	return authInfo, nil
}

// GetAuthInfoByIds 根据 shopId 和 userId 查询授权信息
func (d *DB) GetAuthInfoByIds(link sqlm.ITable, shopId, userId int64) (*store.AuthInfo, error) {
	authInfo := &store.AuthInfo{}
	authc, err := link.Table(store.TABLE_COM_SHOPS_AUTHS).
		Where("shop_id = %d", shopId).
		And("user_id = %d", userId).
		Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(authInfo)
	if err != nil {
		return nil, err
	}
	return authInfo, nil
}
