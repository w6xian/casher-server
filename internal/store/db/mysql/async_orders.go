package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) QueryOrders(link sqlm.ITable, proxyId int64, mobile string) (*store.SyncOrdersReply, error) {
	reply := &store.SyncOrdersReply{}
	orders, err := link.Table(store.TABLE_COM_ADMIN).
		Where("username = '%s'", mobile).
		AndOption(proxyId > 0, "proxy_id = %d", proxyId).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	os := []*store.OrderLite{}
	err = orders.Scan(&os, func(row *sqlm.Row) any {
		return &store.OrderLite{}
	})
	if err != nil {
		return nil, err
	}
	return reply, nil
}
