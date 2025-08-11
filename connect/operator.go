/**
 * Created by lock
 * Date: 2019-08-10
 * Time: 18:35
 */
package connect

import (
	"casher-server/proto"
)

type Operator interface {
	Connect(conn *proto.ConnectRequest) (int, error)
	DisConnect(disConn *proto.DisConnectRequest) (err error)
	HandleMessage(ch *Channel, message []byte)
}

type DefaultOperator struct {
}

func (o *DefaultOperator) Connect(conn *proto.ConnectRequest) (uid int, err error) {
	return 1024, nil
}

func (o *DefaultOperator) DisConnect(disConn *proto.DisConnectRequest) (err error) {
	return nil
}

func (o *DefaultOperator) HandleMessage(ch *Channel, message []byte) {
	// 同步库存

}
