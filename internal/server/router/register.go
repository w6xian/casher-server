package router

import (
	"casher-server/internal/lager"
	"context"
	"net/http"

	v1 "casher-server/internal/server/router/api/v1"

	"github.com/gorilla/mux"
)

func Register(ctx context.Context, r *mux.Router, v *v1.Api) {
	// 收银相关
	r.HandleFunc("/", JsonV2(func(w http.ResponseWriter, req *http.Request) ([]byte, error) {
		return []byte("hello this is casher server"), nil
	})).Methods(http.MethodGet, http.MethodOptions)

	// 设置 setting
	// ### 同步操作
	r.HandleFunc("/shop-setting-table-id", JsonV2(v.Call, lager.RegLager(v.Lager, "查询配置表ID"))).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)

}
