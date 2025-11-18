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
	// products
	r.HandleFunc("/product-info", JsonV2(v.ProductInfo, lager.RegLager(v.Lager, "商品信息"))).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	r.HandleFunc("/products-info", JsonV2(v.ProductsInfo, lager.RegLager(v.Lager, "商品信息"))).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
	// system
	r.HandleFunc("/system-info", JsonV2(v.SystemInfo, lager.RegLager(v.Lager, "系统信息"))).Methods(http.MethodGet, http.MethodPost, http.MethodOptions)
}
