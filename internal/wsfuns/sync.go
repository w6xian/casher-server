package wsfuns

import (
	"context"
	"fmt"
	"time"

	"github.com/w6xian/sloth"
	"github.com/w6xian/sloth/message"
	"github.com/w6xian/tlv"
)

func (s *WsServerApi) SyncTables(ctx context.Context, req []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	ctx, close := s.Start(ctx)
	defer close()

	header := ctx.Value(sloth.HeaderKey).(message.Header)
	if header == nil {
		return nil, fmt.Errorf("header is nil")
	}
	// 校验 sign
	err := s.checkSign(header)
	if err != nil {
		return nil, err
	}
	tracker, err := s.TrackerFromHeader(ctx, header)
	if err != nil {
		return nil, err
	}
	resp, err := s.Store.SyncTables(ctx, tracker)
	if err != nil {
		return nil, err
	}
	return tlv.JsonEnpack(resp)
}

func (s *WsServerApi) SyncTableCreate(ctx context.Context, tableName string, proto string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	ctx, close := s.Start(ctx)
	defer close()

	header := ctx.Value(sloth.HeaderKey).(message.Header)
	if header == nil {
		return nil, fmt.Errorf("header is nil")
	}
	// 校验 sign
	err := s.checkSign(header)
	if err != nil {
		return nil, err
	}
	tracker, err := s.TrackerFromHeader(ctx, header)
	if err != nil {
		return nil, err
	}
	resp, err := s.Store.SyncTableCreate(ctx, tracker, tableName, proto)
	if err != nil {
		return nil, err
	}
	return []byte(resp.PragmaData), nil
}

func (s *WsServerApi) SyncTableData(ctx context.Context, tableName string, lastId int64, lastTime int64) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	ctx, close := s.Start(ctx)
	defer close()

	header := ctx.Value(sloth.HeaderKey).(message.Header)
	if header == nil {
		return nil, fmt.Errorf("header is nil")
	}
	// 校验 sign
	err := s.checkSign(header)
	if err != nil {
		return nil, err
	}
	tracker, err := s.TrackerFromHeader(ctx, header)
	if err != nil {
		return nil, err
	}
	resp, err := s.Store.SyncTableData(ctx, tracker, tableName, lastId)
	if err != nil {
		return nil, err
	}
	return tlv.JsonEnpack(resp)
}
