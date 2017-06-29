package timex

import "time"

var (
	daysInMonth = [13]int{
		-1, // time.Month is 1 based
		31, // time.January
		28, // time.February
		31, // time.March
		30, // time.April
		31, // time.May
		30, // time.June
		31, // time.July
		31, // time.August
		30, // time.September
		31, // time.October
		30, // time.November
		31, // time.December
	}
)

// DaysInMonth returns the number of days in the month for the year.
func DaysInMonth(y int, m time.Month) int {
	if m == time.February && IsLeapYear(y) {
		return 29
	}
	return daysInMonth[m]
}

// DaysBetweenWeekdays returns the number of days between two
// time.Weekday values. The order of parameters does matter. If you
// pass `w1 == time.Tuesday` and `w2 == time.Wednesday`, you'll get
// a result of 1. If you reverse them (`w1 == time.Wednesday` and
// `w2 == time.Tuesday`), you'll get a result of 6.
func DaysBetweenWeekdays(w1, w2 time.Weekday) (dist int) {
	if w1 > w2 {
		// Wed to Tue case. The +1 is because time.Weekday is 0 based
		dist = int(time.Saturday - w1 + w2 + 1)
	} else {
		dist = int(w2 - w1)
	}
	return dist
}

// IsLeapYear returns whether the year is a leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// NextBusinessWeekday returns the next business weekday after the
// current weekday. It assumes a standard work week of Monday
// through Friday. It will wrap from Friday, Saturday, or Sunday
// to Monday.
func NextBusinessWeekday(w time.Weekday) time.Weekday {
	var nw time.Weekday
	switch w {
	case time.Friday, time.Saturday, time.Sunday:
		nw = time.Monday
	default:
		nw = w + 1
	}
	return nw
}

// NextMonth returns the next month after the current month. It will wrap
// from December to January.
func NextMonth(m time.Month) time.Month {
	var nm time.Month
	switch m {
	case time.December:
		nm = time.January
	default:
		nm = m + 1
	}
	return nm
}

// NextWeekday returns the next weekday after the current weekday. It
// will wrap from Saturday to Sunday.
func NextWeekday(w time.Weekday) time.Weekday {
	var nw time.Weekday
	switch w {
	case time.Saturday:
		nw = time.Sunday
	default:
		nw = w + 1
	}
	return nw
}

// PrevBusinessWeekday returns the previous business weekday before the
// current weekday. It assumes a standard work week of Monday
// through Friday. It will wrap from Saturday, Sunday, or Monday to
// Friday.
func PrevBusinessWeekday(w time.Weekday) time.Weekday {
	var pw time.Weekday
	switch w {
	case time.Saturday, time.Sunday, time.Monday:
		pw = time.Friday
	default:
		pw = w - 1
	}
	return pw
}

// PrevMonth returns the previous month before the current month. It will wrap
// from January to December.
func PrevMonth(m time.Month) time.Month {
	var pm time.Month
	switch m {
	case time.January:
		pm = time.December
	default:
		pm = m - 1
	}
	return pm
}

// PrevWeekday returns the previous weekday after the current weekday. It
// will wrap from Sunday to Saturday.
func PrevWeekday(w time.Weekday) time.Weekday {
	var pw time.Weekday
	switch w {
	case time.Sunday:
		pw = time.Saturday
	default:
		pw = w - 1
	}
	return pw
}
