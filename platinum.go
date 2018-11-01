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

func Now() *Time {
	return &Time{
		time:   time.Now(),
		format: "2006-01-02 15:04:05",
	}
}

func (t *Time) ToDateTimeString() string {
	return t.time.Format(t.format)
}
