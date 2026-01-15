package store

import (
	"casher-server/internal/action"
	"context"
)

type BroadcastReq struct {
	Tracker *Tracker `json:"tracker"`
}

func (c *BroadcastReq) Validate() error {
	return nil
}

type BroadcastResp struct {
}

func (s *Store) Broadcast(ctx context.Context, req *BroadcastReq) (*BroadcastResp, error) {
	s.WsProxy.Broadcast(ctx, action.BROADCAST, "action:setting_query")
	resp := &BroadcastResp{}
	return resp, nil
}
