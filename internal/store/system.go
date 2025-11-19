package store

import (
	"context"
	"fmt"
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
}

func (s *Store) SystemInfo(ctx context.Context, req *SystemReq) (*SystemResp, error) {
	// 有多少人在线
	resp := &SystemResp{}
	resp.Data = fmt.Sprintf("online: %d", 2)
	return resp, nil
}
