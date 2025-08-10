package egg

import (
	"encoding/binary"
	"io"
	"strings"
)

func EnJsonArgs(args map[string]string) []string {
	var ks []string
	var vs []string
	for k, v := range args {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	// rst := map[uint8]string{
	// 	0: strings.Join(ks, ","),
	// }

	var rst []string
	rst = append(rst, strings.Join(ks, ","))
	rst = append(rst, vs...)
	// for i, arg := range vs {
	// 	pos := uint8(i + 1)
	// 	rst[pos] = arg
	// }
	return rst
}

func DeJsonArgs(params []string) map[string]string {
	var rst = map[string]string{}
	if len(params) >= 2 {
		v := params[0]
		ks := strings.Split(v, ",")
		params = params[1:]
		if len(ks) == len(params) {
			for k, v := range ks {
				rst[v] = params[k]
			}
		}
	}
	return rst
}

func EnMapArgments(args Arguments) []Value {
	var ks []string
	var vs []Value
	vs = append(vs, Value(""))
	for k, v := range args {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	keys := Value(strings.Join(ks, ","))
	vs[0] = keys
	return vs
}

func DeMapArgments(params []Value) Arguments {
	var rst = Arguments{}
	var kv Value = params[0]
	ks := strings.Split(string(kv), ",")
	if len(ks) == len(params)-1 {
		for k, v := range ks {
			rst[v] = params[k+1]
		}
	}

	return rst
}

func readMulti(r io.Reader, l int) ([]byte, error) {
	b := make([]byte, 0, l)
	for {
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
		if l == len(b) {
			return b, nil
		}
	}
}

func uint24ToBytes(v uint32) []byte {
	v = v & 0x00ffffff
	b := make([]byte, 3)
	b[0] = byte(v)
	b[1] = byte((v >> 8) & 0xff)
	b[2] = byte((v >> 16) & 0xff)
	return b
}

func Uint64(i uint64) Value {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func Int64(i int64) Value {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func String(v string) Value {
	return []byte(v)
}

func ToUint64(buf []byte) uint64 {
	return binary.LittleEndian.Uint64(buf)
}

func uint8ToBytes(v uint8) []byte {
	b := make([]byte, 1)
	b[0] = v
	return b
}

func bytesToUint24(buf []byte) int {
	if len(buf) < 3 {
		return 0
	}
	v := int(buf[0])
	a := int(buf[1]) << 8
	b := int(buf[2]) << 16
	v += int(a)
	v += int(b)
	return v
}
