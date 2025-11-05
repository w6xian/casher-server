package store

import "context"

type CallReq struct {
	Tracker *Tracker `json:"tracker"`
}

func (c *CallReq) Validate() error {
	return nil
}

type CallResp struct {
}

func (s *Store) Call(ctx context.Context, req *CallReq) (*CallResp, error) {
	s.WsLogic.Broadcast(ctx, []byte("action:setting_query"))
	resp := &CallResp{}
	return resp, nil
}
