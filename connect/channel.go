/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 15:18
 */
package connect

import (
	"casher-server/proto"
	"context"
	"net"

	"github.com/gorilla/websocket"
)

// in fact, Channel it's a user Connect session
type Channel struct {
	Room      *Room
	Next      *Channel
	Prev      *Channel
	broadcast chan *proto.Msg
	userId    int64
	conn      *websocket.Conn
	connTcp   *net.TCPConn
}

func NewChannel(size int) (c *Channel) {
	c = new(Channel)
	c.broadcast = make(chan *proto.Msg, size)
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
