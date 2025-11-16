package buff

import (
	"bytes"
	"encoding/json"
	"sync"
)

var bufPool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

func Build(data any) []byte {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)
	json.NewEncoder(buf).Encode(data)
	return buf.Bytes()
}
