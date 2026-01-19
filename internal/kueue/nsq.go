package kueue

import (
	"github.com/nsqio/go-nsq"
)

type Kueue struct {
	topic   string
	addr    string
	handler nsq.Handler
}

func NewQueue(topic string, addr string) *Kueue {
	return &Kueue{
		topic: topic,
		addr:  addr,
	}
}

// channel
func (k *Kueue) Channel() string {
	return "MAIN"
}

func (k *Kueue) Start() error {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(k.topic, k.Channel(), cfg)
	consumer.SetLoggerLevel(nsq.LogLevelError)
	if err != nil {
		return err
	}
	consumer.AddHandler(k.handler)
	if err = consumer.ConnectToNSQLookupds([]string{k.addr}); err != nil {
		return err
	}
	return nil
}

func (k *Kueue) Handler(handler nsq.Handler) error {
	k.handler = handler
	return nil
}

func (k *Kueue) Tell(msg []byte) error {
	return nil
}
