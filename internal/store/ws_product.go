package store

import (
	"casher-server/internal/errors"
	"casher-server/internal/utils"
	"casher-server/internal/utils/def"
	"context"
)

type CallBackReq struct {
	AppId string         `json:"app_id"`
	Data  map[string]any `json:"data"`
}

type ProductUnionId struct {
	UnionId string `json:"union_id"`
}

type ProductsUnionId struct {
	UnionId []string `json:"union_id"`
}

func (s *Store) ProductInfo(ctx context.Context, req *CallReq) (*CallResp, error) {
	link := s.GetLink(ctx)
	if link == nil {
		return nil, errors.New("link not found")
	}
	pReq := ProductUnionId{}
	if err := utils.Deserialize([]byte(req.Data), &pReq); err != nil {
		return nil, err
	}
	if pReq.UnionId == "" {
		return nil, errors.New("union_id is empty")
	}

	shop, err := s.driver.GetShopByAppId(link, req.AppId)
	if err != nil {
		return nil, err
	}

	resp, err := s.WsLogic.Call(ctx, req.UserId,
		def.GetString(req.Method, "shop.ProductInfo"),
		CallBackReq{
			AppId: shop.AppId,
			Data: map[string]any{
				"union_id": pReq.UnionId,
				"columns":  []string{"stock", "union_id"},
			},
		})
	if err != nil {
		return nil, err
	}

	return &CallResp{
		Data: string(resp),
	}, nil
}

func (s *Store) ProductsInfo(ctx context.Context, req *CallReq) (*CallResp, error) {
	link := s.GetLink(ctx)
	if link == nil {
		return nil, errors.New("link not found")
	}
	pReq := ProductsUnionId{}
	if err := utils.Deserialize([]byte(req.Data), &pReq); err != nil {
		return nil, err
	}
	if len(pReq.UnionId) == 0 {
		return nil, errors.New("union_id is empty")
	}

	shop, err := s.driver.GetShopByAppId(link, req.AppId)
	if err != nil {
		return nil, err
	}

	resp, err := s.WsLogic.Call(ctx, req.UserId,
		def.GetString(req.Method, "shop.ProductsInfo"),
		CallBackReq{
			AppId: shop.AppId,
			Data: map[string]any{
				"union_ids": pReq.UnionId,
				"columns":   []string{"*"},
			},
		})
	if err != nil {
		return nil, err
	}
	return &CallResp{
		Data: string(resp),
	}, nil
}
