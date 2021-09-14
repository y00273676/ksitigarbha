package xtime

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//ParseMonth 解析 yyyy-MM格式
func ParseMonth(src string) (month *Month, err error) {
	var items = strings.Split(src, "-")
	if len(items) != 2 {
		err = fmt.Errorf("invalid format of month:%v", src)
		return
	}
	var (
		theYear, theMonth int64
	)
	theYear, err = strconv.ParseInt(items[0], 10, 64)
	if err != nil {
		return
	}
	theMonth, err = strconv.ParseInt(items[1], 10, 64)
	if err != nil {
		return
	}
	month = &Month{
		Year:  int(theYear),
		Month: int(theMonth),
	}
	return
}

type Month struct {
	Year  int
	Month int
}

func (m *Month) Before(n *Month) bool {
	if m.Year < n.Year {
		return true
	}
	if m.Year == n.Year {
		return m.Month < n.Month
	}
	return false
}

func (m *Month) After(n *Month) bool {
	if m.Year > n.Year {
		return true
	}
	if m.Year == n.Year {
		return m.Month > n.Month
	}
	return false
}

func (m *Month) Equals(n *Month) bool {
	if m.Year == n.Year && m.Month == n.Month {
		return true
	}
	return false
}

func (m *Month) BeginDate(loc *time.Location) Date {
	return NewDate(m.Year, time.Month(m.Month), 1, loc)
}
func (m *Month) EndDate(loc *time.Location) Date {
	return m.BeginDate(loc).EndDateOfMonth()
}

func (m *Month) String() string {
	return fmt.Sprintf("%04d-%02d", m.Year, m.Month)
}

func (m *Month) AddMonths(delta int) *Month {
	var total = m.Year*12 + m.Month
	var abs = total + delta
	var year = abs / 12
	var month = abs % 12
	if month == 0 {
		month = 12
		year--
	}
	return &Month{Year: year, Month: month}
}

func (m *Month) Value() (driver.Value, error) {
	return int64(m.Number()), nil
}
func (m *Month) Scan(src interface{}) error {
	var int64Number = int(src.(int64))
	if month, err := ParseNumberMonth(int64Number); err != nil {
		return err
	} else {
		*m = *month
	}
	return nil
}
func (m *Month) MarshalJSON() ([]byte, error) {
	var display = m.String()
	return []byte(fmt.Sprintf(`"%s"`, display)), nil
}
func (m *Month) UnmarshalJSON(src []byte) error {
	inst, err := ParseMonth(string(src))
	if err != nil {
		return err
	}
	*m = *inst
	return nil
}

func (m *Month) Number() int {
	return m.Year*100 + m.Month //month一定<=12所以
}
func ParseNumberMonth(num int) (*Month, error) {
	var month = num % 100
	var year = num / 100
	if year <= 0 || month <= 0 {
		return nil, fmt.Errorf("错误的month,number:%v", num)
	}
	return &Month{Year: year, Month: month}, nil
}
