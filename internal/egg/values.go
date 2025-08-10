package egg

type Value []byte

func (v Value) Uint64() uint64 {
	return ToUint64(v)
}

func (v Value) String() string {
	return string(v)
}

type KeyValue map[string]Value
type Arguments map[string]Value

func (args Arguments) K(key string) Value {
	if v, ok := args[key]; ok {
		return v
	}
	return []byte{}
}

func (args Arguments) HasCallBack() bool {
	_, ok := args["__cbk"]
	return ok
}

func (args Arguments) CallBack() Value {
	return args.K("__cbk")
}

func (args Arguments) SetCallBack(v Value) bool {
	args["__cbk"] = v
	return true
}

type Arg map[string]string
