package xtime

import (
	"ksitigarbha/timezone"
	"time"
)

var oneWeekDuration = int64(7 * 24 * time.Hour)
var monthFirstDayPlus = map[time.Weekday]int{
	time.Monday:    0,
	time.Tuesday:   6,
	time.Wednesday: 5,
	time.Thursday:  4,
	time.Friday:    3,
	time.Saturday:  2,
	time.Sunday:    1,
}

var monthLastDayPlus = map[time.Weekday]int{
	time.Monday:    6,
	time.Tuesday:   5,
	time.Wednesday: 4,
	time.Thursday:  3,
	time.Friday:    2,
	time.Saturday:  1,
	time.Sunday:    0,
}

func FirstMondayOfMonth(year, month int) time.Time {
	var monthFirstDay = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, timezone.China)
	var plusDay = monthFirstDayPlus[monthFirstDay.Weekday()]
	var targetDate = monthFirstDay.AddDate(0, 0, plusDay)
	return targetDate
}

//isLeap 拷贝的time包的内部实现
func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

//DaysOfMonth 一个月有多少天
func DaysOfMonth(year, month int) int {
	if month == 2 {
		var daysOfFeb = 28
		if isLeap(year) {
			daysOfFeb += 1
		}
		return daysOfFeb
	}
	//以7月为断点，7月之前奇数月31天，7月之后偶数月31天，2月单独计算
	return 31 - (month-1)%7%2
}

//一个月的最后一个周日.以自然月的最后一天如果已经开始一周，要扩展到下个月.直到一个周日出现
func LastSundayOfMonth(year, month int) time.Time {
	var days = DaysOfMonth(year, month)
	var monthLastDay = time.Date(year, time.Month(month), days, 23, 59, 59, 999, timezone.China)
	var plus = monthLastDayPlus[monthLastDay.Weekday()]
	var targetDate = monthLastDay.AddDate(0, 0, plus)
	return targetDate
}

//Months计算两个日期之间相隔多少个自然月.begin<=end(2018-10-01,2018-11-01算两个月)
func Months(begin, end Date) int {
	var (
		beginYear  = begin.Year()
		beginMonth = begin.Month()
		endYear    = end.Year()
		endMonth   = end.Month()
	)
	return 12*(endYear-beginYear) + (int(endMonth) - int(beginMonth)) + 1 //算自然月。+1

}

//WeeksOfMonth 一个月中一共有多少个完整的周.
// 一月的第一个周一为第一周.
//最后一天如果不是周日要往下个月继续算,到已经开始的这周完整到周日截止.
func WeeksOfMonth(year, month int) int {
	var theDay = FirstMondayOfMonth(year, month)
	var count = 1
	//怎么也不可能超过5周.
	for i := 0; i < 5; i++ {
		theDay = theDay.AddDate(0, 0, 7)
		if theDay.Month() != time.Month(month) {
			break
		}
		count++
	}
	return count
}
