/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 15:19
 */
package connect

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (c *Connect) InitWebsocket(wsLogic *WsLogic) error {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		c.serveWs(wsLogic.Server, w, r)
	})
	return nil
}

func (c *Connect) serveWs(server *Server, w http.ResponseWriter, r *http.Request) {
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  server.Profile.Server.ReadBufferSize,
		WriteBufferSize: server.Profile.Server.WriteBufferSize,
	}
	fmt.Println("ws addr:", server.Profile.Server.WsAddr)
	//cross origin domain support
	upGrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upGrader.Upgrade(w, r, nil)

	if err != nil {
		c.Lager.Error("serverWs err:%s", zap.Error(err))
		return
	}
	// 一个连接一个channel
	ch := NewChannel(server.Profile.Server.BroadcastSize)
	fmt.Println("ch:", ch)
	//default broadcast size eq 512
	ch.conn = conn
	//send data to websocket conn
	go server.writePump(ch, c)
	//get data from websocket conn
	// 需要确认客户端是否合法，一个是JWT,一个是ClientID
	go server.readPump(ch, c)
}
