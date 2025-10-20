package sqlite

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) GetCRMCompanyById(link sqlm.ITable, id int64) (*store.CompInfo, error) {
	company := &store.CompInfo{}
	authc, err := link.Table("crm_companies").Where("id = %d", id).Query()
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
	authc, err := link.Table("crm_companies").Where("open_id = '%s'", openId).Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}
