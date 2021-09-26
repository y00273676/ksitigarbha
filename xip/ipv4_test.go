package xip_test

import (
	"encoding/json"
	"ksitigarbha/xip"
	"net"
	"strings"
	"testing"
)

func TestIPV4_MarshalText(t *testing.T) {
	s := struct {
		V xip.IPV4
	}{
		V: xip.IPV4(net.IPv4(192, 168, 1, 1)),
	}
	byt, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("expect no error. got %v", err)
	}
	if !strings.Contains(string(byt), "192.168.1.1") {
		t.Error("expect to have ip in json string")
	}
}

func TestIPV4_UnmarshalText(t *testing.T) {
	jsonVal := `"192.168.1.1"`
	var s xip.IPV4
	if err := json.Unmarshal([]byte(jsonVal), &s); err != nil {
		t.Fatalf("expect no error. got %v", err)
	}
	if s.String() != "192.168.1.1" {
		t.Error("expect to have ip in INet.String()")
	}
}

func TestLocal(t *testing.T) {
	var localAddr = xip.Local()
	t.Log(localAddr)
	var localIPV4 = xip.LocalIPV4()
	t.Log(localIPV4.String())
}
