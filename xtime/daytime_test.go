package xtime

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"
)

func TestDayTimeBefore(t *testing.T) {
	t1 := NewDayTime(12, 30, 0)
	t2 := NewDayTime(13, 0, 0)
	if !t1.Before(t2) {
		log.Fatalf("t1:<%s> should be before t2:<%s>", t1, t2)
	}
}
func TestDayTimeAfter(t *testing.T) {
	t1 := NewDayTime(12, 30, 0)
	t2 := NewDayTime(13, 0, 0)
	if !t2.After(t1) {
		log.Fatalf("t2:<%s> should be After t1:<%s>", t2, t1)
	}
}

func TestDayTimeString(t *testing.T) {
	t1 := NewDayTime(12, 30, 0)
	if t1.String() != "12:30" {
		log.Fatalf("daytime instance String() should be  %s", "12:30")
	}
}

func TestDayTimeJSON(t *testing.T) {
	t1 := NewDayTime(12, 30, 0)
	b, err := json.Marshal(t1)
	if err != nil {
		log.Fatalf("DayTime MarshalJSON error:%+v", err)
		return
	}
	expect, _ := json.Marshal(45000)
	if !bytes.Equal(b, []byte(expect)) {
		log.Fatalf("json of t1:%+v expects to be :%s,but result is :%d", t1, expect, b)
	}
}

func TestDayTimeUnmarshalJSON(t *testing.T) {
	src := []byte{52, 53, 48, 48, 48} //45000
	dist := DayTime{}
	err := json.Unmarshal(src, &dist)
	if err != nil {
		log.Fatalf("UnmarshalJSON DayTime error:%+v", err)
		return
	}
	if dist.Hour != 12 || dist.Minute != 30 || dist.Second != 0 {
		log.Fatalf("src:%s unmarshal something wrong,exepcts the hour = %d,minute = %d,second = %d", src, 12, 30, 0)
		return
	}
}

func TestDayTimeRangeString(t *testing.T) {
	t1 := NewDayTime(10, 00, 00)
	t2 := NewDayTime(13, 30, 00)
	r := &DayTimeRange{
		From: t1,
		To:   t2,
	}
	result := r.String()
	expect := `{from:"10:00",to:"13:30"}`
	if result != expect {
		log.Fatalf("DayTimeRange String() should be :%s,but result is :%s", expect, result)
	}
}

func TestDayTimeRangeJSON(t *testing.T) {
	t1 := NewDayTime(10, 00, 00)
	t2 := NewDayTime(13, 30, 00)
	r := &DayTimeRange{
		From: t1,
		To:   t2,
	}
	result, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("DayTimeRange MarshalJSON error:%+v", err)
		return
	}
	var expect = []byte(`{"from":36000,"to":48600}`)

	if !bytes.Equal(result, expect) {
		log.Fatalf("DayTimeRange MarshalJSON() should be: %s,but result is: %s", expect, result)
	}
}
