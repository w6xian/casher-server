package config

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/w6xian/sqlm"
	"go.uber.org/zap/zapcore"
)

type StdLog interface {
	Debug(s string, args ...zapcore.Field)
	Info(s string, args ...zapcore.Field)
	Warn(s string, args ...zapcore.Field)
	Error(s string, args ...zapcore.Field)
	Panic(s string, args ...zapcore.Field)
	Fatal(s string, args ...zapcore.Field)
}

var opt *Profile
var once sync.Once

type Apps struct {
	Secret   string `mapstructure:"secret"`
	AppId    string `mapstructure:"app_id"`
	AppKey   string `mapstructure:"app_key"`
	Mode     string `mapstructure:"mode"`
	Language string `mapstructure:"language"`
}

type Server struct {
	Address         string        `mapstructure:"address"`
	Network         string        `mapstructure:"network"`
	Origins         []string      `mapstructure:"origins"`
	WriteWait       time.Duration `mapstructure:"write_wait"`
	PongWait        time.Duration `mapstructure:"pong_wait"`
	PingPeriod      time.Duration `mapstructure:"ping_period"`
	MaxMessageSize  int64         `mapstructure:"max_message_size"`
	ReadBufferSize  int           `mapstructure:"read_buffer_size"`
	WriteBufferSize int           `mapstructure:"write_buffer_size"`
	BroadcastSize   int           `mapstructure:"broadcast_size"`
	RpcAddr         string        `mapstructure:"rpc_addr"`
	WsAddr          string        `mapstructure:"ws_addr"`
}

type Machine struct {
	Id   int64  `mapstructure:"id"`
	Code string `mapstructure:"code"`
}

type ConnectBase struct {
	CertPath string `mapstructure:"certPath"`
	KeyPath  string `mapstructure:"keyPath"`
}

type ConnectRpcAddressWebsockts struct {
	Address string `mapstructure:"address"`
}

type ConnectRpcAddressTcp struct {
	Address string `mapstructure:"address"`
}

type ConnectBucket struct {
	CpuNum        int    `mapstructure:"cpuNum"`
	Channel       int    `mapstructure:"channel"`
	Room          int    `mapstructure:"room"`
	SrvProto      int    `mapstructure:"svrProto"`
	RoutineAmount uint64 `mapstructure:"routineAmount"`
	RoutineSize   int    `mapstructure:"routineSize"`
}

type ConnectWebsocket struct {
	ServerId string `mapstructure:"serverId"`
	Bind     string `mapstructure:"bind"`
}

type ConnectTcp struct {
	ServerId      string `mapstructure:"serverId"`
	Bind          string `mapstructure:"bind"`
	SendBuf       int    `mapstructure:"sendbuf"`
	ReceiveBuf    int    `mapstructure:"receivebuf"`
	KeepAlive     bool   `mapstructure:"keepalive"`
	Reader        int    `mapstructure:"reader"`
	ReadBuf       int    `mapstructure:"readBuf"`
	ReadBufSize   int    `mapstructure:"readBufSize"`
	Writer        int    `mapstructure:"writer"`
	WriterBuf     int    `mapstructure:"writerBuf"`
	WriterBufSize int    `mapstructure:"writeBufSize"`
}

type ConnectConfig struct {
	ConnectBase                ConnectBase                `mapstructure:"connect-base"`
	ConnectRpcAddressWebSockts ConnectRpcAddressWebsockts `mapstructure:"connect-rpcAddress-websockts"`
	ConnectRpcAddressTcp       ConnectRpcAddressTcp       `mapstructure:"connect-rpcAddress-tcp"`
	ConnectBucket              ConnectBucket              `mapstructure:"connect-bucket"`
	ConnectWebsocket           ConnectWebsocket           `mapstructure:"connect-websocket"`
	ConnectTcp                 ConnectTcp                 `mapstructure:"connect-tcp"`
	ServerId                   string                     `mapstructure:"server_id"`
}

type Logger struct {
	FilePath    string `mapstructure:"file_path"`    // 日志文件路径
	Level       int8   `mapstructure:"level"`        // 日志级别
	MaxSize     int    `mapstructure:"max_size"`     // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups  int    `mapstructure:"max_backups"`  // 日志文件最多保存多少个备份
	MaxAge      int    `mapstructure:"max_age"`      // 文件最多保存多少天
	Compress    bool   `mapstructure:"compress"`     // 是否压缩
	ServiceName string `mapstructure:"service_name"` // 服务名
	Stdout      bool   `mapstructure:"std_out"`
	Filename    string `mapstructure:"file_name"`
	LocalTime   bool   `mapstructure:"localtime"`
	Debug       bool   `mapstructure:"debug"`
}

type Profile struct {
	Id         string       `mapstructure:"id"`
	Machine    Machine      `mapstructure:"machine"`
	Apps       *Apps        `mapstructure:"apps"`
	AppKey     string       `mapstructure:"app_key"`
	Server     *Server      `mapstructure:"server"`
	Store      *sqlm.Server `mapstructure:"store"`
	Logger     *Logger      `mapstructure:"log"`
	ShopId     int64        `mapstructure:"shop_id"`
	ProxyId    int64        `mapstructure:"proxy_id"`
	UserId     int64        `mapstructure:"user_id"`
	EmployeeId int64        `mapstructure:"employee_id"`
	log        StdLog
	Connect    ConnectConfig `mapstructure:"connect"`
}

func (opts *Profile) SetLogger(log StdLog) *Profile {
	opts.log = log
	return opt
}

func (opts *Profile) GetLogger() StdLog {
	return opts.log
}
func (opts *Profile) FromFile(f string, t string) error {
	MustLoad(f, opts, t)
	return nil
}

func NewProfile() *Profile {
	id := uuid.New()
	once.Do(func() {
		// 创建连接
		opt = &Profile{
			Id:         id.String(),
			ShopId:     0,
			ProxyId:    0,
			UserId:     0,
			EmployeeId: 0,
			Machine: Machine{
				Id:   1,
				Code: "SST-1",
			},
			Apps: &Apps{
				Secret:   "",
				AppId:    "",
				AppKey:   "",
				Mode:     "prod",
				Language: "zh",
			},
			Server: &Server{
				RpcAddr:         "0.0.0.0:8965",
				WsAddr:          "0.0.0.0:8966",
				Address:         "0.0.0.0:8888",
				Network:         "tcp",
				Origins:         []string{"*"},
				WriteWait:       10 * time.Second,
				PongWait:        60 * time.Second,
				PingPeriod:      54 * time.Second,
				MaxMessageSize:  512,
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
				BroadcastSize:   512,
			},

			Store: &sqlm.Server{
				Database:     "cloud",
				Host:         "127.0.0.1",
				Port:         3306,
				Protocol:     "mysql",
				Pretable:     "mi_",
				Charset:      "utf8mb4",
				MaxOpenConns: 64,
				MaxIdleConns: 64,
				MaxLifetime:  int(time.Second) * 60,
				DSN:          "sqlm_demo.db",
			},
			Logger: &Logger{
				ServiceName: "",   // 服务名
				FilePath:    "./", // 日志文件路径
				Filename:    "cash.log",
				Level:       -1,
				MaxSize:     500,   // 每个日志文件保存的最大尺寸 单位：M
				MaxBackups:  2024,  // 日志文件最多保存多少个备份
				MaxAge:      180,   // 文件最多保存多少天
				Compress:    false, // 是否压缩
				Stdout:      true,
				LocalTime:   true,
				Debug:       false,
			},
		}
	})
	return opt
}
