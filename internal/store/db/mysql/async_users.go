package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) QueryUsers(link sqlm.ITable, req *store.AsyncRequest) (*store.AsyncUsersReply, error) {
	reply := &store.AsyncUsersReply{}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	users, err := link.Table(store.TABLE_COM_SHOPS_USERS).
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("shop_id=%d", req.Tracker.ShopId).
		And("id > %d", req.CloudId).
		OrderASC("id").
		Limit(limit).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	us := []*store.UserLite{}
	err = users.Scan(&us, func(row *sqlm.Row) any {
		return &store.UserLite{}
	})
	if err != nil {
		return nil, err
	}
	row, err := link.Table(store.TABLE_COM_SHOPS_USERS).
		Count().
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("shop_id=%d", req.Tracker.ShopId).
		And("id > %d", req.CloudId).
		Query()
	if err != nil {
		return nil, err
	}
	c := row.Get("total").NullInt64()
	reply.Users = us
	reply.TotalNum = c.Int64
	return reply, nil
}

func (db *DB) QueryUsersExtra(link sqlm.ITable, req *store.AsyncRequest) (*store.AsyncUsersExtraReply, error) {
	reply := &store.AsyncUsersExtraReply{}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	levels, err := link.Table(store.TABLE_COM_SHOPS_USERS_LEVELS).
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("shop_id=%d", req.Tracker.ShopId).
		And("id > %d", req.CloudId).
		OrderASC("id").
		Limit(limit).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	ls := []*store.LevelLite{}
	err = levels.Scan(&ls, func(row *sqlm.Row) any {
		return &store.LevelLite{}
	})
	if err != nil {
		return nil, err
	}

	tags, err := link.Table(store.TABLE_COM_SHOPS_USERS_TAGS).
		Where("proxy_id=%d", req.Tracker.ProxyId).
		And("shop_id=%d", req.Tracker.ShopId).
		And("id > %d", req.CloudId).
		OrderASC("id").
		Limit(limit).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	ts := []*store.TagLite{}
	err = tags.Scan(&ts, func(row *sqlm.Row) any {
		return &store.TagLite{}
	})
	if err != nil {
		return nil, err
	}

	reply.Levels = ls
	reply.Tags = ts
	return reply, nil
}
