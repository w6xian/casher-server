package this

import (
	"sync"

	"github.com/google/uuid"
)

var opt *Options
var once sync.Once

type Mysql struct {
	Database     string `json:"database"`
	Protocol     string `json:"protocol"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Pretable     string `json:"pretable"`
	Charset      string `json:"charset"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Maxconnetion int    `json:"maxconnection"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idel_conns"`
	MaxLifetime  int    `json:"max_life_time"`
	DNS          string `json:"dns"`
}

type Apps struct {
	Tag             []string `json:"tag"`
	Path            string   `json:"path"`
	File            bool     `json:"file"` //要不要输入到文件
	NullOrSql       string   `json:"null_or_sql"`
	TypePackage     string   `json:"type_package"`
	DefPackage      string   `json:"def_package"`
	UseTableName    bool     `json:"use_table_name"`
	UseTableColumns bool     `json:"use_table_columns"`
	KeywordPrefix   string   `json:"kw_prefix"`
	Keyword         []string `json:"keyword"`
	AppKey          string   `json:"app_key"`
	AppSecret       string   `json:"app_secret"`
	AppId           string   `json:"app_id"`
}

type Http struct {
	Scheme string `json:"scheme"`
	Root   string `json:"root"`
	Port   int    `json:"port"`
	Domain string `json:"domain"`
	Index  string `json:"index"`
	Server string `json:"server"`
}

// websocket
type Ws struct {
	Port   int    `json:"port"`
	Domain string `json:"domain"`
}

type Options struct {
	Id      string `json:"id"`
	Version string `json:"version"`
	App     Apps   `json:"app"`
	Mysql   Mysql  `json:"mysql"`
	Http    Http   `json:"http"`
	Ws      Ws     `json:"ws"`
}

func GetVersion() string {
	return "0.0.1Beta"
}

func NewOptions() *Options {
	id := uuid.New()
	once.Do(func() {
		// 创建连接
		opt = &Options{
			Id:      id.String(),
			Version: GetVersion(),
			App: Apps{
				Tag:           []string{"json"},
				KeywordPrefix: "pk_",
				AppSecret:     "",
				AppKey:        "",
				AppId:         "",
			},
			Http: Http{
				Scheme: "http",
				Root:   "dist",
				Port:   20327,
				Domain: "proxy.51d.ink",
				Server: "console.51d.ink",
				Index:  "index.html",
			},
			Ws: Ws{
				Port:   8510,
				Domain: "ws.51d.ink",
			},
			Mysql: Mysql{
				Host:     "127.0.0.1",
				Port:     3306,
				Database: "cloud",
				Protocol: "mysql",
				Username: "root",
				Password: "1Qazxsw2",
				Pretable: "mi_",
				Charset:  "utf8mb4",
			},
		}
	})
	return opt
}
