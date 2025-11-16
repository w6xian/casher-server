package mapv

type mapv map[string]any

func NewMapv(values map[string]any) mapv {
	return mapv(values)
}

func (v mapv) Int64(key string) int64 {
	return v[key].(int64)
}

// 基本类型转换
// int32 int uint uint32 string ..
func (v mapv) Int32(key string) int32 {
	return v[key].(int32)
}

func (v mapv) Int(key string) int {
	return v[key].(int)
}

func (v mapv) Uint(key string) uint {
	return v[key].(uint)
}

func (v mapv) Uint32(key string) uint32 {
	return v[key].(uint32)
}

func (v mapv) Uint64(key string) uint64 {
	return v[key].(uint64)
}

func (v mapv) String(key string) string {
	return v[key].(string)
}

func (v mapv) Byte(key string) byte {
	return v[key].(byte)
}

func (v mapv) Bytes(key string) []byte {
	return v[key].([]byte)
}
