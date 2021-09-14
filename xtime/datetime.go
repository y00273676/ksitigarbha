package xtime

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"time"

	"ksitigarbha/timezone"
)

type DateTime struct{ time.Time }

func (d DateTime) format() string {
	return d.Time.Format(DateTimePattern)
}
func (d DateTime) AddDates(years, months, days int) DateTime {
	return DateTime{Time: d.Time.AddDate(years, months, days)}
}
func (d DateTime) AddHours(hours int) DateTime {
	return DateTime{Time: d.Time.Add(time.Duration(hours) * time.Hour)}
}
func (d DateTime) AddMinutes(minutes int) DateTime {
	return DateTime{Time: d.Time.Add(time.Duration(minutes) * time.Minute)}
}
func (d DateTime) AddSeconds(secs int) DateTime {
	return DateTime{Time: d.Time.Add(time.Duration(secs) * time.Second)}
}

func (d DateTime) AddDays(days int) DateTime {
	return d.AddDates(0, 0, days)
}
func (d DateTime) After(other DateTime) bool {
	return d.Time.After(other.Time)
}
func (d DateTime) Before(other DateTime) bool {
	return d.Time.Before(other.Time)
}

// String returns the Date as a string
func (d DateTime) String() string {
	return d.format()
}

// Equals returns true if the dates are equal
func (d DateTime) Equals(other DateTime) bool {
	return d.Time.Equal(other.Time)
}
func (d DateTime) Days(other DateTime) int {
	//统一时区再进行计算.
	hours := d.Time.UTC().Sub(other.Time.UTC()).Hours()
	return int(hours) / HoursPerDay
}

//Date 按照该原本时间的Location
func (d DateTime) Date() Date {
	return NewDateFromTime(d.Time)
}

//DateInLocation 按照指定的loc进行转换
func (d DateTime) DateInLocation(loc *time.Location) Date {
	return NewDateFromTime(d.Time.In(loc))
}

func (d DateTime) In(loc *time.Location) DateTime {
	return NewDateTimeFromTime(d.Time.In(loc))
}

//Timestamp 转化为Timestamp
func (d DateTime) Timestamp() Timestamp {
	return Timestamp{d.Time}
}

// UnmarshalJSON converts a byte array into a Date
func (d *DateTime) UnmarshalJSON(text []byte) error {
	if string(text) == "null" {
		// Nulls are converted to zero times
		var zero DateTime
		*d = zero
		return nil
	}
	b := bytes.NewBuffer(text)
	dec := json.NewDecoder(b)
	var s string
	if err := dec.Decode(&s); err != nil {
		return err
	}
	value, err := time.ParseInLocation(DateTimePattern, s, timezone.China)
	if err != nil {
		return err
	}
	d.Time = value
	return nil
}

// MarshalJSON returns the JSON output of a Date.
// Null will return a zero value date.
func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.format() + `"`), nil
}

// UnmarshalText converts a byte array into a DateTime.
func (d *DateTime) UnmarshalText(text []byte) error {
	var s = string(text)
	if s == "" {
		// Empty texts are converted to zero times.
		var zero DateTime
		*d = zero
		return nil
	}
	value, err := time.ParseInLocation(DateTimePattern, s, timezone.China)
	if err != nil {
		return err
	}
	d.Time = value
	return nil
}

// MarshalText returns a text output of a DateTime.
// An empty string will be returned for a zero value DateTime.
func (d *DateTime) MarshalText() ([]byte, error) {
	if d.IsZero() {
		return []byte(""), nil
	}
	return []byte(d.format()), nil
}

// Scan converts an SQL value into a DateTime
func (d *DateTime) Scan(value interface{}) error {
	d.Time = value.(time.Time)
	return nil
}

// Value returns the d formatted for insert into database
func (d DateTime) Value() (driver.Value, error) {
	return d.Time, nil
}

func CurrentDateTime() DateTime {
	return NewDateTimeFromTime(time.Now())
}
func NewDateTimeFromTime(t time.Time) DateTime {
	year, month, day := t.Date()
	return NewDateTime(year, month, day, t.Hour(), t.Minute(), t.Second(), t.Location())
}

// New creates a new Date
func NewDateTime(year int, month time.Month, day, hour, min, sec int, loc *time.Location) DateTime {
	// Remove all second and nano second information and mark as UTC
	return DateTime{Time: time.Date(year, month, day, hour, min, sec, 0, loc)}
}

func ParseDateTime(value string, loc *time.Location) (DateTime, error) {
	return ParseDateTimeUsingLayout(DateTimePattern, value, loc)
}

// ParseDateTimeUsingLayout calls Parse with a different d layout
func ParseDateTimeUsingLayout(format, value string, loc *time.Location) (DateTime, error) {
	t, err := time.ParseInLocation(format, value, loc)
	if err != nil {
		return DateTime{}, err
	}
	return DateTime{Time: t}, nil
}
