package store

import (
	"casher-server/internal/lager"
	"casher-server/internal/utils"
	"context"
	"fmt"
	"strings"
)

type IMap interface {
	Map() map[string]any
}

// oldData := map[string]any{"name": "张三", "age": 20}
// newData := map[string]any{"name": "张三", "age": 21, "gender": "男"}
// s.TableChange(ctx, "shop", newData, oldData)
// s.TableChange(ctx, "shops", newData, oldData)

func (s *Store) TableChange(ctx context.Context, tableName string, newData, oldData map[string]any) {
	logReq := lager.FromContext(ctx)
	tc := lager.CompareTableChange(newData, oldData)
	logReq.SetData(map[string]*lager.TablesChange{
		tableName: {
			TableChanges: []*lager.TableChange{tc},
		},
	})
}

func (s *Store) NewTableRecord(ctx context.Context, tableName string, pk int64) {
	logReq := lager.FromContext(ctx)
	link := s.GetLink(ctx)
	//万能钥匙
	data := s.driver.GetMapById(link, tableName, pk)
	tc := &lager.TableChange{
		ChangedFields: fmt.Sprintf("[%s]", strings.Join(utils.MapKeys(data), ",")),
		BeforeData:    "{}",
		AfterData:     utils.JsonString(data),
	}
	logReq.SetData(map[string]*lager.TablesChange{
		tableName: {
			TableChanges: []*lager.TableChange{tc},
		},
	})
}
