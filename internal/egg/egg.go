package egg

import (
	"bytes"
	"errors"
)

type Egg struct {
	Type  uint8
	Len   int
	Ai    uint8
	Value Value
}

func (egg *Egg) Encode() []byte {
	rst := make([]byte, 0, egg.Len+TLAV_HEAD_LEN)
	rst = append(rst, []byte(TLAV)...)
	rst = append(rst, uint8ToBytes(egg.Type)...)
	rst = append(rst, uint24ToBytes(uint32(egg.Len))...)
	rst = append(rst, uint8ToBytes(egg.Ai)...)
	rst = append(rst, egg.Value...)
	return rst
}

func (egg *Egg) Maps() (map[uint8]string, error) {
	if egg.Type != EGG_TYPE_VALUES {
		return nil, errors.New("not a map")
	}
	m := map[uint8]string{}
	buf := egg.Value[0:]
	for {
		if len(buf) < 5 {
			break
		}
		pb := buf[0:4]
		if string(pb) != TLAV {
			break
		}
		buf = buf[4:]
		k := buf[0:1]
		ls := buf[1:4]
		l := bytesToUint24(ls)
		// ai := buf[4:5]
		need := l + 5
		if len(buf) < need {
			break
		}
		v := buf[5:need]
		m[uint8(k[0])] = string(v)
		buf = buf[need:]
	}
	return m, nil
}

func (egg *Egg) MustJson() map[string]string {
	var m []string
	buf := egg.Value[0:]
	i := 0
	for {
		if len(buf) < 5 {
			break
		}
		pb := buf[0:4]
		if string(pb) != TLAV {
			break
		}
		buf = buf[4:]
		ls := buf[1:4]
		l := bytesToUint24(ls)
		// ai := buf[4:5]
		need := l + 5
		if len(buf) < need {
			break
		}
		v := buf[5:need]
		m = append(m, v.String())
		buf = buf[need:]
		i = i + 1
	}
	return DeJsonArgs(m)
}
func (egg *Egg) MustValues() []Value {
	var m []Value
	buf := egg.Value[0:]
	for {
		if len(buf) < 5 {
			break
		}
		pb := buf[0:4]
		if string(pb) != TLAV {
			break
		}
		buf = buf[4:]
		ls := buf[1:4]
		l := bytesToUint24(ls)
		// ai := buf[4:5]
		need := l + 5
		if len(buf) < need {
			break
		}
		v := buf[5:need]
		m = append(m, v)
		buf = buf[need:]
	}
	return m
}

func (egg *Egg) Length() int {
	return egg.Len + TLAV_HEAD_LEN
}

func (egg *Egg) MustEventMap() (string, Arguments) {
	var name string
	var kv Arguments
	if egg.Type == EGG_TYPE_EVENT_MAP || egg.Type == EGG_TYPE_EVENT_BACK_MAP {
		buf := egg.Value[0:]
		r := bytes.NewReader(buf)
		if n, err := ReadEggWithType(r, EGG_TYPE_EVENT_NAME); err == nil {
			if a, err := ReadEggWithType(r, EGG_TYPE_EVENT_ARGUMENTS); err == nil {
				name = n.Value.String()
				m := a.MustValues()
				kv = DeMapArgments(m)
			}
		}
	}
	return name, kv
}

func (egg *Egg) MustSettingMap() (string, Arguments) {
	var name string
	var kv Arguments
	if egg.Type == EGG_TYPE_SETTING || egg.Type == EGG_TYPE_EVENT_BACK_MAP {
		buf := egg.Value[0:]
		r := bytes.NewReader(buf)
		if n, err := ReadEggWithType(r, EGG_TYPE_EVENT_NAME); err == nil {
			if a, err := ReadEggWithType(r, EGG_TYPE_EVENT_ARGUMENTS); err == nil {
				name = n.Value.String()
				m := a.MustValues()
				kv = DeMapArgments(m)
			}
		}
	}
	return name, kv
}

func (egg *Egg) MustEventList() (string, []*Egg) {
	var name string
	var ls []*Egg
	if egg.Type == EGG_TYPE_EVENT_LIST {
		buf := egg.Value[0:]
		r := bytes.NewReader(buf)
		if n, err := ReadEggWithType(r, EGG_TYPE_EVENT_NAME); err == nil {
			name = n.Value.String()
			for {
				if e, err := ReadEgg(r); err == nil {
					ls = append(ls, e)
				} else {
					break
				}
			}
		}
	}
	return name, ls
}
