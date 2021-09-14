package xtime

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"ksitigarbha/timezone"
)

type Timestamp struct{ time.Time }

func (m Timestamp) timestamp() int64 {
	return m.UnixMilli()
}
func (m Timestamp) UnixMilli() int64 {
	return m.Time.UnixNano() / int64(time.Millisecond)
}
func (m Timestamp) format() string {
	return m.Time.Format(TimestampPattern)
}
func (m Timestamp) AddDates(years, months, days int) Timestamp {
	return Timestamp{Time: m.Time.AddDate(years, months, days)}
}
func (m Timestamp) AddHours(hours int) Timestamp {
	return Timestamp{Time: m.Time.Add(time.Duration(hours) * time.Hour)}
}
func (m Timestamp) AddMinutes(minutes int) Timestamp {
	return Timestamp{Time: m.Time.Add(time.Duration(minutes) * time.Minute)}
}
func (m Timestamp) AddSeconds(secs int) Timestamp {
	return Timestamp{Time: m.Time.Add(time.Duration(secs) * time.Second)}
}

func (m Timestamp) AddDays(days int) Timestamp {
	return m.AddDates(0, 0, days)
}
func (m Timestamp) After(other Timestamp) bool {
	return m.Time.After(other.Time)
}
func (m Timestamp) Before(other Timestamp) bool {
	return m.Time.Before(other.Time)
}

// String returns the Date as a string
func (m Timestamp) String() string {
	return m.format()
}

// Equals returns true if the dates are equal
func (m Timestamp) Equals(other Timestamp) bool {
	return m.Time.Equal(other.Time)
}
func (m Timestamp) Days(other Timestamp) int {
	hours := m.Time.Sub(other.Time).Hours()
	return int(hours) / HoursPerDay
}

// UnmarshalJSON converts a byte array into a Date
func (m *Timestamp) UnmarshalJSON(text []byte) error {
	var ts int64
	err := json.Unmarshal(text, &ts)
	if err != nil {
		return err
	}
	t := time.Unix(ts/1000, 0).In(timezone.China)
	*m = Timestamp{t}
	return nil
}

// MarshalJSON returns the JSON output of a Date.
// Null will return a zero value date.
func (m Timestamp) MarshalJSON() ([]byte, error) {
	if m.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(m.timestamp())
}

// UnmarshalText converts a byte array into a Timestamp.
func (m *Timestamp) UnmarshalText(text []byte) error {
	ts, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}
	t := time.Unix(ts/1000, 0).In(timezone.China)
	*m = Timestamp{t}
	return nil
}

func (m *Timestamp) DateInLoc(loc *time.Location) Date {
	var t = m.Time.In(loc)
	year, month, date := t.Date()
	return NewDate(year, month, date, loc)
}

// MarshalText returns the text of a Timestamp.
// An empty string will be returned for a zero value Timestamp.
func (m Timestamp) MarshalText() ([]byte, error) {
	if m.IsZero() {
		return []byte(""), nil
	}
	return []byte(fmt.Sprintf("%d", m.timestamp())), nil
}

// Scan converts an SQL value into a Timestamp
func (m *Timestamp) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	m.Time = value.(time.Time)
	return nil
}

// Value returns the Timestamp formatted for insert into database
func (m Timestamp) Value() (driver.Value, error) {
	return m.Time, nil
}

func CurrentTimestamp() Timestamp {
	return NewTimestampFromTime(time.Now())
}

func NewTimestampFromTime(t time.Time) Timestamp {
	year, month, day := t.Date()
	return NewMilliSecondTimeInLocation(year, month, day, t.Hour(), t.Minute(), t.Second(), t.Location())
}

func NewMilliSecondTimeInLocation(year int, month time.Month, day, hour, min, sec int, loc *time.Location) Timestamp {
	// Remove all second and nano second information and mark as UTC
	return Timestamp{Time: time.Date(year, month, day, hour, min, sec, 0, loc)}
}

// New creates a new Date
func NewMilliSecondTime(year int, month time.Month, day, hour, min, sec int) Timestamp {
	return NewMilliSecondTimeInLocation(year, month, day, hour, min, sec, time.UTC)
}

func ParseTimestamp(value int64) (Timestamp, error) {
	t := time.Unix(value/1000, 0)
	return Timestamp{t}, nil
}
