package connect

import (
	"casher-server/proto"
)

type Operator interface {
	Connect(conn *proto.CmdReq) (int64, error)
	DisConnect(disConn *proto.DisConnectRequest) (err error)
	HandleMessage(ch *Channel, message []byte)
}

type DefaultOperator struct {
}

func (o *DefaultOperator) Connect(conn *proto.CmdReq) (uid int64, err error) {
	return 0, nil
}

func (o *DefaultOperator) DisConnect(disConn *proto.DisConnectRequest) (err error) {
	return nil
}

func (o *DefaultOperator) HandleMessage(ch *Channel, message []byte) {
	//

}
