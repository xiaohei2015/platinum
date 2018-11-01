package platinum

import (
	"time"
)

type Time struct {
	time   time.Time
	format string
}

func (t *Time) String() string {
	return t.time.Format(t.format)
}

func Create(tm time.Time, format string) *Time {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return &Time{
		time:   tm,
		format: format,
	}
}

func Now() *Time {
	return Create(time.Now(), "")
}

func CreateFromDate(year int, month time.Month, day int) *Time {
	return Create(time.Date(year, month, day, 0, 0, 0, 0, time.Local), "")
}

func CreateFromDateWithLocation(year int, month time.Month, day int, loc *time.Location) *Time {
	return Create(time.Date(year, month, day, 0, 0, 0, 0, loc), "")
}

func (t *Time) Format(format string) *Time {
	t.format = format
	return t
}

func (t *Time) ToDateTimeString() string {
	return t.time.Format(t.format)
}
