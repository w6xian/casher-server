package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

// GetPublicProductBySn 根据商品编码查询商品
func (db *DB) InsertPublicProduct(link sqlm.ITable, prd *store.ProductModel) (int64, error) {
	id, err := link.Table(store.TABLE_CLOUD_PUBLIC_PRODUCTS).
		Insert(prd.ToSqlMap())
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetPublicProductBySn 根据商品编码查询商品
func (db *DB) InsertPublicProductVersion(link sqlm.ITable, prd *store.ProductModel) (int64, error) {
	id, err := link.Table(store.TABLE_CLOUD_PUBLIC_PRODUCTS_VERSIONS).
		Insert(prd.ToSqlMap())
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetPublicProductBySn 根据商品编码查询商品
func (db *DB) GetPublicProductBySn(link sqlm.ITable, sn string) (*store.ProductModel, error) {
	product := &store.ProductModel{}
	authc, err := link.Table(store.TABLE_CLOUD_PUBLIC_PRODUCTS).Where("sn = '%s'", sn).Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetPublicProductById 根据商品ID查询商品
func (db *DB) GetPublicProductById(link sqlm.ITable, id int64) (*store.ProductModel, error) {
	product := &store.ProductModel{}
	authc, err := link.Table(store.TABLE_CLOUD_PUBLIC_PRODUCTS).Where("id = %d", id).Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// QueryPublicProductBySn 根据商品编码查询商品
func (db *DB) QueryPublicProductBySn(link sqlm.ITable, sn string) ([]*store.ProductLiteModel, error) {
	prds, err := link.Table(store.TABLE_CLOUD_PUBLIC_PRODUCTS).
		Where("sn like '%s%%'", sn).
		Limit(10).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	products := []*store.ProductLiteModel{}
	err = prds.Scan(&products, func(row *sqlm.Row) any {
		return &store.ProductLiteModel{}
	})
	if err != nil {
		return nil, err
	}
	return products, nil
}

// QueryPublicProductByName 根据商品名称查询商品
func (db *DB) QueryPublicProductByName(link sqlm.ITable, name string) ([]*store.ProductLiteModel, error) {
	prds, err := link.Table(store.TABLE_CLOUD_PUBLIC_PRODUCTS).
		Where("name like '%s%%'", name).
		Limit(10).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	products := []*store.ProductLiteModel{}
	err = prds.Scan(&products, func(row *sqlm.Row) any {
		return &store.ProductLiteModel{}
	})
	if err != nil {
		return nil, err
	}
	return products, nil
}
