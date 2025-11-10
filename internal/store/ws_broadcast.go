package store

import (
	"casher-server/internal/command"
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
	s.WsLogic.Broadcast(ctx, command.ACTION_BROADCAST, "action:setting_query")
	resp := &BroadcastResp{}
	return resp, nil
}
