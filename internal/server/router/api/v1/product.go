package v1

import (
	"casher-server/internal/muxhttp"
	"casher-server/internal/store"
	"fmt"
	"net/http"
)

func (v *Api) ProductInfo(w http.ResponseWriter, req *http.Request) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "%s", err)
		}
	}()
	// API 上下文开始
	req, stop := v.Start(req)
	defer stop()

	ctx, close := v.DbConnectWithClose(req.Context())
	defer close()
	reqData := &store.CallReq{}
	aErr := v.PostV(req, reqData)
	if aErr != nil {
		return nil, muxhttp.NewArgsErr(aErr)
	}
	// // 校验参数
	tracker, tErr := v.GetTracker(req, true)
	if tErr != nil {
		return nil, tErr
	}
	reqData.Tracker = tracker
	vErr := reqData.Validate()
	if vErr != nil {
		return nil, muxhttp.NewArgsValidErr(vErr)
	}
	if reqData.AppId == "" {
		reqData.AppId = "610800923266441381"
		reqData.UserId = 10
		reqData.Data = `{"union_id":"MDraSyLNC6faYf71eGT6sG"}`
	}
	resp, err := v.Store.ProductInfo(ctx, reqData)
	if err != nil {
		return nil, muxhttp.NewErr(err)
	}

	return muxhttp.NewRiskData(resp.Data).ToBytes()
}

func (v *Api) ProductsInfo(w http.ResponseWriter, req *http.Request) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "%s", err)
		}
	}()
	// API 上下文开始
	req, stop := v.Start(req)
	defer stop()

	ctx, close := v.DbConnectWithClose(req.Context())
	defer close()
	reqData := &store.CallReq{}
	aErr := v.PostV(req, reqData)
	if aErr != nil {
		return nil, muxhttp.NewArgsErr(aErr)
	}
	// // 校验参数
	tracker, tErr := v.GetTracker(req, true)
	if tErr != nil {
		return nil, tErr
	}
	reqData.Tracker = tracker
	vErr := reqData.Validate()
	if vErr != nil {
		return nil, muxhttp.NewArgsValidErr(vErr)
	}
	if reqData.AppId == "" {
		reqData.AppId = "610800923266441381"
		reqData.UserId = 10
		reqData.Data = `{"union_id":["MDraSyLNC6faYf71eGT6sG","PqWexQmyVXm4fkps3ZRHZF"]}`
	}
	resp, err := v.Store.ProductsInfo(ctx, reqData)
	if err != nil {
		return nil, muxhttp.NewErr(err)
	}

	return muxhttp.NewRiskData(resp.Data).ToBytes()
}
