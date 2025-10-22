package main

import (
	"casher-server/api/rpc"
	"casher-server/connect"
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"casher-server/internal/config"
	"casher-server/internal/i18n"
	"casher-server/internal/store"
	"casher-server/internal/store/db"

	"github.com/kardianos/service"
	"github.com/soheilhy/cmux"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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
	Profile *config.Profile
}

func (p *Deamon) Start(s service.Service) error {
	// 进程守护
	p.run(s)
	return nil
}

func (p *Deamon) run(s service.Service) {
	ctx := p.Context
	// 初始化日志
	p.initConfig()
	// 初始化 i18n
	p.initLanguage()

	// 日志
	optsLog := p.Profile.Logger
	hook := &lumberjack.Logger{
		Filename:   optsLog.FilePath + optsLog.Filename,
		MaxSize:    optsLog.MaxSize, // megabytes
		MaxBackups: optsLog.MaxBackups,
		MaxAge:     optsLog.MaxAge,   //days
		Compress:   optsLog.Compress, // disabled by default
		LocalTime:  optsLog.LocalTime,
	}
	defer hook.Close()

	encoderConfig := ecszap.EncoderConfig{
		EncodeName:     zapcore.FullNameEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   ecszap.FullCallerEncoder,
	}

	syncer := []zapcore.WriteSyncer{
		zapcore.AddSync(hook),
	}
	if optsLog.Stdout {
		syncer = append(syncer, zapcore.AddSync(os.Stdout))
	}

	core := ecszap.NewCore(encoderConfig,
		// zapcore.AddSync(hook),
		zapcore.NewMultiWriteSyncer(syncer...),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.Level(p.Profile.Logger.Level)
		}),
	)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	// 数据库 这样设计是为了方便以后换数据库引擎
	dbDriver, err := db.NewDBDriver(p.Profile)
	if err != nil {
		panic(err)
	}
	storeInstance, err := store.New(dbDriver, p.Profile, logger)
	if err != nil {
		panic(err)
	}
	if err = storeInstance.Migrate(ctx); err != nil {
		panic(err)
	}

	ln, err := net.Listen("tcp", p.Profile.Server.WsAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("server ws addr:", p.Profile.Server.WsAddr)
	muxServer := cmux.New(ln)
	//Otherwise, we match it againts a websocket upgrade request.
	wsListener := muxServer.Match(cmux.HTTP1HeaderField("Upgrade", "websocket"))
	// wsl := m.Match(cmux.HTTP1HeaderField("Upgrade", "websocket"))
	// httpListener := muxServer.Match(cmux.HTTP1Fast())
	// rpcxListener := muxServer.Match(cmux.Any())
	go func() {
		//初始化加入对应的
		connect.New(p.Context, p.Profile, logger).Server()
		http.Serve(wsListener, nil)
	}()
	go func() {
		rpc.InitLogicRpcServer(p.Context, p.Profile, logger, storeInstance)
	}()

	if err := muxServer.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		panic(err)
	}

}

// func rpcxPrefixByteMatcher() cmux.Matcher {
// 	var magic = byte(0x08)
// 	return func(r io.Reader) bool {
// 		buf := make([]byte, 1)
// 		n, _ := r.Read(buf)
// 		fmt.Println(buf)
// 		return n == 1 && buf[0] == magic
// 	}
// }

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

func (h *Deamon) initConfig() {
	h.FlagSet.String("config", "conf.toml", "path to config file")
	configFile := h.FlagSet.Lookup("config").Value.String()
	fmt.Println("configFile=", configFile)
	// 文件里读取配置
	err := h.Profile.FromFile(configFile, config.TOML)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func (h *Deamon) initLanguage() {
	// 初始化 i18n
	h.FlagSet.String("lang", "locales", "path to language files")
	langDir := h.FlagSet.Lookup("lang").Value.String()
	err := i18n.Init(langDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
