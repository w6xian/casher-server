package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"casher-server/internal/config"
	"casher-server/internal/exts"

	"github.com/kardianos/service"
)

var AppName = "sstWebsockServer"
var version string = "1.0.0"

func init() {
	// config下面处理config配制

}

func main() {
	ctx, cancel := exts.WithCancel(context.Background())
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)

	svcConfig := &service.Config{
		Name:        AppName,
		DisplayName: "casher-ws-service",
		Description: "services for websocket",
	}
	// 进程守护

	flagSet := flag.NewFlagSet(AppName, flag.ContinueOnError)
	profile := config.NewProfile()
	prg := &Deamon{
		Context: ctx,
		FlagSet: flagSet,
		Profile: profile,
	}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			log.Println("服务安装成功")
			return
		}

		if os.Args[1] == "uninstall" {
			s.Uninstall()
			log.Println("服务卸载成功")
			return
		}
	}

	go func() {
		<-interruptChan
		s.Stop()
		cancel()
		signal.Stop(interruptChan)
	}()

	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
