package mysql

import (
	"casher-server/internal/store"
	"casher-server/internal/utils/id"
	"time"

	"github.com/w6xian/sqlm"
)

func (db *DB) GetCRMCompanyById(link sqlm.ITable, id int64) (*store.CompInfo, error) {
	company := &store.CompInfo{}
	authc, err := link.Table(store.TABLE_CRM_COMPANIES).Where("id = %d", id).Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}
func (db *DB) GetCRMCompanyByOpenId(link sqlm.ITable, openId string) (*store.CompInfo, error) {
	company := &store.CompInfo{}
	authc, err := link.Table(store.TABLE_CRM_COMPANIES).Where("open_id = '%s'", openId).Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (db *DB) CreateCompanyLite(link sqlm.ITable, req *store.ShopInfoReq) (int64, error) {
	reqId, err := id.NextId(1)
	if err != nil {
		reqId = time.Now().UnixNano()
	}
	id, err := link.Table(store.TABLE_CRM_COMPANIES).
		Insert(map[string]any{
			"open_id": reqId,
			"name":    req.Name,
			"status":  1,
			"intime":  sqlm.UnixTime(),
		})
	if err != nil {
		return 0, err
	}

	return id, nil
}
