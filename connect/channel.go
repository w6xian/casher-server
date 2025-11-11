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

func (ch *Channel) Call(ctx context.Context, mtd string, args any) ([]byte, error) {
	ch.Lock.Lock()
	defer ch.Lock.Unlock()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	msg := NewWsJsonCallObject(mtd, []byte(utils.Serialize(args)))
	// 发送调用请求
	select {
	case ch.rpcCaller <- msg:
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	// 等待调用结果
	for {
		select {
		case <-ticker.C:
			return nil, fmt.Errorf("call timeout")
		case back := <-ch.rpcBacker:
			if back.Id == msg.Id {
				return []byte(back.Data), nil
			}
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
func (ch *Channel) Reply(id string, data []byte) error {
	ch.Lock.Lock()
	defer ch.Lock.Unlock()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	msg := NewWsJsonBackObject(id, data)
	// 发送调用请求
	select {
	case <-ticker.C:
		return fmt.Errorf("reply timeout")
	case ch.rpcBacker <- msg:
	default:
	}
	return nil
}
