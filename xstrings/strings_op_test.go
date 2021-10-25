package xstrings_test

import (
	"ksitigarbha/xstrings"
	"log"
	"testing"
)

func TestSub(t *testing.T) {
	cases := []struct {
		Src    string
		Start  int
		Length int
		Expect string
	}{
		{"hello", 0, 5, "hello"},
		{"hello", 0, 4, "hell"},
		{"hello", 0, 10, "hello"},
	}
	for _, c := range cases {
		result := xstrings.Sub(c.Src, c.Start, c.Length)
		if result != c.Expect {
			log.Fatalf("input:%s,start:%v,length:%v,expoect:%v, but result is :%v", c.Src, c.Start, c.Length, c.Expect, result)
		}
	}
}

func TestChop(t *testing.T) {
	cases := []struct {
		Src    string
		Start  int
		End    int
		Expect string
	}{
		{"hello", 0, 5, "hello"},
		{"hello", 0, 4, "hell"},
		{"hello", 0, 0, ""},
	}
	for _, c := range cases {
		result := xstrings.Sub(c.Src, c.Start, c.End)
		if result != c.Expect {
			log.Fatalf("input:%s,start:%v,end:%v,expoect:%v, but result is :%v", c.Src, c.Start, c.End, c.Expect, result)
		}
	}
}
