package fresh

import (
	"fmt"
	"time"

	"github.com/askasoft/pango/str"
	"github.com/askasoft/pango/tmu"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = time.RFC3339 //"2006-01-02T15:04:05Z07:00"
)

type Date struct {
	time.Time
}

func ParseDate(s string) (Date, error) {
	t, err := tmu.ParseInLocation(s, time.UTC, DateFormat)
	if err != nil {
		return Date{}, err
	}
	return Date{t}, nil
}

func (d *Date) String() string {
	return d.Format(DateFormat)
}

func (d *Date) MarshalJSON() ([]byte, error) {
	bs := make([]byte, 0, len(DateFormat)+2)
	bs = append(bs, '"')
	bs = d.AppendFormat(bs, DateFormat)
	bs = append(bs, '"')
	return bs, nil
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	js := str.UnsafeString(data)

	// Ignore null, like in the main JSON package.
	if js == "" || js == "null" {
		return
	}

	if !str.StartsWithByte(js, '"') || !str.EndsWithByte(js, '"') {
		return fmt.Errorf("fresh: invalid date format %q", js)
	}

	d.Time, err = tmu.ParseInLocation(js[1:len(js)-1], time.UTC, DateFormat)
	return
}

type Time struct {
	time.Time
}

func ParseTime(s string) (Time, error) {
	t, err := tmu.ParseInLocation(s, time.UTC, TimeFormat, DateFormat)
	if err != nil {
		return Time{}, err
	}
	return Time{t}, nil
}

func (t *Time) String() string {
	return t.Time.UTC().Format(TimeFormat)
}

func (t *Time) MarshalJSON() ([]byte, error) {
	bs := make([]byte, 0, len(TimeFormat)+2)
	bs = append(bs, '"')
	bs = t.Time.UTC().AppendFormat(bs, TimeFormat)
	bs = append(bs, '"')
	return bs, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	// Ignore null, like in the main JSON package.
	js := str.UnsafeString(data)
	if js == "null" {
		return
	}

	if !str.StartsWithByte(js, '"') || !str.EndsWithByte(js, '"') {
		return fmt.Errorf("fresh: invalid time format %q", js)
	}

	// parse time and date format to prevent error
	t.Time, err = tmu.ParseInLocation(js[1:len(js)-1], time.UTC, TimeFormat, DateFormat)
	return
}

type TimeSpent = tmu.HHMM

func ParseTimeSpent(s string) (TimeSpent, error) {
	return tmu.ParseHHMM(s)
}
