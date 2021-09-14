package xtime

import (
	"encoding/json"
	"ksitigarbha/timezone"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	day := NewDate(2015, 3, 1, time.Local)
	if day.String() != "2015-03-01" {
		t.Error("The string output of March 1st should be 2015-03-01")
	}

	output, err := json.Marshal(day)
	if err != nil {
		t.Error("JSON marshaling of dates should not error")
	}
	if string(output) != `"2015-03-01"` {
		t.Error(`JSON marshaling of March 1st should be "2015-03-01"`)
	}

	// Zero dates should return null
	var zero Date
	output, err = json.Marshal(zero)
	if err != nil {
		t.Error("JSON marshaling of zero dates should not error")
	}
	if string(output) != "null" {
		t.Error("json.Marshal of a zero date should be null")
	}

	nextDay := NewDate(2015, 3, 2, time.Local)
	if !nextDay.Equals(day.AddDays(1)) {
		t.Error("The day after March 1st should be March 2nd")
	}

	parsed, err := ParseDate("2015-03-01", time.Local)
	if err != nil {
		t.Error("Parsing of properly formatted dates should not error")
	}
	if !parsed.Equals(day) {
		t.Error("The parsed string should equal March 1st")
	}

	if day.UnmarshalJSON([]byte(`"2015-03-01"`)) != nil {
		t.Error("UnmarshalJSON of a valid slice of bytes should not error")
	}

	// Parsing null should return a zero date
	if zero.UnmarshalJSON([]byte("null")) != nil {
		t.Error("Unmarshaling a null date should not error")
	}
	if !zero.IsZero() {
		t.Error("A null date should unmarshal to zero")
	}

	jan1 := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	if !jan1.Before(day.Time) {
		t.Error("January 1st should be before March 1st")
	}
}

func TestDate_AddDate(t *testing.T) {
	feb1 := NewDate(2016, 2, 1, time.Local)
	if feb1.AddDate(0, 1, 0) != NewDate(2016, 3, 1, time.Local) {
		t.Error("February 1st plus one month should be March 1st")
	}

	if feb1.AddDate(0, 1, -1) != NewDate(2016, 2, 29, time.Local) {
		t.Error("February 1st plus one month and minus one day should be February 29th (in 2016)")
	}

	if feb1.AddDate(1, -1, 0) != NewDate(2017, 1, 1, time.Local) {
		t.Errorf("February 1st plus one year and minus one month should be January 1st, 2017")
	}
}

func TestDate_Within(t *testing.T) {
	march1 := NewDate(2015, 3, 1, time.Local)
	dec1 := NewDate(2015, 12, 1, time.Local)

	feb := EntireMonth(2015, 2, time.Local)
	march := EntireMonth(2015, 3, time.Local)

	if march1.Within(feb) {
		t.Error("March 1st should not be within February")
	}
	if march1 != march.Start {
		t.Error("March 1st should equal the start of March")
	}
	if !march1.Within(march) {
		t.Error("March 1st should be within March")
	}

	// Test unbounded ranges
	novOnward := Range{Start: NewDate(2015, 11, 1, time.Local)}
	beforeNov := Range{End: NewDate(2015, 10, 31, time.Local)}

	if !dec1.Within(novOnward) {
		t.Error("December 1st should be within November onward")
	}
	if dec1.Within(beforeNov) {
		t.Error("December 1st should not be within before November")
	}

	if !march1.Within(beforeNov) {
		t.Error("March 1st should be within before November")
	}
	if march1.Within(novOnward) {
		t.Error("March 1st should not be within November onward")
	}
}

func TestDate_Month(t *testing.T) {
	var cases = []struct {
		Input            string
		BeginDateOfMonth string
		EndDateOfMonth   string
	}{
		{Input: "2000-02-01", BeginDateOfMonth: "2000-02-01", EndDateOfMonth: "2000-02-29"},
		{Input: "2008-02-01", BeginDateOfMonth: "2008-02-01", EndDateOfMonth: "2008-02-29"},
		{Input: "2009-02-01", BeginDateOfMonth: "2009-02-01", EndDateOfMonth: "2009-02-28"},
		{Input: "2009-02-15", BeginDateOfMonth: "2009-02-01", EndDateOfMonth: "2009-02-28"},
		{Input: "2009-12-15", BeginDateOfMonth: "2009-12-01", EndDateOfMonth: "2009-12-31"},
	}
	for _, c := range cases {
		d, _ := ParseDate(c.Input, timezone.China)
		begin, _ := ParseDate(c.BeginDateOfMonth, timezone.China)
		end, _ := ParseDate(c.EndDateOfMonth, timezone.China)
		if !d.BeginDateOfMonth().Equals(begin) || !d.EndDateOfMonth().Equals(end) {
			t.Fatalf("input:%s,expect BeginDateOfMonth:%v,but result is: %v,expect EndDateOfMonth: %v,but result is: %v ", c.Input, c.BeginDateOfMonth, d.BeginDateOfMonth(), c.EndDateOfMonth, d.EndDateOfMonth())
		}
	}

}
