/**
 * Created by lock
 * Date: 2019-08-10
 * Time: 18:32
 */
package connect

import (
	"casher-server/internal/config"
	"casher-server/proto"
	"casher-server/tools"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Server struct {
	Buckets   []*Bucket
	Profile   *config.Profile
	bucketIdx uint32
	operator  Operator
	Lager     *zap.Logger
}

func NewServer(b []*Bucket, o Operator, profile *config.Profile, lager *zap.Logger) *Server {
	s := new(Server)
	s.Buckets = b
	s.Profile = profile
	s.bucketIdx = uint32(len(b))
	s.operator = o
	s.Lager = lager
	return s
}

// reduce lock competition, use google city hash insert to different bucket
func (s *Server) Bucket(userId int) *Bucket {
	userIdStr := fmt.Sprintf("%d", userId)
	idx := tools.CityHash32([]byte(userIdStr), uint32(len(userIdStr))) % s.bucketIdx
	return s.Buckets[idx]
}

func (s *Server) writePump(ch *Channel, c *Connect) {
	//PingPeriod default eq 54s
	ticker := time.NewTicker(s.Profile.Server.PingPeriod)
	defer func() {
		ticker.Stop()
		ch.conn.Close()
	}()

	for {
		select {
		case message, ok := <-ch.broadcast:
			//write data dead time , like http timeout , default 10s
			ch.conn.SetWriteDeadline(time.Now().Add(s.Profile.Server.WriteWait))
			if !ok {
				c.Lager.Warn("SetWriteDeadline not ok")
				ch.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := ch.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.Lager.Warn(" ch.conn.NextWriter err  ", zap.String("err", err.Error()))
				return
			}
			c.Lager.Info("message write body", zap.ByteString("body", message.Body))
			w.Write(message.Body)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			//heartbeat，if ping error will exit and close current websocket conn
			ch.conn.SetWriteDeadline(time.Now().Add(s.Profile.Server.WriteWait))
			c.Lager.Info("websocket.PingMessage ", zap.Int("type", websocket.PingMessage))
			if err := ch.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (s *Server) readPump(ch *Channel, c *Connect) {
	defer func() {
		c.Lager.Info("start exec disConnect ...")
		if ch.Room == nil || ch.userId == 0 {
			c.Lager.Info("roomId and userId eq 0")
			ch.conn.Close()
			return
		}
		c.Lager.Info("exec disConnect ...")
		disConnectRequest := new(proto.DisConnectRequest)
		disConnectRequest.RoomId = ch.Room.Id
		disConnectRequest.UserId = ch.userId
		s.Bucket(ch.userId).DeleteChannel(ch)
		if err := s.operator.DisConnect(disConnectRequest); err != nil {
			c.Lager.Warn("DisConnect err  ", zap.String("err", err.Error()))
		}
		ch.conn.Close()
	}()

	ch.conn.SetReadLimit(s.Profile.Server.MaxMessageSize)
	ch.conn.SetReadDeadline(time.Now().Add(s.Profile.Server.PongWait))
	ch.conn.SetPongHandler(func(string) error {
		ch.conn.SetReadDeadline(time.Now().Add(s.Profile.Server.PongWait))
		return nil
	})

	for {
		_, message, err := ch.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Lager.Error("readPump ReadMessage err  ", zap.String("err", err.Error()))
				return
			}
		}
		if message == nil {
			return
		}
		if ch.userId > 0 {
			// 已登录过后，可以互动消息
			s.operator.HandleMessage(ch, message)
			return
		}
		var connReq *proto.ConnectRequest
		c.Lager.Info("get a message ", zap.ByteString("body", message))
		if err := json.Unmarshal([]byte(message), &connReq); err != nil {
			c.Lager.Error("message struct ", zap.String("err", err.Error()))
		}
		if connReq == nil || connReq.AuthToken == "" {
			c.Lager.Error("s.operator.Connect no authToken")
			return
		}
		connReq.ServerId = c.ServerId //config.Conf.Connect.ConnectWebsocket.ServerId
		// 拿到用用户信息
		userId, err := s.operator.Connect(connReq)
		if err != nil {
			c.Lager.Error("s.operator.Connect error  ", zap.String("err", err.Error()))
			return
		}
		if userId == 0 {
			c.Lager.Error("Invalid AuthToken ,userId empty")
			return
		}
		c.Lager.Info("websocket rpc call return userId:%d,RoomId:%d", zap.Int("userId", userId), zap.Int("RoomId", connReq.RoomId))

		b := s.Bucket(userId)
		//insert into a bucket
		err = b.Put(userId, connReq.RoomId, ch)
		if err != nil {
			c.Lager.Error("conn close err: ", zap.String("err", err.Error()))
			ch.conn.Close()
		}
	}
}
