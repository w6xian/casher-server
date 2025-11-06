package store

import "context"

type WsReq struct {
	Tracker *Tracker `json:"tracker"`
}

func (c *WsReq) Validate() error {
	return nil
}

func (s *Store) NoticeNewOrder(ctx context.Context, req *WsReq) (*CallResp, error) {
	s.WsLogic.Room(ctx, 2, ACTION_NOTICE_ORDER, "[111,125]")
	resp := &CallResp{}
	return resp, nil
}
