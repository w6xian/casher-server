package egg

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

func FromString(s string) *Egg {
	v := []byte(s)
	return &Egg{Type: EGG_TYPE_STRING, Len: len(v), Ai: 0, Value: v}
}

func FromUint64(i uint64) *Egg {
	v := Uint64(i)
	return &Egg{Type: EGG_TYPE_UINT64, Len: len(v), Ai: 0, Value: v}
}

func FromBytes(v []byte) *Egg {
	return &Egg{Type: EGG_TYPE_BIN, Len: len(v), Ai: 0, Value: v}
}

// 用户自己处理对应的key,对应的Value需要什么取什么,key-map格式
func FromEventMap(event string, data Arguments) *Egg {
	var v []byte
	evt := FromString(event)
	v = append(v, evt.Encode()...)
	buf := FromMapArguments(data)
	v = append(v, buf.Encode()...)
	tp := EGG_TYPE_EVENT_MAP
	if strings.HasPrefix(event, ":") {
		tp = EGG_TYPE_EVENT_BACK_MAP
	}
	return &Egg{Type: tp, Len: len(v), Ai: 0, Value: v}
}

func FromEventList(event string, eggs ...*Egg) *Egg {
	var v []byte
	evt := FromString(event)
	v = append(v, evt.Encode()...)
	for _, e := range eggs {
		v = append(v, e.Encode()...)
	}
	return &Egg{Type: EGG_TYPE_EVENT_LIST, Len: len(v), Ai: 0, Value: v}
}

func NewPing() *Egg {
	v := []byte{EGG_TYPE_PING}
	return &Egg{Type: EGG_TYPE_PING, Len: len(v), Ai: 0, Value: v}
}

func NewPong() *Egg {
	v := []byte{EGG_TYPE_PONG}
	return &Egg{Type: EGG_TYPE_PONG, Len: len(v), Ai: 0, Value: v}
}

func NewSession(s string) *Egg {
	v := []byte(s)
	return &Egg{Type: EGG_TYPE_SESSION, Len: len(v), Ai: 0, Value: v}
}

func FromTypeValues(kvs map[uint8]string) *Egg {
	var v []byte
	for key, value := range kvs {
		data := []byte(value)
		rst := make([]byte, 0, len(data)+TLAV_HEAD_LEN)
		rst = append(rst, []byte(TLAV)...)
		rst = append(rst, uint8ToBytes(key)...)
		rst = append(rst, uint24ToBytes(uint32(len(data)))...)
		rst = append(rst, uint8ToBytes(0)...)
		rst = append(rst, data...)
		v = append(v, rst...)
	}
	return &Egg{Type: EGG_TYPE_VALUES, Len: len(v), Ai: 0, Value: v}
}

func FromJson(values map[string]string, et ...uint8) *Egg {
	t := EGG_TYPE_JSON
	if len(et) > 0 {
		t = et[0]
	}
	kvs := EnJsonArgs(values)
	var v []byte
	for idx, value := range kvs {
		data := []byte(value)
		rst := make([]byte, 0, len(data)+TLAV_HEAD_LEN)
		rst = append(rst, []byte(TLAV)...)
		rst = append(rst, uint8ToBytes(uint8(EGG_TYPE_STRING))...)
		rst = append(rst, uint24ToBytes(uint32(len(data)))...)
		rst = append(rst, uint8ToBytes(uint8(idx))...)
		rst = append(rst, data...)
		v = append(v, rst...)
	}
	return &Egg{Type: t, Len: len(v), Ai: 0, Value: v}
}

func FromMapArguments(values Arguments) *Egg {
	t := EGG_TYPE_EVENT_ARGUMENTS
	kvs := EnMapArgments(values)
	var v []byte
	for key, value := range kvs {
		data := []byte(value)
		rst := make([]byte, 0, len(data)+TLAV_HEAD_LEN)
		rst = append(rst, []byte(TLAV)...)
		rst = append(rst, uint8ToBytes(EGG_TYPE_BIN)...)
		rst = append(rst, uint24ToBytes(uint32(len(data)))...)
		rst = append(rst, uint8ToBytes(uint8(key))...)
		rst = append(rst, data...)
		v = append(v, rst...)
	}
	return &Egg{Type: t, Len: len(v), Ai: 0, Value: v}
}

func FromValue(v []byte, t uint8) *Egg {
	return &Egg{Type: t, Len: len(v), Ai: 0, Value: v}
}

func FromType(t uint8, v []byte) *Egg {
	return &Egg{Type: t, Len: len(v), Ai: 0, Value: v}
}

func ReadEgg(r io.Reader) (*Egg, error) {
	pb, err := readMulti(r, 4)
	if err != nil {
		return nil, err
	}
	if string(pb) != TLAV {
		return nil, errors.New("TLAV 头信息不匹配")
	}
	t, err := readMulti(r, 1)
	if err != nil {
		return nil, err
	}

	ls, err := readMulti(r, 3)
	if err != nil {
		return nil, err
	}
	l := bytesToUint24(ls)

	ai, err := readMulti(r, 1)
	if err != nil {
		return nil, err
	}

	v, err := readMulti(r, l)
	if err != nil {
		return nil, err
	}
	return &Egg{Type: t[0], Len: l, Ai: ai[0], Value: v}, nil
}

func ReadEggWithType(r io.Reader, lt ...uint8) (*Egg, error) {
	pb, err := readMulti(r, 4)
	if err != nil {
		return nil, err
	}
	if string(pb) != TLAV {
		return nil, errors.New("TLAV 头信息不匹配")
	}
	t, err := readMulti(r, 1)
	if err != nil {
		return nil, err
	}
	if len(lt) > 0 {
		if bytes.IndexByte(lt, t[0]) == -1 {
			return nil, errors.New("type not match")
		}
	}

	ls, err := readMulti(r, 3)
	if err != nil {
		return nil, err
	}
	l := bytesToUint24(ls)

	ai, err := readMulti(r, 1)
	if err != nil {
		return nil, err
	}

	v, err := readMulti(r, l)
	if err != nil {
		return nil, err
	}
	return &Egg{Type: t[0], Len: l, Ai: ai[0], Value: v}, nil
}

func ReadFromBytes(buf []byte) (*Egg, error) {
	if len(buf) < TLAV_HEAD_LEN {
		return nil, errors.New("buffer too small")
	}
	pb := buf[0:4]
	if string(pb) != TLAV {
		return nil, errors.New("TLAV 头信息不匹配")
	}
	buf = buf[4:]
	t := buf[0:1]
	ls := buf[1:4]
	l := bytesToUint24(ls)
	ai := buf[4:5]
	need := l + 5
	if len(buf) < need {
		return nil, errors.New("buffer too small")
	}
	v := buf[5:need]
	return &Egg{Type: t[0], Len: l, Ai: ai[0], Value: v}, nil
}
