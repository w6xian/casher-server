/**
 * Created by lock
 * Date: 2019-08-10
 * Time: 18:00
 */
package proto

type Msg struct {
	Ver       int    `json:"ver"`  // protocol version
	Operation int    `json:"op"`   // operation for request
	SeqId     string `json:"seq"`  // sequence number chosen by client
	Body      []byte `json:"body"` // binary body bytes
}

type JsonCallMsg struct {
	Id     string // user id
	Method string // service method name
	Args   any    // binary body bytes
	Reply  any    // binary body bytes
}
type JsonCallObject struct {
	Id     string `json:"id"`     // user id
	Action int    `json:"action"` // operation for request
	Method string `json:"method"` // service method name
	Data   string `json:"data"`   // binary body bytes
}

type JsonBackObject struct {
	Id   string `json:"id"`   // user id
	Data string `json:"data"` // binary body bytes
}

type PushMsgRequest struct {
	UserId int
	Msg    Msg
}

type PushRoomMsgRequest struct {
	RoomId int64
	Msg    Msg
}

type PushRoomCountRequest struct {
	RoomId int64
	Count  int
}
