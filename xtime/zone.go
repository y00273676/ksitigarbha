package xtime

import (
	"ksitigarbha/timezone"
	"time"
)

//NewTimeAccordingLocation 根据当前时间数据构造一个新的时区的的时间.解决有的时间不好根据时区解析的问题
func NewTimeAccordingLocation(t time.Time, loc *time.Location) time.Time {
	var y, m, d = t.Date()
	var dist = time.Date(y, m, d, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
	return dist
}

//NewChinaTime 转化为+08时区(中国时区)
func NewChinaTime(t time.Time) time.Time {
	return NewTimeAccordingLocation(t, timezone.China)
}
