package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) QueryProducts(link sqlm.ITable, req *store.AsyncRequest) (*store.AsyncProductsReply, error) {
	reply := &store.AsyncProductsReply{}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	products, err := link.Table(store.TABLE_COM_SHOPS_PRODUCTS).
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("shop_id=%d", req.Tracker.ShopId).
		And("id > %d", req.CloudId).
		OrderASC("id").
		Limit(limit).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	ps := []*store.ProductLite{}
	err = products.Scan(&ps, func(row *sqlm.Row) any {
		return &store.ProductLite{}
	})
	if err != nil {
		return nil, err
	}
	row, err := link.Table(store.TABLE_COM_SHOPS_PRODUCTS).
		Count().
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("shop_id=%d", req.Tracker.ShopId).
		And("id > %d", req.CloudId).
		Query()
	if err != nil {
		return nil, err
	}
	c := row.Get("total").NullInt64()
	reply.Products = ps
	reply.TotalNum = c.Int64
	return reply, nil
}

func (db *DB) QueryProductsExtra(link sqlm.ITable, req *store.AsyncRequest) (*store.AsyncProductsExtraReply, error) {
	reply := &store.AsyncProductsExtraReply{}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	cs, err := link.Table(store.TABLE_COM_PRODUCTS_CATEGORIES).
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("id > %d", req.CloudId).
		And("mall=%d", 1).
		And("status=%d", 1).
		OrderASC("id").
		Limit(limit).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	categories := []*store.CategoryLite{}
	err = cs.Scan(&categories, func(row *sqlm.Row) any {
		return &store.CategoryLite{}
	})
	if err != nil {
		return nil, err
	}

	tags, err := link.Table(store.TABLE_CRM_BRANDS).
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("id > %d", req.CloudId).
		OrderASC("id").
		Limit(limit).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	brands := []*store.BrandLite{}
	err = tags.Scan(&brands, func(row *sqlm.Row) any {
		return &store.BrandLite{}
	})
	if err != nil {
		return nil, err
	}

	reply.Categories = categories
	reply.Brands = brands
	return reply, nil
}
