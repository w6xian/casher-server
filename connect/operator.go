package connect

import (
	"casher-server/proto"
)

type DefaultOperator struct {
}

func (o *DefaultOperator) Connect(conn *proto.CmdReq) (userId int64, roomId int64, err error) {
	userId = conn.UserId
	roomId = conn.RoomId
	return
}

func (o *DefaultOperator) DisConnect(disConn *proto.DisConnectRequest) (err error) {
	return nil
}

func (o *DefaultOperator) HandleMessage(ch *Channel, message []byte) {
	//

}
