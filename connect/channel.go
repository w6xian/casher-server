/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 15:18
 */
package connect

import (
	"casher-server/internal/utils"
	"casher-server/proto"
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// in fact, Channel it's a user Connect session
type Channel struct {
	Lock      sync.Mutex
	Room      *Room
	Next      *Channel
	Prev      *Channel
	broadcast chan *proto.Msg
	rpcCaller chan *JsonCallObject
	rpcBacker chan *JsonBackObject
	UserId    int64
	conn      *websocket.Conn
	connTcp   *net.TCPConn
}

func NewChannel(size int) (c *Channel) {
	c = new(Channel)
	c.Lock = sync.Mutex{}
	c.broadcast = make(chan *proto.Msg, size)
	c.rpcCaller = make(chan *JsonCallObject, 10)
	c.rpcBacker = make(chan *JsonBackObject, 10)
	c.Next = nil
	c.Prev = nil
	return
}

func (ch *Channel) Push(ctx context.Context, msg *proto.Msg) (err error) {
	select {
	case ch.broadcast <- msg:
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	return
}

func (c *Channel) Reply(id string, data []byte) (err error) {
	if c.conn == nil {
		return
	}
	msg := NewWsJsonBackObject(id, data)
	select {
	case c.rpcBacker <- msg:
	default:
	}
	return
}

func (ch *Channel) Call(ctx context.Context, mtd string, args any) ([]byte, error) {
	ch.Lock.Lock()
	defer ch.Lock.Unlock()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	msg := NewWsJsonCallObject(mtd, utils.Serialize(args))
	// 发送调用请求
	select {
	case <-ticker.C:
		return []byte{}, fmt.Errorf("call timeout")
	case ch.rpcCaller <- msg:
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	ticker.Reset(5 * time.Second)
	// 等待调用结果
	for {
		select {
		case <-ticker.C:
			return []byte{}, fmt.Errorf("reply timeout")
		case back, ok := <-ch.rpcBacker:
			if back.Id == msg.Id && ok {
				return []byte(back.Data), nil
			}
		case <-ctx.Done():
			return []byte{}, ctx.Err()
		}
	}
}
