package sqlite

import (
	"casher-server/internal/store"

	"github.com/w6xian/sqlm"
)

func (db *DB) GetAuthByComId(link sqlm.ITable, comId int64) (*store.Auths, error) {
	auths := &store.Auths{}
	authc, err := link.Table("cloud_companies_shops_auth").Where("com_id = %d", comId).Query()
	if err != nil {
		return nil, err
	}
	err = authc.Scan(auths)
	if err != nil {
		return nil, err
	}
	return auths, nil
}

// 新增授权
func (db *DB) InsertAuth(link sqlm.ITable, auth *store.Auths) (int64, error) {
	id, err := link.Table("cloud_companies_shops_auth").Insert(map[string]any{
		"proxy_id":       auth.ProxyId,
		"app_id":         auth.AppId,
		"app_sec":        auth.AppSec,
		"com_id":         auth.ComId,
		"com_name":       auth.ComName,
		"auth_user_id":   auth.AuthUserId,
		"auth_user_name": auth.AuthUserName,
		"expire_time":    auth.ExpireTime,
		"status":         auth.Status,
		"intime":         auth.Intime,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
