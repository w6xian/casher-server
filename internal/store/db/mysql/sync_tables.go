package mysql

import (
	"casher-server/internal/store"
	"strings"

	"github.com/w6xian/sqlm"
)

func (db *DB) SyncTables(link sqlm.ITable) ([]*store.SyncTableRow, error) {
	rows, err := link.Table("sync_tables").
		Select("id", "name", "pk_col", "intime").
		Where("id>%d", 0).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	syncRows := []*store.SyncTableRow{}
	rows.ScanMulti(&syncRows)
	return syncRows, nil
}
func (db *DB) SyncTableCreate(link sqlm.ITable, tableName string) (*store.SyncTableRow, error) {
	row, err := link.Table("sync_tables").
		Select("id", "pragma_data").
		Where("name ='%s'", tableName).
		Query()
	if err != nil {
		return nil, err
	}
	syncRows := &store.SyncTableRow{}
	row.Scan(syncRows)
	return syncRows, nil
}

func (db *DB) SyncTableUpdate(link sqlm.ITable, tableName string, lastId int64, lastTime int64) ([]byte, error) {
	rows, err := link.Table(tableName).
		Where("id <= %d", lastId).
		And("uptime > %d", lastTime).
		OrderASC("id").
		Limit(100).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	return []byte(rows.ToString()), nil
}

// SyncQueryTable 查询表数据
func (db *DB) SyncTableData(link sqlm.ITable, tableName string, lastId int64) ([]byte, error) {
	rows, err := link.Table(strings.TrimPrefix(tableName, "mi_")).
		Where("id > %d", lastId).
		OrderASC("id").
		Limit(100).
		QueryMulti()
	if err != nil {
		return nil, err
	}
	return []byte(rows.ToString()), nil
}
