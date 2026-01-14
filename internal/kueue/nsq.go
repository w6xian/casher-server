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
	return "sensor01"
}

func (k *Kueue) Start() error {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(k.topic, k.Channel(), cfg)
	if err != nil {
		return err
	}
	consumer.AddHandler(k.handler)
	return nil
}

func (k *Kueue) Handler(handler nsq.Handler) error {
	k.handler = handler
	return nil
}

func (k *Kueue) Tell(msg []byte) error {
	return nil
}
