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
	// ### 异步操作
	r.HandleFunc("/notice-order", JsonV2(v.NewOrder, lager.RegLager(v.Lager, "创建订单"))).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	r.HandleFunc("/call", JsonV2(v.Call, lager.RegLager(v.Lager, "调用返回"))).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
}
