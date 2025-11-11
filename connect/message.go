package connect

import (
	"casher-server/internal/command"
	"casher-server/internal/utils/id"
)

type JsonCallObject struct {
	Id     string `json:"id"`     // user id
	Action int    `json:"action"` // operation for request
	Method string `json:"method"` // service method name
	Data   string `json:"data"`   // binary body bytes
}

func NewWsJsonCallObject(method string, data []byte) *JsonCallObject {
	return &JsonCallObject{
		Id:     id.ShortID(),
		Action: command.ACTION_CALL,
		Method: method,
		Data:   string(data),
	}
}

type JsonBackObject struct {
	Id     string `json:"id"` // user id
	Action int64  `json:"action"`
	Data   string `json:"data"` // binary body bytes
}

func NewWsJsonBackObject(id string, data []byte) *JsonBackObject {
	return &JsonBackObject{
		Id:     id,
		Action: command.ACTION_REPLY,
		Data:   string(data),
	}
}
