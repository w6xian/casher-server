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
	Data string `json:"data"`
}

func (s *Store) SystemInfo(ctx context.Context, req *SystemReq) (*SystemResp, error) {

	resp := &SystemResp{}
	return resp, nil
}
