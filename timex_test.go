package timex_test

import (
	"testing"
	"time"

	. "github.com/justrudd/timex"
)

const (
	nonLeapYear = 2015
	leapYear    = 2016
)

func TestDaysInMonth(t *testing.T) {
	cases := []struct {
		year     int
		month    time.Month
		expected int
	}{
		{nonLeapYear, time.January, 31},
		{nonLeapYear, time.February, 28},
		{nonLeapYear, time.March, 31},
		{nonLeapYear, time.April, 30},
		{nonLeapYear, time.May, 31},
		{nonLeapYear, time.June, 30},
		{nonLeapYear, time.July, 31},
		{nonLeapYear, time.August, 31},
		{nonLeapYear, time.September, 30},
		{nonLeapYear, time.October, 31},
		{nonLeapYear, time.November, 30},
		{nonLeapYear, time.December, 31},

		{leapYear, time.January, 31},
		{leapYear, time.February, 29},
		{leapYear, time.March, 31},
		{leapYear, time.April, 30},
		{leapYear, time.May, 31},
		{leapYear, time.June, 30},
		{leapYear, time.July, 31},
		{leapYear, time.August, 31},
		{leapYear, time.September, 30},
		{leapYear, time.October, 31},
		{leapYear, time.November, 30},
		{leapYear, time.December, 31},
	}

	for _, c := range cases {
		got := DaysInMonth(c.year, c.month)
		if got != c.expected {
			t.Errorf("DaysInMonth(%d, %s) == %d, want %d", c.year, c.month, got, c.expected)
		}
	}
}

func TestDaysBetweenWeekdays(t *testing.T) {
	cases := []struct {
		w1, w2   time.Weekday
		expected int
	}{
		{time.Sunday, time.Sunday, 0},
		{time.Sunday, time.Monday, 1},
		{time.Sunday, time.Tuesday, 2},
		{time.Sunday, time.Wednesday, 3},
		{time.Sunday, time.Thursday, 4},
		{time.Sunday, time.Friday, 5},
		{time.Sunday, time.Saturday, 6},

		// make sure we wrap properly
		{time.Wednesday, time.Tuesday, 6},
		{time.Thursday, time.Monday, 4},
		{time.Saturday, time.Sunday, 1},
	}

	for _, c := range cases {
		got := DaysBetweenWeekdays(c.w1, c.w2)
		if got != c.expected {
			t.Errorf("DaysBetweenWeekdays(%s, %s) == %v, want %v", c.w1, c.w2, got, c.expected)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	cases := []struct {
		year     int
		expected bool
	}{
		{1900, false},
		{2000, true},
		{2012, true},
		{2100, false},
	}

	for _, c := range cases {
		got := IsLeapYear(c.year)
		if got != c.expected {
			t.Errorf("IsLeapYear(%d) == %t, want %t", c.year, got, c.expected)
		}
	}
}

func TestNextBusinessWeekday(t *testing.T) {
	cases := []struct {
		w        time.Weekday
		expected time.Weekday
	}{
		{time.Sunday, time.Monday},
		{time.Monday, time.Tuesday},
		{time.Tuesday, time.Wednesday},
		{time.Wednesday, time.Thursday},
		{time.Thursday, time.Friday},
		{time.Friday, time.Monday},
		{time.Saturday, time.Monday},
	}

	for _, c := range cases {
		got := NextBusinessWeekday(c.w)
		if got != c.expected {
			t.Errorf("NextBusinessWeekday(%s) == %s, want %s", c.w, got, c.expected)
		}
	}
}

func TestNextMonth(t *testing.T) {
	cases := []struct {
		month    time.Month
		expected time.Month
	}{
		{time.January, time.February},
		{time.February, time.March},
		{time.March, time.April},
		{time.April, time.May},
		{time.May, time.June},
		{time.June, time.July},
		{time.July, time.August},
		{time.August, time.September},
		{time.September, time.October},
		{time.October, time.November},

		// ensure it wraps around
		{time.December, time.January},
	}

	for _, c := range cases {
		got := NextMonth(c.month)
		if got != c.expected {
			t.Errorf("NextMonth(%v) == %v, want %v", c.month, got, c.expected)
		}
	}
}

func TestNextWeekday(t *testing.T) {
	cases := []struct {
		weekday  time.Weekday
		expected time.Weekday
	}{
		{time.Sunday, time.Monday},
		{time.Monday, time.Tuesday},
		{time.Tuesday, time.Wednesday},
		{time.Wednesday, time.Thursday},
		{time.Thursday, time.Friday},
		{time.Friday, time.Saturday},

		// ensure it wraps around
		{time.Saturday, time.Sunday},
	}

	for _, c := range cases {
		got := NextWeekday(c.weekday)
		if got != c.expected {
			t.Errorf("NextWeekday(%v) == %v, want %v", c.weekday, got, c.expected)
		}
	}
}

func TestPrevBusinessWeekday(t *testing.T) {
	cases := []struct {
		w        time.Weekday
		expected time.Weekday
	}{
		{time.Sunday, time.Friday},
		{time.Monday, time.Friday},
		{time.Tuesday, time.Monday},
		{time.Wednesday, time.Tuesday},
		{time.Thursday, time.Wednesday},
		{time.Friday, time.Thursday},
		{time.Saturday, time.Friday},
	}

	for _, c := range cases {
		got := PrevBusinessWeekday(c.w)
		if got != c.expected {
			t.Errorf("PrevBusinessWeekday(%s) == %s, want %s", c.w, got, c.expected)
		}
	}
}

func TestPrevMonth(t *testing.T) {
	cases := []struct {
		month    time.Month
		expected time.Month
	}{
		// ensure it wraps around
		{time.January, time.December},

		{time.February, time.January},
		{time.March, time.February},
		{time.April, time.March},
		{time.May, time.April},
		{time.June, time.May},
		{time.July, time.June},
		{time.August, time.July},
		{time.September, time.August},
		{time.October, time.September},
		{time.December, time.November},
	}

	for _, c := range cases {
		got := PrevMonth(c.month)
		if got != c.expected {
			t.Errorf("PrevMonth(%v) == %v, want %v", c.month, got, c.expected)
		}
	}
}

func TestPrevWeekday(t *testing.T) {
	cases := []struct {
		weekday  time.Weekday
		expected time.Weekday
	}{
		// ensure it wraps
		{time.Sunday, time.Saturday},

		{time.Monday, time.Sunday},
		{time.Tuesday, time.Monday},
		{time.Wednesday, time.Tuesday},
		{time.Thursday, time.Wednesday},
		{time.Friday, time.Thursday},
		{time.Saturday, time.Friday},
	}

	for _, c := range cases {
		got := PrevWeekday(c.weekday)
		if got != c.expected {
			t.Errorf("PrevWeekday(%v) == %v, want %v", c.weekday, got, c.expected)
		}
	}
}
