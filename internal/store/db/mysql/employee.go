package mysql

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) CreateEmployeeLite(link sqlm.ITable, req *store.ShopInfoReq) (int64, error) {

	id, err := link.Table(store.TABLE_COM_EMPLOYEES).
		Insert(map[string]any{

			"name":   req.Name,
			"status": 1,
			"intime": sqlm.UnixTime(),
		})
	if err != nil {
		return 0, err
	}

	return id, nil
}
