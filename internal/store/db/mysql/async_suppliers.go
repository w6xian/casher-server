package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) QuerySuppliers(link sqlm.ITable, req *store.AsyncRequest) (*store.AsyncSuppliersReply, error) {
	reply := &store.AsyncSuppliersReply{}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	products, err := link.Table(store.TABLE_CRM_SUPPLIERS).
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("id > %d", req.CloudId).
		OrderASC("id").
		Limit(limit).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	ps := []*store.SupplierLite{}
	err = products.Scan(&ps, func(row *sqlm.Row) any {
		return &store.SupplierLite{}
	})
	if err != nil {
		return nil, err
	}
	row, err := link.Table(store.TABLE_CRM_SUPPLIERS).
		Count().
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("id > %d", req.CloudId).
		Query()
	if err != nil {
		return nil, err
	}
	c := row.Get("total").NullInt64()
	reply.Suppliers = ps
	reply.TotalNum = c.Int64
	return reply, nil
}
