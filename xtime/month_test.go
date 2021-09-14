package xtime_test

import (
	"ksitigarbha/timezone"
	"ksitigarbha/xtime"
	"testing"
)

//ParseMonth 解析 yyyy-MM格式
func TestParseMonth(t *testing.T) {
	var testcase = []struct {
		Src   string
		Month *xtime.Month
	}{
		{Src: "2018-10", Month: &xtime.Month{Year: 2018, Month: 10}},
		{Src: "2018-11", Month: &xtime.Month{Year: 2018, Month: 11}},
		{Src: "2018-12", Month: &xtime.Month{Year: 2018, Month: 12}},
	}
	for _, item := range testcase {
		month, err := xtime.ParseMonth(item.Src)
		if err != nil {
			t.Fatalf("src:%v,expect month:%v,but there is error:%v", item.Src, item.Month, err)
		}
		if !month.Equals(item.Month) {
			t.Fatalf("src:%v,expect month:%v,but result is:%v", item.Src, item.Month, month)

		}
	}

}

func TestMonthBefore(t *testing.T) {
	var testcase = []struct {
		Src    *xtime.Month
		Target *xtime.Month
		Before bool
	}{
		{Src: &xtime.Month{2018, 10}, Target: &xtime.Month{2018, 10}, Before: false},
		{Src: &xtime.Month{2018, 11}, Target: &xtime.Month{2018, 10}, Before: false},
		{Src: &xtime.Month{2019, 10}, Target: &xtime.Month{2018, 10}, Before: false},
		{Src: &xtime.Month{2019, 9}, Target: &xtime.Month{2018, 10}, Before: false},
	}
	for _, item := range testcase {
		var result = item.Src.Before(item.Target)
		if result != item.Before {
			t.Fatalf("src:%v,target:%v,expect before:%v,but result is:%v", item.Src, item.Target, item.Before, result)
		}
	}
}

func TestMonthAddMonths(t *testing.T) {
	var testcase = []struct {
		Src    *xtime.Month
		Plus   int
		Target *xtime.Month
	}{
		{Src: &xtime.Month{2018, 10}, Plus: 1, Target: &xtime.Month{2018, 11}},
		{Src: &xtime.Month{2018, 11}, Plus: 2, Target: &xtime.Month{2019, 1}},
		{Src: &xtime.Month{2019, 10}, Plus: 13, Target: &xtime.Month{2020, 11}},
		{Src: &xtime.Month{2019, 9}, Plus: 10, Target: &xtime.Month{2020, 7}},
	}
	for _, item := range testcase {
		var result = item.Src.AddMonths(item.Plus)
		if !result.Equals(item.Target) {
			t.Fatalf("src:%v,plus:%v,expect result:%v,but result is:%v", item.Src, item.Plus, item.Target, result)
		}
	}
}

func TestMonthBeginEnd(t *testing.T) {
	var cases = []struct {
		Input            string
		BeginDateOfMonth string
		EndDateOfMonth   string
	}{
		{Input: "2000-02", BeginDateOfMonth: "2000-02-01", EndDateOfMonth: "2000-02-29"},
		{Input: "2008-02", BeginDateOfMonth: "2008-02-01", EndDateOfMonth: "2008-02-29"},
		{Input: "2009-02", BeginDateOfMonth: "2009-02-01", EndDateOfMonth: "2009-02-28"},
		{Input: "2009-03", BeginDateOfMonth: "2009-03-01", EndDateOfMonth: "2009-03-31"},
	}
	for _, c := range cases {
		m, _ := xtime.ParseMonth(c.Input)
		begin, _ := xtime.ParseDate(c.BeginDateOfMonth, timezone.China)
		end, _ := xtime.ParseDate(c.EndDateOfMonth, timezone.China)
		if !m.BeginDate(timezone.China).Equals(begin) || !m.EndDate(timezone.China).Equals(end) {
			t.Fatalf("input:%s,expect BeginDateOfMonth:%v,but result is: %v,expect EndDateOfMonth: %v,but result is: %v ",
				c.Input, c.BeginDateOfMonth, m.BeginDate(timezone.China), c.EndDateOfMonth, m.EndDate(timezone.China))

		}
	}
}
