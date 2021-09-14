package xtime_test

import (
	"ksitigarbha/timezone"
	"ksitigarbha/xtime"
	"testing"
)

func TestFirstMondayOfMonth(t *testing.T) {
	var testcase = []struct {
		Year   int
		Month  int
		Expect string
	}{
		{Year: 2018, Month: 10, Expect: "2018-10-01"},
		{Year: 2018, Month: 11, Expect: "2018-11-05"},
	}
	for _, item := range testcase {
		var result = xtime.FirstMondayOfMonth(item.Year, item.Month)
		if result.Format(xtime.ISO8601Date) != item.Expect {
			t.Fatalf("year:%v,month:%v,expect result:%v,but result is:%v", item.Year, item.Month, item.Expect, result)
		}
	}
}

func TestLastSundayOfMonth(t *testing.T) {
	var testcase = []struct {
		Year   int
		Month  int
		Expect string
	}{
		{Year: 2018, Month: 10, Expect: "2018-11-04"},
		{Year: 2018, Month: 11, Expect: "2018-12-02"},
		{Year: 2018, Month: 12, Expect: "2019-01-06"},
	}
	for _, item := range testcase {
		var result = xtime.LastSundayOfMonth(item.Year, item.Month)
		if result.Format(xtime.ISO8601Date) != item.Expect {
			t.Fatalf("year:%v,month:%v,expect result:%v,but result is:%v", item.Year, item.Month, item.Expect, result)
		}
	}
}

func TestWeeksOfMonth(t *testing.T) {
	var testcase = []struct {
		Year   int
		Month  int
		Expect int
	}{
		{Year: 2018, Month: 10, Expect: 5},
		{Year: 2018, Month: 11, Expect: 4},
		{Year: 2018, Month: 12, Expect: 5},
	}
	for _, item := range testcase {
		var result = xtime.WeeksOfMonth(item.Year, item.Month)
		if result != item.Expect {
			t.Fatalf("year:%v,month:%v,expect result:%v,but result is:%v", item.Year, item.Month, item.Expect, result)
		}
	}
}
func TestMonths(t *testing.T) {
	var testcase = []struct {
		Begin  string
		End    string
		Expect int
	}{
		{Begin: "2018-10-01", End: "2019-02-28", Expect: 5},
		{Begin: "2018-10-01", End: "2019-12-28", Expect: 15},
		{Begin: "2018-10-01", End: "2018-11-01", Expect: 2},
		{Begin: "2018-10-01", End: "2018-10-30", Expect: 1},
	}
	for _, item := range testcase {
		var beginDate, _ = xtime.ParseDate(item.Begin, timezone.China)
		var endDate, _ = xtime.ParseDate(item.End, timezone.China)
		var result = xtime.Months(beginDate, endDate)
		if result != item.Expect {
			t.Fatalf("begin:%v,end:%v,result is:%v,but expect is:%v", item.Begin, item.End, result, item.Expect)
		}
	}
}
