package v1

import (
	"casher-server/internal/muxhttp"
	"casher-server/internal/store"
	"casher-server/internal/utils"
	"fmt"
	"net/http"
)

func (v *Api) SystemInfo(w http.ResponseWriter, req *http.Request) ([]byte, error) {
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
	reqData := &store.SystemReq{}
	aErr := v.PostV(req, reqData)
	if aErr != nil {
		return nil, muxhttp.NewArgsErr(aErr)
	}
	fmt.Println(string(utils.Serialize(reqData)))
	err := store.CheckSign(reqData, reqData.AppId)
	if err != nil {
		return nil, muxhttp.NewArgsErr(err)
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
	resp, err := v.Store.SystemInfo(ctx, reqData)
	if err != nil {
		return nil, muxhttp.NewErr(err)
	}
	resp.AppId = reqData.AppId
	// 增加签名
	err = store.SetSign(resp, reqData.AppId)
	if err != nil {
		return nil, muxhttp.NewErr(err)
	}

	return muxhttp.NewRiskData(resp).ToBytes()
}
