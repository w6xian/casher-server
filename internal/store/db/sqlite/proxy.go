package sqlite

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) GetProxyInfoById(link sqlm.ITable, id int64) (*store.ProxyInfo, error) {
	proxy := &store.ProxyInfo{}
	authc, err := link.Table("cloud_companies").Where("com_id = %d", id).Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(proxy)
	if err != nil {
		return nil, err
	}
	return proxy, nil
}
