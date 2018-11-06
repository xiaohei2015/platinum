package platinum

import (
	"fmt"
	"time"
)

type Time struct {
	time   time.Time
	format string
}

type DiffTime struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
}

const (
	DateDefaultFormat     = "2006-01-02"
	TimeDefaultFormat     = "15:04:05"
	DateTimeDefaultFormat = "2006-01-02 15:04:05"
)

const (
	YearsPerCentury    = 100
	YearsPerDecade     = 10
	MonthsPerYear      = 12
	MonthsPerQuarter   = 3
	WeeksPerYear       = 52
	WeeksperMonth      = 4
	DaysPerWeek        = 7
	WorkingDaysPerWeek = 5
	HoursPerDay        = 24
	MinutesPerHour     = 60
	SecondsPerMinute   = 60
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

func (t *Time) GetTimeStamp() int64 {
	return t.time.Unix()
}

func (t *Time) GetYear() int {
	return t.time.Year()
}

func (t *Time) GetMonth() time.Month {
	return t.time.Month()
}

func (t *Time) GetDay() int {
	return t.time.Day()
}

func (t *Time) GetHour() int {
	return t.time.Hour()
}

func (t *Time) GetMinute() int {
	return t.time.Minute()
}

func (t *Time) GetSecond() int {
	return t.time.Second()
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
	weekNum := delta / WorkingDaysPerWeek
	//how many days
	dayNum := delta % WorkingDaysPerWeek
	//how many extra weekend
	extWeekendNum := 0
	if dayNum > (WorkingDaysPerWeek - iWeekDay) {
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
	weekNum := delta / WorkingDaysPerWeek
	//how many days
	dayNum := delta % WorkingDaysPerWeek
	//how many extra weekend
	extWeekendNum := 0
	if dayNum > (iWeekDay - 1) {
		extWeekendNum++
	}
	t.time = t.time.AddDate(0, 0, -1*(weekNum*7+extWeekendNum*2+dayNum))
	return t
}

func (t *Time) AddWeeks(delta int) *Time {
	t.time = t.time.AddDate(0, 0, delta*DaysPerWeek)
	return t
}

func (t *Time) SubWeeks(delta int) *Time {
	t.time = t.time.AddDate(0, 0, -1*delta*DaysPerWeek)
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//return the max number of the unit
//typeDate == 0 (means how many months in a year), return 12
//typeDate == 1 (means how many days in a month), return value depends the month and whether a leap year
//typeDate == 2 (means how many hours in a day), return 24
//typeDate == 3 (means how many minutes in a hour), return 60
//typeDate == 4 (means how many seconds in a minute), return 60
func getTimeUnit(typeDate int, dataTime ...*Time) int {
	result := 0
	switch typeDate {
	case 0:
		result = 12
		break
	case 1:
		y := Now()
		if dataTime != nil {
			y = dataTime[0]
		}
		month := y.GetMonth()
		if month == time.January ||
			month == time.March ||
			month == time.May ||
			month == time.July ||
			month == time.August ||
			month == time.October ||
			month == time.December {
			result = 31
		} else if month == time.April ||
			month == time.June ||
			month == time.September ||
			month == time.November {
			result = 30
		} else if month == time.February {
			if y.IsLeapYear() {
				result = 29
			} else {
				result = 28
			}
		}
		break
	case 2:
		result = 24
		break
	case 3:
		result = 60
		break
	case 4:
		result = 60
		break
	}
	return result
}

// borrow time from more high unit
// typeDate == 4 , borrow from minute to second
// typeDate == 3 , borrow from hour to minute
// typeDate == 2 , borrow from day to hour
// typeDate == 1 , borrow from month to day
// typeDate == 0 , borrow from year to month
func borrowTime(dataDate [6]int, typeDate int) [6]int {
	if (typeDate == 4 && dataDate[4] == 0) ||
		(typeDate == 3 && dataDate[3] == 0) ||
		(typeDate == 2 && dataDate[2] == 1) ||
		(typeDate == 1 && dataDate[1] == 1) {
		borrowTime(dataDate, typeDate-1)
	}
	dataDate[typeDate] = dataDate[typeDate] - 1
	dataDate[typeDate+1] = dataDate[typeDate+1] + getTimeUnit(typeDate, CreateFromDate(dataDate[0], time.Month(dataDate[1]), dataDate[2]))
	return dataDate
}

func DateTimeDiff(sourceTime, targetTime *Time) DiffTime {
	if sourceTime.GetTimeStamp() > targetTime.GetTimeStamp() {
		sourceTime, targetTime = targetTime, sourceTime
	}
	sourceYear, sourceMonth, sourceDay := sourceTime.time.Date()
	sourceHour, sourceMinute, sourceSecond := sourceTime.time.Clock()
	targetYear, targetMonth, targetDay := targetTime.time.Date()
	targetHour, targetMinute, targetSecond := targetTime.time.Clock()
	var diffTime DiffTime
	sourceDate := [6]int{sourceYear, int(sourceMonth), sourceDay, sourceHour, sourceMinute, sourceSecond}
	targetDate := [6]int{targetYear, int(targetMonth), targetDay, targetHour, targetMinute, targetSecond}
	if sourceDate[5] > targetDate[5] {
		targetDate = borrowTime(targetDate, 4)
	}
	diffTime.second = targetDate[5] - sourceDate[5]
	if sourceDate[4] > targetDate[4] {
		targetDate = borrowTime(targetDate, 3)
	}
	diffTime.minute = targetDate[4] - sourceDate[4]
	if sourceDate[3] > targetDate[3] {
		targetDate = borrowTime(targetDate, 2)
	}
	diffTime.hour = targetDate[3] - sourceDate[3]
	if sourceDate[2] > targetDate[2] {
		targetDate = borrowTime(targetDate, 1)
	}
	diffTime.day = targetDate[2] - sourceDate[2]
	if sourceDate[1] > targetDate[1] {
		targetDate = borrowTime(targetDate, 0)
	}
	diffTime.month = targetDate[1] - sourceDate[1]
	diffTime.year = targetDate[0] - sourceDate[0]
	return diffTime
}

//Get the difference in seconds, uniform to be UTC and then compare
//func (t *Time) DiffInSeconds(targetTime *Time, absolute ...bool) int64 {
//	isAbsolute := true
//	if absolute != nil {
//		isAbsolute = absolute[0]
//	}
//	newSourceTime, _ := time.Parse(t.format, t.ToDateTimeString())
//	newTargetTime, _ := time.Parse(targetTime.format, targetTime.ToDateTimeString())
//	result := newTargetTime.Unix() - newSourceTime.Unix()
//	if isAbsolute && result < 0 {
//		result = -1 * result
//	}
//	return result
//}

//Get the difference in seconds using timestamps
func (t *Time) DiffInSeconds(targetTime *Time, absolute ...bool) int64 {
	isAbsolute := true
	if absolute != nil {
		isAbsolute = absolute[0]
	}
	result := targetTime.GetTimeStamp() - t.GetTimeStamp()
	if isAbsolute && result < 0 {
		result = -1 * result
	}
	return result
}

func (t *Time) DiffInMinutes(targetTime *Time, absolute ...bool) int64 {
	isAbsolute := true
	if absolute != nil {
		isAbsolute = absolute[0]
	}
	diffSecs := t.DiffInSeconds(targetTime, isAbsolute)
	return diffSecs / SecondsPerMinute
}

func (t *Time) DiffInHours(targetTime *Time, absolute ...bool) int64 {
	isAbsolute := true
	if absolute != nil {
		isAbsolute = absolute[0]
	}
	diffSecs := t.DiffInSeconds(targetTime, isAbsolute)
	return diffSecs / SecondsPerMinute / MinutesPerHour
}

func (t *Time) DiffInDays(targetTime *Time, absolute ...bool) int64 {
	isAbsolute := true
	if absolute != nil {
		isAbsolute = absolute[0]
	}
	diffSecs := t.DiffInSeconds(targetTime, isAbsolute)
	return diffSecs / HoursPerDay / SecondsPerMinute / MinutesPerHour
}

func (t *Time) DiffInWeeks(targetTime *Time, absolute ...bool) int64 {
	isAbsolute := true
	if absolute != nil {
		isAbsolute = absolute[0]
	}
	diffSecs := t.DiffInSeconds(targetTime, isAbsolute)
	return diffSecs / DaysPerWeek / HoursPerDay / SecondsPerMinute / MinutesPerHour
}

func (t *Time) DiffInMonths(targetTime *Time, absolute ...bool) int {
	diffTime := DateTimeDiff(t, targetTime)
	return MonthsPerYear*diffTime.year + diffTime.month
}

func (t *Time) DiffInYears(targetTime *Time, absolute ...bool) int {
	diffTime := DateTimeDiff(t, targetTime)
	return diffTime.year
}
