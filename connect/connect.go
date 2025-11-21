/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 18:18
 */
package connect

import (
	"casher-server/internal/config"
	"casher-server/internal/queue"
	"casher-server/internal/utils/id"
	"casher-server/proto"
	"context"
	"reflect"
	"sync"
	"time"

	"fmt"
	_ "net/http/pprof"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/louis-xie-programmer/go-local-cache/cache"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"modernc.org/mathutil"
)

var once sync.Once
var WsLogicObjc *WsLogic

type WsLogic struct {
	Server *Server
}

func InitWsLogicServer() *WsLogic {
	once.Do(func() {
		WsLogicObjc = &WsLogic{}
	})
	return WsLogicObjc
}

func (c *WsLogic) Call(ctx context.Context, userId int64, mtd string, data any) ([]byte, error) {
	if c.Server == nil {
		return nil, errors.New("server not found")
	}
	b := c.Server.Bucket(userId)
	ch := b.Channel(userId)
	if ch == nil {
		return nil, errors.New("channel not found")
	}

	resp, err := ch.Call(ctx, mtd, data)
	if err != nil {
		fmt.Println("Connect layer Call() error", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (c *WsLogic) Channel(ctx context.Context, userId int64, action int, data string) {
	if c.Server == nil {
		return
	}
	b := c.Server.Bucket(userId)
	ch := b.Channel(userId)
	if ch == nil {
		return
	}
	cmd := proto.CmdReq{
		Id:     id.ShortID(),
		Ts:     time.Now().Unix(),
		Action: action,
		Data:   data,
	}
	msg := &proto.Msg{
		Body: cmd.Bytes(),
	}
	if err := ch.Push(ctx, msg); err != nil {
		fmt.Println("Connect layer Push() error", zap.Error(err))
	}
}

func (c *WsLogic) Room(ctx context.Context, roomId int64, action int, data string) {
	if c.Server == nil {
		return
	}
	room := c.Server.Room(roomId)
	if room == nil {
		return
	}
	if room.drop {
		return
	}
	cmd := proto.CmdReq{
		Id:     id.ShortID(),
		Ts:     time.Now().Unix(),
		Action: action,
		Data:   data,
	}
	msg := &proto.Msg{
		Body: cmd.Bytes(),
	}
	fmt.Println("Connect layer Push() roomId", roomId)
	room.Push(ctx, msg)
}

func (c *WsLogic) Broadcast(ctx context.Context, action int, data string) {
	if c.Server == nil {
		return
	}

	cmd := proto.CmdReq{
		Id:     id.ShortID(),
		Ts:     time.Now().Unix(),
		Action: action,
		Data:   data,
	}
	msg := &proto.Msg{
		Body: cmd.Bytes(),
	}
	if err := c.Server.Broadcast(ctx, msg); err != nil {
		return
	}
}

type Connect struct {
	ServerId     string
	Lager        *zap.Logger
	Profile      *config.Profile
	Actor        *queue.ActorPool
	cache        *cache.Cache
	serviceMapMu sync.RWMutex
	serviceMap   map[string]*serviceFns
}

// server-bucket-channel
// each bucket has a channel to send message to client

func New(ctx context.Context, profile *config.Profile, lager *zap.Logger, cache *cache.Cache, actor *queue.ActorPool) *Connect {
	svr := new(Connect)
	svr.Profile = profile
	svr.serviceMapMu = sync.RWMutex{}
	svr.serviceMap = make(map[string]*serviceFns)
	svr.Lager = lager
	svr.cache = cache
	svr.Actor = actor
	return svr
}

func (c *Connect) Server(wsLogic *WsLogic, r *mux.Router) {
	// get Connect layer config
	connectConfig := c.Profile.Connect
	//set the maximum number of CPUs that can be executing
	runtime.GOMAXPROCS(connectConfig.ConnectBucket.CpuNum)

	bsNum := mathutil.Max(connectConfig.ConnectBucket.CpuNum, 1)
	bsNum = mathutil.Min(bsNum, runtime.NumCPU())
	//init Connect layer rpc server, logic client will call this
	bs := make([]*Bucket, bsNum)
	for i := 0; i < bsNum; i++ {
		bs[i] = NewBucket(BucketOptions{
			ChannelSize:   connectConfig.ConnectBucket.Channel,
			RoomSize:      connectConfig.ConnectBucket.Room,
			RoutineAmount: connectConfig.ConnectBucket.RoutineAmount,
			RoutineSize:   connectConfig.ConnectBucket.RoutineSize,
		})
	}
	operator := NewDefaultOperator(c.cache)
	wsLogic.Server = NewServer(bs, operator, c.Profile, c.Lager, c.serviceMap)
	c.ServerId = fmt.Sprintf("%s-%s", "ws", id.ShortID())
	c.Lager.Info("Connect layer server id", zap.String("server_id", c.ServerId))
	//start Connect layer server handler persistent connection
	if err := c.InitWebsocket(wsLogic, r); err != nil {
		c.Lager.Panic("Connect layer InitWebsocket() error", zap.Error(err))
		panic("Connect layer InitWebsocket() error")
	}
}

func (c *Connect) RegisterName(name string, rcvr any, metadata string) error {
	_, err := c.register(name, rcvr)
	if err != nil {
		return err
	}
	return nil
}

// RegisterName registers the receiver object for the given name.
// methodName is the name of the method to register.
func (c *Connect) register(name string, rcvr any) (string, error) {
	c.serviceMapMu.Lock()
	defer c.serviceMapMu.Unlock()

	service := new(serviceFns)
	getType := reflect.TypeOf(rcvr)
	service.v = reflect.ValueOf(rcvr)
	k := getType.Kind()
	if k == reflect.Pointer {
		el := getType.Elem()
		sname := fmt.Sprintf("%s.%s", el.PkgPath(), el.Name())
		service.name = sname
	} else {
		sname := fmt.Sprintf("%s.%s", getType.PkgPath(), getType.Name())
		service.name = sname
	}
	// Install the methods
	service.method = suitableMethods(getType)
	fmt.Println("suitableMethods methods = ", service.name)
	c.serviceMap[name] = service
	return service.name, nil
}

// Precompute the reflect type for context.
var typeOfContext = reflect.TypeOf((*context.Context)(nil)).Elem()

// Precompute the reflect type for error.
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

func suitableMethods(typ reflect.Type) map[string]reflect.Method {
	methods := make(map[string]reflect.Method)
	for m := 0; m < typ.NumMethod(); m++ {
		m := typ.Method(m)
		// 这里可以加一些方法需要什么样的参数，比如第一个参数必须是context.Context
		if m.Type.NumIn() < 2 || m.Type.In(1) != reflect.TypeOf((*context.Context)(nil)).Elem() {
			continue
		}
		// Method must be exported.
		if m.PkgPath != "" {
			continue
		}
		if !m.IsExported() {
			continue
		}
		// 只限定第一个参数，一这是context.Context，后面的参数可以是任意类型
		if m.Type.NumIn() < 2 {
			panic(fmt.Sprintf("method %s must have at least 1 arguments", m.Name))
		}
		arg1 := m.Type.In(1)
		// 判定第一个参数是不是context.Context
		if !arg1.Implements(typeOfContext) {
			panic(fmt.Sprintf("method %s must have at least 1 arguments, first argument must be context.Context", m.Name))
		}
		// 返回值最后一个值需要是error
		if m.Type.NumOut() < 1 {
			panic(fmt.Sprintf("method %s must have 1-2 return value and last return value must be error", m.Name))
		}
		if m.Type.NumOut() > 2 {
			panic(fmt.Sprintf("method %s must have 1-2 return values and last return value must be error", m.Name))
		}
		out := m.Type.Out(m.Type.NumOut() - 1)
		if !out.Implements(typeOfError) {
			panic(fmt.Sprintf("method %s must have at least 1 return value, last return value must be error", m.Name))
		}
		methods[m.Name] = m
	}
	return methods
}

type serviceFns struct {
	name   string                    // name of service
	v      reflect.Value             // receiver of methods for the service
	method map[string]reflect.Method // registered methods
}
