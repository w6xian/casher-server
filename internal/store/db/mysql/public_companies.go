package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

// GetPublicCompanyBySn 根据公司编码查询公司
func (db *DB) GetPublicCompanyBySn(link sqlm.ITable, sn string) (*store.CompanyModel, error) {
	company := &store.CompanyModel{}
	authc, err := link.Table(store.TABLE_CLOUD_PUBLIC_COMPANIES).
		Where("sn = '%s'", sn).
		Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

// GetPublicCompanyByName 根据公司名称查询公司
func (db *DB) GetPublicCompanyByName(link sqlm.ITable, name string) (*store.CompanyModel, error) {
	company := &store.CompanyModel{}
	authc, err := link.Table(store.TABLE_CLOUD_PUBLIC_COMPANIES).
		Where("name = '%s'", name).
		Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

// GetPublicCompanyById 根据公司ID查询公司
func (db *DB) GetPublicCompanyById(link sqlm.ITable, id int64) (*store.CompanyModel, error) {
	company := &store.CompanyModel{}
	authc, err := link.Table(store.TABLE_CLOUD_PUBLIC_COMPANIES).
		Where("id = %d", id).
		Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

// QueryPublicCompanyBySn 根据公司编码查询公司
func (db *DB) QueryPublicCompanyBySn(link sqlm.ITable, sn string) ([]*store.CompanyLiteModel, error) {
	prds, err := link.Table(store.TABLE_CLOUD_PUBLIC_COMPANIES).
		Where("sn like '%s%%'", sn).
		Limit(10).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	companies := []*store.CompanyLiteModel{}
	err = prds.Scan(&companies, func(row *sqlm.Row) any {
		return &store.CompanyLiteModel{}
	})
	if err != nil {
		return nil, err
	}
	return companies, nil
}

// QueryPublicCompanyByName 根据公司名称查询公司
func (db *DB) QueryPublicCompanyByName(link sqlm.ITable, name string) ([]*store.CompanyLiteModel, error) {
	prds, err := link.Table(store.TABLE_CLOUD_PUBLIC_COMPANIES).
		Where("name like '%s%%'", name).
		Limit(10).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	companies := []*store.CompanyLiteModel{}
	err = prds.Scan(&companies, func(row *sqlm.Row) any {
		return &store.CompanyLiteModel{}
	})
	if err != nil {
		return nil, err
	}
	return companies, nil
}
