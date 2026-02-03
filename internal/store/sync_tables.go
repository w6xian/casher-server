package store

import (
	"casher-server/internal/lager"
	"context"
	"encoding/json"
	"fmt"
)

type SyncTableRow struct {
	Id         int64  `json:"id"`
	TableName  string `json:"name"`
	PkCol      string `json:"pk_col"`
	InTime     int64  `json:"intime"`
	PragmaData string `json:"pragma_data"`
}

func (s *Store) SyncTables(ctx context.Context, tracker *lager.Tracker) ([]*SyncTableRow, error) {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("SyncTableCreate", "SyncTableCreate", "sync")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	// 3 查询订单信息
	tables, err := db.SyncTables(link)
	if err != nil {
		log.ErrorExit("SyncTables Query err", err)
		return nil, tracker.Error("msg_sync_tables_err", err.Error())
	}

	return tables, nil
}

func (s *Store) SyncTableCreate(ctx context.Context, tracker *lager.Tracker, tableName string, proto string) (*SyncTableRow, error) {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("SyncTableCreate", "SyncTableCreate", "sync")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	// 3 查询订单信息
	res, err := db.SyncTableCreate(link, tableName, proto)
	if err != nil {
		log.ErrorExit("SyncTableCreate Query err", err)
		return nil, tracker.Error("msg_sync_table_create_err", err.Error())
	}

	return res, nil
}

type SyncDataModel struct {
	Columns []string `json:"column"`
	Data    [][]any  `json:"data"`
	Id      int64    `json:"id"`
	InTime  int64    `json:"intime"`
}

func (s *Store) SyncTableData(ctx context.Context, tracker *lager.Tracker, tableName string, lastId int64) (*SyncDataModel, error) {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("SyncTableData", "SyncTableData", "sync")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	// 3 查询订单信息
	res, err := db.SyncTableData(link, tableName, lastId)
	if err != nil {
		log.ErrorExit("SyncTableData Query err", err)
		return nil, tracker.Error("msg_sync_table_data_err", err.Error())
	}
	rst := []map[string]any{}
	unErr := json.Unmarshal([]byte(res), &rst)
	if unErr != nil {
		log.ErrorExit("SyncTableData Unmarshal err", unErr)
		return nil, tracker.Error("msg_sync_table_data_err", unErr.Error())
	}
	cls := []string{}
	for n, _ := range rst[0] {
		cls = append(cls, n)
	}
	fmt.Println(cls)
	datas := [][]any{}
	for _, row := range rst {
		data := []any{}
		for _, c := range cls {
			data = append(data, row[c])
		}
		datas = append(datas, data)
	}
	model := &SyncDataModel{
		Columns: cls,
		Data:    datas,
	}
	return model, nil
}
