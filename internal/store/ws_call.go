package store

import (
	"casher-server/internal/errors"
	"context"
	"fmt"
)

type CallReq struct {
	AppId   string   `json:"app_id"`
	UserId  int64    `json:"user_id"`
	Method  string   `json:"method"`
	Data    string   `json:"data"`
	Tracker *Tracker `json:"tracker"`
}

func (c *CallReq) Validate() error {
	return nil
}

type CallResp struct {
}

func (s *Store) Call(ctx context.Context, req *CallReq) (*CallResp, error) {
	link := s.GetLink(ctx)
	if link == nil {
		return nil, errors.New("link not found")
	}
	shop, err := s.driver.GetShopByAppId(link, req.AppId)
	if err != nil {
		return nil, err
	}

	resp, err := s.WsLogic.Call(ctx, req.UserId, "abc", map[string]string{
		"name":      "test",
		"app_id":    req.AppId,
		"shop_name": shop.Name,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Call resp=", string(resp))
	return &CallResp{}, nil
}
