package store

import (
	"context"
)

type SystemReq struct {
	APIReq
	UserId  int64    `json:"user_id"`
	Method  string   `json:"method"`
	Data    string   `json:"data"`
	Tracker *Tracker `json:"tracker"`
}

func (c *SystemReq) Validate() error {
	return nil
}

type SystemResp struct {
	APIReq
	Data string `json:"data"`
	// 在线用户数
	Online int64 `json:"online"`
	// bucket 数
	Buckets int64 `json:"buckets"`
	// 队列数
	Queues int64 `json:"queues"`
	// room 数
	Rooms int64 `json:"rooms"`
}

func (s *Store) SystemInfo(ctx context.Context, req *SystemReq) (*SystemResp, error) {
	// 有多少人在线
	resp := &SystemResp{}
	svr := s.WsProxy.Server.Serve
	prof, err := svr.PProf(ctx)
	if err != nil {
		return nil, err
	}
	resp.Buckets = prof.Buckets
	resp.Queues = prof.Connects
	resp.Rooms = int64(len(prof.Rooms))
	return resp, nil
}
