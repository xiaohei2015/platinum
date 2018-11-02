package platinum

import (
	"fmt"
	"time"
)

type Time struct {
	time   time.Time
	format string
}

const (
	DateDefaultFormat     = "2006-01-02"
	TimeDefaultFormat     = "15:04:05"
	DateTimeDefaultFormat = "2006-01-02 15:04:05"
)

const (
	WeekWorkDayNum = 5
	WeekDayNum     = 7
)

var WeekDayMap = map[string]int{
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
	"Sunday":    0,
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (t *Time) String() string {
	return t.time.Format(t.format)
}

func Create(tm time.Time, format string) *Time {
	if format == "" {
		format = DateTimeDefaultFormat
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Update time format
func (t *Time) Format(format string) *Time {
	t.format = format
	return t
}

//Update time location
func (t *Time) Location(loc *time.Location) *Time {
	return Create(t.time.In(loc), "")
}

func (t *Time) ToDateString() string {
	return t.time.Format(DateDefaultFormat)
}

func (t *Time) ToDateTimeString() string {
	return t.time.Format(t.format)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (t *Time) IsWeekend() bool {
	weekday := t.time.Weekday().String()
	if weekday == "Sunday" || weekday == "Saturday" {
		return true
	} else {
		return false
	}
}

func (t *Time) IsWeekday() bool {
	return !t.IsWeekend()
}

func (t *Time) IsYesterday() bool {
	return t.ToDateString() == Yesterday().ToDateString()
}

func (t *Time) IsToday() bool {
	return t.ToDateString() == Today().ToDateString()
}

func (t *Time) IsTomorrow() bool {
	return t.ToDateString() == Tomorrow().ToDateString()
}

func (t *Time) IsFuture() bool {
	return t.time.Unix() > time.Now().Unix()
}

func (t *Time) IsPast() bool {
	return t.time.Unix() < time.Now().Unix()
}

func (t *Time) IsLeapYear() bool {
	year := t.time.Year()
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	} else {
		return false
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (t *Time) AddYears(delta int) *Time {
	t.time = t.time.AddDate(delta, 0, 0)
	return t
}

func (t *Time) SubYears(delta int) *Time {
	t.time = t.time.AddDate(-1*delta, 0, 0)
	return t
}

func (t *Time) AddMonths(delta int) *Time {
	t.time = t.time.AddDate(0, delta, 0)
	return t
}

func (t *Time) SubMonths(delta int) *Time {
	t.time = t.time.AddDate(0, -1*delta, 0)
	return t
}

func (t *Time) AddDays(delta int) *Time {
	t.time = t.time.AddDate(0, 0, delta)
	return t
}

func (t *Time) SubDays(delta int) *Time {
	t.time = t.time.AddDate(0, 0, -1*delta)
	return t
}

func (t *Time) AddWeekdays(delta int) *Time {
	if delta < 0 {
		return t.SubWeekdays(-1 * delta)
	}
	iWeekDay := WeekDayMap[t.time.Weekday().String()]
	//how many weeks
	weekNum := delta / WeekWorkDayNum
	//how many days
	dayNum := delta % WeekWorkDayNum
	//how many extra weekend
	extWeekendNum := 0
	if dayNum > (WeekWorkDayNum - iWeekDay) {
		extWeekendNum++
	}
	t.time = t.time.AddDate(0, 0, weekNum*7+extWeekendNum*2+dayNum)
	return t
}

func (t *Time) SubWeekdays(delta int) *Time {
	if delta < 0 {
		return t.AddWeekdays(-1 * delta)
	}
	iWeekDay := WeekDayMap[t.time.Weekday().String()]
	//how many weeks
	weekNum := delta / WeekWorkDayNum
	//how many days
	dayNum := delta % WeekWorkDayNum
	//how many extra weekend
	extWeekendNum := 0
	if dayNum > (iWeekDay - 1) {
		extWeekendNum++
	}
	t.time = t.time.AddDate(0, 0, -1*(weekNum*7+extWeekendNum*2+dayNum))
	return t
}

func (t *Time) AddWeeks(delta int) *Time {
	t.time = t.time.AddDate(0, 0, delta*WeekDayNum)
	return t
}

func (t *Time) SubWeeks(delta int) *Time {
	t.time = t.time.AddDate(0, 0, -1*delta*WeekDayNum)
	return t
}

func (t *Time) AddHours(delta int) *Time {
	dur, _ := time.ParseDuration(fmt.Sprintf("%dh", delta))
	t.time = t.time.Add(dur)
	return t
}

func (t *Time) SubHours(delta int) *Time {
	dur, _ := time.ParseDuration(fmt.Sprintf("%dh", -1*delta))
	t.time = t.time.Add(dur)
	return t
}

func (t *Time) AddMinutes(delta int) *Time {
	dur, _ := time.ParseDuration(fmt.Sprintf("%dm", delta))
	t.time = t.time.Add(dur)
	return t
}

func (t *Time) SubMinutes(delta int) *Time {
	dur, _ := time.ParseDuration(fmt.Sprintf("%dm", -1*delta))
	t.time = t.time.Add(dur)
	return t
}

func (t *Time) AddSeconds(delta int) *Time {
	dur, _ := time.ParseDuration(fmt.Sprintf("%ds", delta))
	t.time = t.time.Add(dur)
	return t
}

func (t *Time) SubSeconds(delta int) *Time {
	dur, _ := time.ParseDuration(fmt.Sprintf("%ds", -1*delta))
	t.time = t.time.Add(dur)
	return t
}
