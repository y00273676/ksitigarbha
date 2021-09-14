package xtime

import (
	"log"
	"testing"
	"time"
)

func TestParseDateTime(t *testing.T) {
	datetime, err := ParseDateTime("2017-09-28 15:40:32", time.Local)
	if err != nil {
		log.Fatalf("parse error:%v", err)
	}
	if datetime.Year() != 2017 {
		log.Fatalf("Year expect to be :%v,but result is :%v", 2017, datetime.Year())
	}
	if datetime.Month() != time.September {
		log.Fatalf("Year expect to be :%v,but result is :%v", time.September, datetime.Month())
	}
	if datetime.Day() != 28 {
		log.Fatalf("Year expect to be :%v,but result is :%v", 28, datetime.Day())
	}
	if datetime.Hour() != 15 {
		log.Fatalf("Year expect to be :%v,but result is :%v", 15, datetime.Hour())
	}
	if datetime.Minute() != 40 {
		log.Fatalf("Year expect to be :%v,but result is :%v", 40, datetime.Minute())
	}
	if datetime.Second() != 32 {
		log.Fatalf("Year expect to be :%v,but result is :%v", 32, datetime.Second())
	}
	if datetime.format() != "2017-09-28 15:40:32" {
		log.Fatalf("expect to be :%s,but result is: %s", "2017-09-28 15:40:32", datetime.format())
	}
}
