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

func CreateFromDate(year int, month time.Month, day int) *Time {
	return Create(time.Date(year, month, day, 0, 0, 0, 0, time.Local), "")
}

func CreateFromDateTime(year int, month time.Month, day int, hour int, minute int, second int) *Time {
	return Create(time.Date(year, month, day, hour, minute, second, 0, time.Local), "")
}

func Now() *Time {
	return Create(time.Now(), "")
}

func Today() *Time {
	t := time.Now()
	year, month, day := t.Date()
	return CreateFromDate(year, month, day)
}

func Tomorrow() *Time {
	t := time.Now().AddDate(0, 0, 1)
	year, month, day := t.Date()
	return CreateFromDate(year, month, day)
}

func Yesterday() *Time {
	t := time.Now().AddDate(0, 0, -1)
	year, month, day := t.Date()
	return CreateFromDate(year, month, day)
}

/**
 * Update time format
 */
func (t *Time) Format(format string) *Time {
	t.format = format
	return t
}

/**
 * Update time locaiton
 */
func (t *Time) Location(loc *time.Location) *Time {
	return Create(t.time.In(loc), "")
}

func (t *Time) ToDateTimeString() string {
	return t.time.Format(t.format)
}
