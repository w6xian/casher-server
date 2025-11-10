/**
 * Created by lock
 * Date: 2019-08-10
 * Time: 18:32
 */
package connect

import (
	"casher-server/internal/command"
	"casher-server/internal/config"
	"casher-server/internal/utils"
	"casher-server/proto"
	"casher-server/tools"
	"context"
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
func (s *Server) Bucket(userId int64) *Bucket {
	userIdStr := fmt.Sprintf("%d", userId)
	idx := tools.CityHash32([]byte(userIdStr), uint32(len(userIdStr))) % s.bucketIdx
	return s.Buckets[idx]
}

func (s *Server) Room(roomId int64) *Room {
	for _, b := range s.Buckets {
		if room := b.Room(roomId); room != nil {
			return room
		}
	}
	return nil
}
func (s *Server) Broadcast(ctx context.Context, msg *proto.Msg) error {
	for _, b := range s.Buckets {
		for _, room := range b.rooms {
			room.Push(ctx, msg)
		}
	}
	return nil
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
		case message, ok := <-ch.rpcCaller:
			fmt.Println("=============1")
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

			w.Write(utils.Serialize(message))
			fmt.Println("=============3")
			if err := w.Close(); err != nil {
				c.Lager.Error(" w.Close err  ", zap.String("err", err.Error()))
				return
			}
		case <-ticker.C:
			//heartbeat，if ping error will exit and close current websocket conn
			ch.conn.SetWriteDeadline(time.Now().Add(s.Profile.Server.WriteWait))
			c.Lager.Info("websocket.PingMessage ", zap.Int("type", websocket.PingMessage))
			if err := ch.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
			// fmt.Println("websocket.PingMessage success")
			// cmd := proto.CmdReq{
			// 	Id:     id.ShortID(),
			// 	Ts:     time.Now().Unix(),
			// 	Action: 0xFF,
			// 	Data:   "ping",
			// }
			// if err := ch.conn.WriteJSON(cmd); err != nil {
			// 	c.Lager.Error("WriteJSON err  ", zap.String("err", err.Error()))
			// 	return
			// }
		}
	}
}

func (s *Server) readPump(ch *Channel, c *Connect) {
	defer func() {
		c.Lager.Info("start exec disConnect ...")
		if ch.Room == nil || ch.UserId == 0 {
			c.Lager.Info("roomId and userId eq 0")
			ch.conn.Close()
			return
		}
		c.Lager.Info("exec disConnect ...")
		disConnectRequest := new(proto.DisConnectRequest)
		disConnectRequest.RoomId = ch.Room.Id
		disConnectRequest.UserId = ch.UserId
		s.Bucket(ch.UserId).DeleteChannel(ch)
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
			c.Lager.Error("message is nil")
			return
		}
		var connReq *proto.CmdReq
		if reqErr := json.Unmarshal(message, &connReq); reqErr != nil {
			c.Lager.Error("message struct ", zap.String("err", reqErr.Error()))
			return
		}
		if connReq.Action == command.ACTION_BACK {
			backObj := &proto.JsonBackObject{}
			backObj.Id = connReq.Id
			backObj.Data = connReq.Data
			ch.rpcBacker <- backObj
			continue
		}

		if ch.UserId > 0 && connReq.Action != command.ACTION_LOGIN && connReq.Action != command.ACTION_LOGOUT {
			fmt.Println("connReq:", string(message))
			// 已登录过后，可以互动消息
			s.operator.HandleMessage(ch, connReq)
			continue
		}
		fmt.Println("connReq:", string(message))
		// 拿到用用户信息
		userId, roomId, err := s.operator.Connect(connReq)
		fmt.Println("userId:", userId, roomId, "Connect err:", err)
		if err != nil {
			c.Lager.Error("s.operator.Connect error  ", zap.String("err", err.Error()))
			return
		}
		if userId == 0 {
			c.Lager.Error("Invalid AuthToken ,userId empty")
			// 登录不成功，就等着下一次登录
			continue
		}
		c.Lager.Info("websocket rpc call return userId,RoomId", zap.Int64("userId", userId), zap.Int64("RoomId", connReq.RoomId))
		if connReq.Action == command.ACTION_LOGIN {
			b := s.Bucket(userId)
			//insert into a bucket
			err = b.Put(userId, roomId, ch)
			if err != nil {
				c.Lager.Error("conn close err: ", zap.String("err", err.Error()))
				ch.conn.Close()
			}
		} else if connReq.Action == command.ACTION_LOGOUT {
			b := s.Bucket(userId)
			//insert into a bucket
			err = b.Quit(ch)
			if err != nil {
				c.Lager.Error("conn close err: ", zap.String("err", err.Error()))
				ch.conn.Close()
			}
		} else if connReq.Action == command.ACTION_INVALID {
			c.Lager.Error("Invalid Action ,Action empty")
			// 登录不成功，就等着下一次登录
			continue
		}
	}
}
