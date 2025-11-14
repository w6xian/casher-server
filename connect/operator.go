package connect

import (
	"casher-server/proto"
	"fmt"

	"github.com/louis-xie-programmer/go-local-cache/cache"
)

type DefaultOperator struct {
	cache *cache.Cache
}

func NewDefaultOperator(cache *cache.Cache) *DefaultOperator {
	return &DefaultOperator{
		cache: cache,
	}
}

func (o *DefaultOperator) Connect(conn *proto.CmdReq) (userId int64, roomId int64, err error) {
	authInfo, b := o.cache.Get(conn.AuthToken)
	if !b {
		err = fmt.Errorf("auth token not found")
		return
	}
	auth := authInfo.(proto.IAuthInfo)
	_, roomId, userId = auth.GetUserIds()
	if userId == 0 {
		err = fmt.Errorf("invalid auth info")
		return
	}
	o.cache.Delete(conn.AuthToken)
	return
}

func (o *DefaultOperator) DisConnect(disConn *proto.DisConnectRequest) (err error) {
	return nil
}

func (o *DefaultOperator) HandleMessage(ch *Channel, message *proto.CmdReq) {
	//

}
