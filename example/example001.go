package main

import (
	"fmt"
	"github.com/xiaohei2015/platinum"
	"time"
)

func main() {
	fmt.Println("Now:", platinum.Now())
	// prints: Now: 2018-11-02 11:07:19
	fmt.Println("Right now:", platinum.Now().ToDateTimeString())
	// prints: Right now: 2018-11-02 11:07:19
	fmt.Println("Date:", platinum.CreateFromDate(2018, 10, 31))
	// prints: Date: 2018-10-31 00:00:00
	fmt.Println("DateTime:", platinum.CreateFromDateTime(2018, 10, 31, 13, 22, 33))
	// prints: DateTime: 2018-10-31 13:22:33

	fmt.Println("Today:", platinum.Today())
	// prints: Today: 2018-11-02 00:00:00
	fmt.Println("Tomorrow:", platinum.Tomorrow())
	// prints: Tomorrow: 2018-11-03 00:00:00
	fmt.Println("Yesterday:", platinum.Yesterday())
	// prints: Yesterday: 2018-11-01 00:00:00

	fmt.Println("Location Change:", platinum.Now().Location(time.UTC))
	// prints: Location Change: 2018-11-02 03:07:19
	fmt.Println("Format Change:", platinum.CreateFromDate(2018, 10, 31).Format("2006/01/02"))
	// prints: Format Change: 2018/10/31

	fmt.Println("IsWeekend:", platinum.CreateFromDate(2018, 10, 29).IsWeekend())
	// prints: IsWeekend: false
	fmt.Println("IsWeekday:", platinum.CreateFromDate(2018, 10, 29).IsWeekday())
	// prints: IsWeekday: true
	fmt.Println("IsYesterday:", platinum.CreateFromDate(2018, 11, 1).IsYesterday())
	// prints: IsYesterday: true
	fmt.Println("IsToday:", platinum.CreateFromDate(2018, 11, 1).IsToday())
	// prints: IsToday: false
	fmt.Println("IsTomorrow:", platinum.CreateFromDate(2018, 11, 1).IsTomorrow())
	// prints: IsTomorrow: false
	fmt.Println("IsFuture:", platinum.CreateFromDate(2018, 11, 2).IsFuture())
	// prints: IsFuture: false
	fmt.Println("IsPast:", platinum.CreateFromDate(2018, 11, 2).IsPast())
	// prints: IsPast: true
	fmt.Println("IsLeapYear:", platinum.CreateFromDate(2004, 11, 2).IsLeapYear())
	// prints: IsLeapYear: true

	fmt.Println("AddYears:", platinum.CreateFromDate(2018, 11, 2).AddYears(3))
	// prints: AddYears: 2021-11-02 00:00:00
	fmt.Println("SubYears:", platinum.CreateFromDate(2018, 11, 2).SubYears(3))
	// prints: SubYears: 2015-11-02 00:00:00
	fmt.Println("AddMonths:", platinum.CreateFromDate(2018, 11, 2).AddMonths(3))
	// prints: AddMonths: 2019-02-02 00:00:00
	fmt.Println("SubMonths:", platinum.CreateFromDate(2018, 11, 2).SubMonths(3))
	// prints: SubMonths: 2018-08-02 00:00:00
	fmt.Println("AddDays:", platinum.CreateFromDate(2018, 11, 2).AddDays(3))
	// prints: AddDays: 2018-11-05 00:00:00
	fmt.Println("SubDays:", platinum.CreateFromDate(2018, 11, 2).SubDays(3))
	// prints: SubDays: 2018-10-30 00:00:00
	fmt.Println("AddWeekdays:", platinum.CreateFromDate(2018, 11, 1).AddWeekdays(-9))
	// prints: AddWeekdays: 2018-10-19 00:00:00
	fmt.Println("SubWeekdays:", platinum.CreateFromDate(2018, 11, 1).SubWeekdays(-9))
	// prints: SubWeekdays: 2018-11-14 00:00:00
	fmt.Println("AddWeeks:", platinum.CreateFromDate(2018, 11, 1).AddWeeks(1))
	// prints: AddWeeks: 2018-11-08 00:00:00
	fmt.Println("SubWeeks:", platinum.CreateFromDate(2018, 11, 1).SubWeeks(1))
	// prints: SubWeeks: 2018-10-25 00:00:00
	fmt.Println("AddHours:", platinum.CreateFromDate(2018, 11, 1).AddHours(1))
	// prints: AddHours: 2018-11-01 01:00:00
	fmt.Println("SubHours:", platinum.CreateFromDate(2018, 11, 1).SubHours(1))
	// prints: SubHours: 2018-10-31 23:00:00
	fmt.Println("AddMinutes:", platinum.CreateFromDate(2018, 11, 1).AddMinutes(1))
	// prints: AddMinutes: 2018-11-01 00:01:00
	fmt.Println("SubMinutes:", platinum.CreateFromDate(2018, 11, 1).SubMinutes(1))
	// prints: SubMinutes: 2018-10-31 23:59:00
	fmt.Println("AddSeconds:", platinum.CreateFromDate(2018, 11, 1).AddSeconds(1))
	// prints: AddSeconds: 2018-11-01 00:00:01
	fmt.Println("SubSeconds:", platinum.CreateFromDate(2018, 11, 1).SubSeconds(1))
	// prints: SubSeconds: 2018-10-31 23:59:59

	fmt.Println("DateTimeDiff:", platinum.DateTimeDiff(platinum.CreateFromDate(1984, 1, 28), platinum.CreateFromDate(1989, 1, 27)))
	// prints: DateTimeDiff: {4 11 30 0 0 0}
	fmt.Println("DiffInSeconds:", platinum.CreateFromDate(2018, 11, 3).DiffInSeconds(platinum.CreateFromDate(2018, 11, 2), false))
	// prints: DiffInSeconds: -86400
	fmt.Println("DiffInMinutes:", platinum.CreateFromDate(2018, 11, 3).DiffInMinutes(platinum.CreateFromDate(2018, 11, 2), false))
	// prints: DiffInMinutes: -1440
	fmt.Println("DiffInHours:", platinum.CreateFromDate(2018, 11, 3).DiffInHours(platinum.CreateFromDate(2018, 11, 2), false))
	// prints: DiffInHours: -24
	fmt.Println("DiffInDays:", platinum.CreateFromDate(2018, 11, 3).DiffInDays(platinum.CreateFromDate(2018, 11, 2), false))
	// prints: DiffInDays: -1
	fmt.Println("DiffInWeeks:", platinum.CreateFromDate(2018, 11, 3).DiffInWeeks(platinum.CreateFromDate(2018, 11, 2), false))
	// prints: DiffInWeeks: 0
	fmt.Println("DiffInMonths:", platinum.CreateFromDate(2017, 11, 3).DiffInMonths(platinum.CreateFromDate(2018, 11, 2), false))
	// prints: DiffInMonths: 11
	fmt.Println("DiffInYears:", platinum.CreateFromDate(2017, 11, 1).DiffInYears(platinum.CreateFromDate(2018, 11, 2), false))
	// prints: DiffInYears: 1
}
