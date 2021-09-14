package xtime

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"

	"ksitigarbha/timezone"
)

var (
	_Infinity              = "infinity"
	_NegativeInfinity      = "-infinity"
	_InfinityBytes         = []byte("infinity")
	_NegativeInfinityBytes = []byte("-infinity")
)

func ParseInfinityDate(src string) (InfinityDate, error) {
	var dist InfinityDate
	if IsInfinite(src) {
		dist.Src = src
	} else {
		var err error
		dist.InfinityTime.T, err = time.Parse(ISO8601Date, src)
		if err != nil {
			return dist, fmt.Errorf("invalid InfinityDate:%v", src)
		}
	}
	return dist, nil

}

func ParseInfinityTime(src string) (InfinityTime, error) {
	var dist InfinityTime
	if IsInfinite(src) {
		dist.Src = src
	} else {
		ts, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return dist, fmt.Errorf("invalid InfinityTime:%v", src)
		}
		t := time.Unix(ts/1000, 0).In(timezone.China)
		dist.T = t
	}
	return dist, nil
}

//InfinityTime 的json输出都保证是字符串.这样能够让类型是统一的.
//同时如果不是infinity或者-inifinity的时候,输出的的总是timestamp值.
//这样能够让InfinityTime适应各种情况.时间戳，日期都能使用.客户端格式化到相应的字段即可
type InfinityTime struct {
	Src string
	T   time.Time
}

func IsInfinite(src string) bool {
	if src == _Infinity || src == _NegativeInfinity {
		return true
	}
	return false
}

func IsInfiniteBytes(src []byte) bool {
	if bytes.Equal(_InfinityBytes, src) || bytes.Equal(_NegativeInfinityBytes, src) {
		return true
	}
	return false
}

func (i InfinityTime) IsInfinite() bool {
	return IsInfinite(i.Src)
}

func (i InfinityTime) Timestamp() int64 {
	return i.T.Unix() * 1000
}
func (i InfinityTime) TimestampString() string {
	return strconv.FormatInt(i.T.Unix()*1000, 10)
}

func (i InfinityTime) String() string {
	if i.Src == _Infinity || i.Src == _NegativeInfinity {
		return i.Src
	}
	return i.T.In(timezone.China).Format(TimestampPattern)
}

func (i InfinityTime) Value() (driver.Value, error) {
	if i.IsInfinite() {
		return i.Src, nil
	}
	return i.T, nil
}

func (i *InfinityTime) Scan(src interface{}) error {
	var tmp InfinityTime
	switch t := src.(type) {
	case []byte:
		tmp.Src = string(t)
	case time.Time:
		tmp.T = NewChinaTime(t)
	}
	*i = tmp
	return nil
}

func (i InfinityTime) MarshalJSON() ([]byte, error) {
	if i.Src != "" {
		return []byte(`"` + i.Src + `"`), nil
	}
	return []byte(`"` + i.TimestampString() + `"`), nil
}

func (i *InfinityTime) UnmarshalJSON(src []byte) error {
	str := string(src)
	if IsInfinite(str) {
		i.Src = str
	} else {
		ts, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid InfinityTime:%v", str)
		}
		i.T = time.Unix(ts/1000, 0).In(timezone.China)
	}
	return nil
}

type InfinityDate struct {
	InfinityTime
}

func (i *InfinityDate) MarshalJSON() ([]byte, error) {
	if i.InfinityTime.Src != "" {
		return i.InfinityTime.MarshalJSON()
	}
	return []byte(`"` + i.InfinityTime.T.Format(ISO8601Date) + `"`), nil
}

func (i *InfinityDate) UnmarshalJSON(src []byte) error {
	toCheck, err := strconv.Unquote(string(src))
	if err != nil {
		return fmt.Errorf("invalid InfinityDate:%s,Unquote fail", src)
	}
	var tmp InfinityTime
	if IsInfinite(toCheck) {
		tmp.Src = toCheck
	} else {
		t, err := time.ParseInLocation(ISO8601Date, toCheck, timezone.China)
		if err != nil {
			return fmt.Errorf("Invalid InfinityDate:%s", src)
		}
		tmp.T = t
	}
	*i = InfinityDate{InfinityTime: tmp}
	return nil
}
func (i InfinityDate) Value() (driver.Value, error) {
	if i.InfinityTime.Src != "" {
		return i.InfinityTime.Src, nil
	}
	return i.InfinityTime.T.Format(ISO8601Date), nil
}

func (i InfinityDate) String() string {
	if i.InfinityTime.Src != "" {
		return i.InfinityTime.Src
	}
	return i.InfinityTime.T.Format(ISO8601Date)
}

func (i *InfinityDate) Scan(src interface{}) error {
	var tmp InfinityTime
	switch t := src.(type) {
	case []byte:
		tmp.Src = string(t)
	case time.Time:
		tmp.T = NewChinaTime(t)
	}
	*i = InfinityDate{InfinityTime: tmp}
	return nil
}

func (obj InfinityDate) After(another InfinityDate) bool {
	if obj.IsInfinite() && another.IsInfinite() {
		return false
	}
	if obj.IsInfinite() {
		return true
	}
	if another.IsInfinite() {
		return false
	}
	return obj.InfinityTime.T.After(another.InfinityTime.T)
}

func (obj InfinityDate) Before(another InfinityDate) bool {
	if obj.IsInfinite() && another.IsInfinite() {
		return false
	}
	if obj.IsInfinite() {
		return false
	}
	if another.IsInfinite() {
		return true
	}
	return obj.InfinityTime.T.Before(another.InfinityTime.T)
}
