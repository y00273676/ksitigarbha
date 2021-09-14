package xtime

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"ksitigarbha/timezone"
)

type MilliSeconds struct {
	Src int64
	T   time.Time
}

func (m MilliSeconds) timestamp() int64 {
	return m.Src
}

func (m MilliSeconds) format() string {
	return m.T.Format(TimestampPattern)
}

func (m MilliSeconds) String() string {
	return m.format()
}

func (m MilliSeconds) Equal(other MilliSeconds) bool {
	return m.Src == other.Src
}

func (m MilliSeconds) After(other MilliSeconds) bool {
	return m.Src > other.Src
}
func (m MilliSeconds) Before(other MilliSeconds) bool {
	return m.Src < other.Src
}

func (m *MilliSeconds) UnmarshalJSON(text []byte) error {
	var ts int64
	err := json.Unmarshal(text, &ts)
	if err != nil {
		return err
	}
	m.Src = ts
	t := time.Unix(ts/1000, 0).In(timezone.China)
	m.T = t
	return nil
}

func (m *MilliSeconds) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Src)
}

func (m *MilliSeconds) UnmarshalText(text []byte) error {
	ts, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return err
	}
	m.Src = ts
	var t = time.Unix(ts/1000, 0).In(timezone.China)
	m.T = t
	return nil
}

func (m *MilliSeconds) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", m.Src)), nil
}

func (m *MilliSeconds) Scan(value interface{}) error {
	intVal := value.(int64)
	m.Src = intVal
	t := time.Unix(intVal/1000, 0).In(time.Local)
	m.T = t
	return nil
}

func (m *MilliSeconds) Value() (driver.Value, error) {
	return m.Src, nil
}
