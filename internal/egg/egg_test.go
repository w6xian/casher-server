package egg

import (
	"fmt"
	"testing"
)

func TestJsonArgs(t *testing.T) {
	g := EnJsonArgs(map[string]string{
		"a": "a1",
		"b": "b1",
		"c": "c1",
	})
	fmt.Println(g)
	if g[0] != "a,b,c" {
		t.Error("m[0]!=a,b,c")
	}
	if g[1] != "a1" {
		t.Error("m[1]!=a1")
	}
	if g[2] != "b1" {
		t.Error("m[2]!=b1")
	}
	if g[3] != "c1" {
		t.Error("m[3]!=c1")
	}
}
func TestDeJsonArgs(t *testing.T) {
	g := EnJsonArgs(map[string]string{
		"a": "a1",
		"b": "b1",
		"c": "c1",
	})
	d := DeJsonArgs(g)
	if d["a"] != "a1" {
		t.Error("d[a]!=a1")
	}
	if d["b"] != "b1" {
		t.Error("d[b]!=b1")
	}
	if d["c"] != "c1" {
		t.Error("d[c]!=c1")
	}
}
func TestFromMap(t *testing.T) {
	g := FromTypeValues(map[uint8]string{
		1: "a",
		2: "b",
	})
	if v, err := g.Maps(); err != nil {
		if k, ok := v[1]; ok {
			if k != "a" {
				t.Error("m[1]!=a")
			}
		}
		if k, ok := v[2]; ok {
			if k != "b" {
				t.Error("m[2]!=b")
			}
		}
	}
}

func TestUint64ToBytes(t *testing.T) {
	g := Uint64(12345678910)
	a := ToUint64(g)
	if a != 12345678910 {
		t.Error(g, a)
	}
}

func TestFromEvent(t *testing.T) {
	e := FromEventMap("Test", Arguments{
		"proxy_id": Uint64(2),
		"app_id":   []byte("abcdeabcdeabcdeabcde"),
	})
	data := e.Encode()
	eg, err := ReadFromBytes(data)
	if err != nil {
		t.Error(err)
	}
	event, arg := eg.MustEventMap()
	if event == "" {
		t.Error("解码失败")
	}
	if event != "Test" {
		t.Error("event != Test")
	}
	t.Log(event)
	t.Log("event name is", event)
	proxyId := arg.K("proxy_id").Uint64()
	t.Log("proxyId  is", proxyId)
	if proxyId != 2 {
		t.Error("proxy_id not equal 2")
	}
	appId := arg.K("app_id").String()
	t.Log("appId  is", appId)
	if appId != "abcdeabcdeabcdeabcde" {
		t.Error("appId not equal abcdeabcdeabcdeabcde")
	}
}
