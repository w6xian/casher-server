package main

import (
	"casher-server/connect"
	"context"
	"flag"
	"fmt"
	"os/exec"
	"time"

	"github.com/kardianos/service"
)

const (
	// Time allowed to write a message to the peer.
	// writeWait = 10 * time.Second

	// // Maximum message size allowed from peer.
	// maxMessageSize = 8192

	// // Time allowed to read the next pong message from the peer.
	// pongWait = 60 * time.Second

	// // Send pings to peer with this period. Must be less than pongWait.
	// pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second
)

type Deamon struct {
	Context context.Context
	FlagSet *flag.FlagSet
	pool    []*exec.Cmd
}

func (p *Deamon) Start(s service.Service) error {
	// 进程守护
	p.run(s)
	return nil
}

func (p *Deamon) run(s service.Service) {
	connect.New(s).Run()
}

func (p *Deamon) Stop(s service.Service) error {
	for _, cmd := range p.pool {
		if cmd != nil {
			if err := cmd.Process.Kill(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
	s.Stop()
	return nil
}
