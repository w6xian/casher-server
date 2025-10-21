package mysql

import (
	"casher-server/internal/store"
	"casher-server/internal/utils/id"
	"time"

	"github.com/w6xian/sqlm"
)

func (db *DB) CreateStoreLite(link sqlm.ITable, req *store.ShopInfoReq) (int64, error) {
	reqId, err := id.NextId(1)
	if err != nil {
		reqId = time.Now().UnixNano()
	}
	mchCode := id.ShortID()
	id, err := link.Table(store.TABLE_COM_SHOPS).
		Insert(map[string]any{
			"open_id":         reqId,
			"mch_code":        mchCode,
			"name":            req.Name,
			"sort_name":       req.Name,
			"contact_name":    req.ChiefName,
			"contact_phone":   req.Mobile,
			"contact_mobile":  req.Mobile,
			"categories_id":   6,
			"categories_name": "终端（零售）",
			"level_id":        1,
			"level_name":      "初始等级LV1",
			"avatar":          req.Avatar,
			"tutor_id":        1,
			"tutor_name":      "Leo",
			"manager_id":      1,
			"manager_name":    "Leo",
			"create_id":       1,
			"create_name":     "Leo",
			"area_code":       req.AreaCode,
			"area_path":       req.AreaPath,
			"area_street":     req.AreaStreet,
			"street":          req.Street,
			"address":         req.Address,
			"mark":            req.Mark,
			"theme":           "#B10DC9",
			"status":          1,
			"intime":          sqlm.UnixTime(),
		})
	if err != nil {
		return 0, err
	}

	return id, nil
}
