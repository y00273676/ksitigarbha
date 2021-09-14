package xtime

import (
	"ksitigarbha/timezone"
	"log"
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	// 1506592453450 -- > Thu Sep 28 2017 17:54:13 GMT+0800 (CST)
	var source int64 = 1506592453450
	ts, err := ParseTimestamp(source)
	if err != nil {
		log.Fatalf("parse error:%v", err)
	}
	if ts.Year() != 2017 {
		log.Fatalf("Year expect to be :%v,but result is :%v", 2017, ts.Year())
	}
	if ts.Month() != time.September {
		log.Fatalf("Month expect to be :%v,but result is :%v", time.September, ts.Month())
	}
	if ts.Day() != 28 {
		log.Fatalf("Day expect to be :%v,but result is :%v", 28, ts.Day())
	}
	if ts.Hour() != 17 {
		log.Fatalf("Hour expect to be :%v,but result is :%v", 17, ts.Hour())
	}
	if ts.Minute() != 54 {
		log.Fatalf("Minute expect to be :%v,but result is :%v", 40, ts.Minute())
	}
	if ts.Second() != 13 {
		log.Fatalf("Second expect to be :%v,but result is :%v", 32, ts.Second())
	}
	if ts.format() != "2017-09-28 17:54:13" {
		log.Fatalf("expect to be :%s,but result is: %s", "2017-09-28 17:54:13", ts.format())
	}
}
func TestTimestampDateInLoc(t *testing.T) {
	var src int64 = 1551117600000
	var expectChinaDate = "2019-02-26"
	var expectUTCDate = "2019-02-25"
	var ts, _ = ParseTimestamp(src) //2019-02-25 18:00:00 GMT
	var date = ts.DateInLoc(timezone.China)
	//fmt.Printf("china date:%v\n",date.Time)
	if date.Format(ISO8601Date) != expectChinaDate {
		t.Fatalf("timestamp:%v,should be %s in china timezone,but result is: %v", src, expectChinaDate, date.Format(ISO8601Date))
	}
	var date2 = ts.DateInLoc(time.UTC)
	//fmt.Printf("utc date:%v\n",date2.Time)
	if date2.Format(ISO8601Date) != expectUTCDate {
		t.Fatalf("timestamp:%v,should be %s in china timezone,but result is: %v", src, expectUTCDate, date2.Format(ISO8601Date))
	}

}
