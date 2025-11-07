package store

import (
	"casher-server/internal/command"
	"casher-server/internal/errors"
	"context"
	"encoding/json"
	"fmt"
)

type WsReq struct {
	AppId   string   `json:"app_id"`
	UserId  int64    `json:"user_id"`
	Tracker *Tracker `json:"tracker"`
}

func (c *WsReq) Validate() error {
	return nil
}

type WsCommand struct {
	OpenId string `json:"open_id"`
	ShopId int64  `json:"shop_id"`
	UserId int64  `json:"user_id"`
}

func (c *WsCommand) String() string {
	str, err := json.Marshal(c)
	if err != nil {
		return "{}"
	}
	return string(str)
}

// NoticeNewOrder 通知新订单
func (s *Store) NoticeNewOrder(ctx context.Context, req *WsReq) (*CallResp, error) {
	link := s.GetLink(ctx)
	if link == nil {
		return nil, errors.New("link not found")
	}
	shop, err := s.driver.GetShopByAppId(link, req.AppId)
	if err != nil {
		return nil, err
	}
	cmd := &WsCommand{
		UserId: req.UserId,
		ShopId: shop.Id,
		OpenId: shop.AppId,
	}
	fmt.Println("--------------------------------")
	if req.UserId > 0 {
		s.WsLogic.Channel(ctx, req.UserId, command.ACTION_NOTICE_ORDER, cmd.String())
	} else {
		s.WsLogic.Room(ctx, shop.Id, command.ACTION_NOTICE_ORDER, cmd.String())
	}
	resp := &CallResp{}
	return resp, nil
}
