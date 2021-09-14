package xtime

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"ksitigarbha/timezone"
	"time"
)

// Date 总是代表对应日期的 0 点
type Date struct{ time.Time }

func (date Date) format() string {
	return date.Time.Format(ISO8601Date)
}

// AddDate adds any number of years, months, and days to the date.
// It proxies to the embedded time.Time, but returns a Date
func (date Date) AddDate(years, months, days int) Date {
	return Date{Time: date.Time.AddDate(years, months, days)}
}

// AddDays adds the given number of days to the date
func (date Date) AddDays(days int) Date {
	return date.AddDate(0, 0, days)
}

// After returns true if the given date is after (and not equal) to
// the current date
func (date Date) After(other Date) bool {
	return date.Time.After(other.Time)
}

// Before returns true if the given date is before (and not equal)
// to the current date
func (date Date) Before(other Date) bool {
	return date.Time.Before(other.Time)
}

// String returns the Date as a string
func (date Date) String() string {
	return date.format()
}

// Equals returns true if the dates are equal
func (date Date) Equals(other Date) bool {
	return date.Time.Equal(other.Time)
}

func (date Date) Days(other Date) int {
	hours := date.Time.Sub(other.Time).Hours()
	return int(hours) / HoursPerDay
}

// UnmarshalJSON converts a byte array into a Date
func (date *Date) UnmarshalJSON(text []byte) error {
	if string(text) == "null" {
		// Nulls are converted to zero times
		var zero Date
		*date = zero
		return nil
	}
	b := bytes.NewBuffer(text)
	dec := json.NewDecoder(b)
	var s string
	if err := dec.Decode(&s); err != nil {
		return err
	}
	value, err := time.ParseInLocation(ISO8601Date, s, timezone.China)
	if err != nil {
		return err
	}
	date.Time = value
	return nil
}

// MarshalJSON returns the JSON output of a Date.
// Null will return a zero value date.
func (date Date) MarshalJSON() ([]byte, error) {
	if date.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + date.format() + `"`), nil
}

// UnmarshalText converts a byte slice into a Date.
func (date *Date) UnmarshalText(text []byte) error {
	var s = string(text)
	if len(s) <= 0 {
		// Empty strings are converted to zero times.
		var zero Date
		*date = zero
		return nil
	}

	value, err := time.ParseInLocation(ISO8601Date, s, timezone.China)
	if err != nil {
		return err
	}
	date.Time = value
	return nil
}

// MarshalText returns a text output of a DateTime.
// An empty string will be returned for a zero value DateTime.
func (date Date) MarshalText() ([]byte, error) {
	if date.IsZero() {
		return []byte(""), nil
	}
	return []byte(date.format()), nil
}

// Scan converts an SQL value into a Date
func (date *Date) Scan(value interface{}) error {
	*date = Date{Time: NewChinaTime(value.(time.Time))}
	return nil
}

func (date Date) In(loc *time.Location) Date {
	theTime := date.Time
	t := theTime.In(loc)
	return NewDateFromTime(t)
}

// Value returns the date formatted for insert into PostgreSQL
func (date Date) Value() (driver.Value, error) {
	return date.format(), nil
}

func (date Date) BeginDateOfMonth() Date {
	y, m, _ := date.Date()
	return NewDate(y, m, 1, date.Location())
}
func (date Date) EndDateOfMonth() Date {
	return date.BeginDateOfMonth().AddDate(0, 1, -1)
	// 另一种实现
	//y, m, _ := date.Date()
	// 下个月的第0 天，就是上个月的最后一天，time 包实现的功能
	//return NewDate(y, time.Month(m+1), 0, date.Location())
}

// Within returns true if the Date is within the Range - inclusive
func (date Date) Within(term Range) bool {
	// Empty terms contain nothing
	if term.IsEmpty() {
		return false
	}
	// Only check if the range is bounded
	if !term.Start.IsZero() && date.Before(term.Start) {
		return false
	}
	if !term.End.IsZero() && date.After(term.End) {
		return false
	}
	return true
}

func (date Date) InfinityDate() InfinityDate {
	return InfinityDate{
		InfinityTime{
			Src: "",
			T:   date.Time,
		},
	}
}

// Today converts the local time to a Date
func Today() Date {
	return NewDateFromTime(time.Now())
}

// FromTime creates a Date from a time.Time
func NewDateFromTime(t time.Time) Date {
	y, m, d := t.Date()
	return NewDate(y, m, d, t.Location())
}

// NewDate creates a new Date
func NewDate(year int, month time.Month, day int, loc *time.Location) Date {
	return Date{Time: time.Date(year, month, day, 0, 0, 0, 0, loc)}
}

// ParseDate converts a ISO 8601 date string to a Date, possibly returning an error
func ParseDate(value string, loc *time.Location) (Date, error) {
	return ParseDateUsingLayout(ISO8601Date, value, loc)
}

// ParseDateUsingLayout calls Parse with a different date layout
func ParseDateUsingLayout(format, value string, loc *time.Location) (Date, error) {
	t, err := time.ParseInLocation(format, value, loc)
	if err != nil {
		return Date{}, err
	}
	return Date{Time: t}, nil
}
