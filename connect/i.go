package connect

import (
	"casher-server/proto"
)

type Operator interface {
	Connect(conn *proto.CmdReq) (userId int64, roomId int64, err error)
	DisConnect(disConn *proto.DisConnectRequest) (err error)
	HandleMessage(ch *Channel, message *proto.CmdReq)
}
