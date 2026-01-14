package message

type Message struct {
	// 消息类型
	Type string
	// 消息内容
	Data []byte
}

type IMessager interface {
	// 发送消息
	Tell(msg []byte) error
}
