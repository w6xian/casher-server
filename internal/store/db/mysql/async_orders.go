package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) QueryOrders(link sqlm.ITable, req *store.AsyncRequest) (*store.AsyncOrdersReply, error) {
	reply := &store.AsyncOrdersReply{}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	orders, err := link.Table(store.TABLE_COM_SHOPS_ORDERS).
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("com_id=%d", req.Tracker.ComId).
		And("shop_id=%d", req.Tracker.ShopId).
		And("id > %d", req.CloudId).
		OrderASC("id").
		Limit(limit).
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
	row, err := link.Table(store.TABLE_COM_SHOPS_ORDERS).
		Count().
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("com_id=%d", req.Tracker.ComId).
		And("shop_id=%d", req.Tracker.ShopId).
		And("id > %d", req.CloudId).
		Query()
	if err != nil {
		return nil, err
	}
	c := row.Get("total").NullInt64()
	reply.Orders = os
	reply.TotalNum = c.Int64
	return reply, nil
}
