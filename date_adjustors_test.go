package timex_test

import (
	"testing"
	"time"

	. "github.com/justrudd/timex"
)

const (
	wrap   = true
	noWrap = false
)

var (
	local, _ = time.LoadLocation("Local")
	utc, _   = time.LoadLocation("UTC")
	nyc, _   = time.LoadLocation("America/New_York")
)

func TestFirstDayOfMonth(t *testing.T) {
	cases := []struct {
		t, expected time.Time
	}{
		{
			time.Date(2015, time.July, 27, 15, 41, 0, 0, local),
			time.Date(2015, time.July, 1, 15, 41, 0, 0, local),
		},
		{
			time.Date(2017, time.November, 30, 9, 15, 56, 0, utc),
			time.Date(2017, time.November, 1, 9, 15, 56, 0, utc),
		},
		{
			time.Date(2012, time.January, 19, 0, 1, 34, 0, nyc),
			time.Date(2012, time.January, 1, 0, 1, 34, 0, nyc),
		},
	}

	for _, c := range cases {
		got := FirstDayOfMonth(c.t)
		if got != c.expected {
			t.Errorf("FirstDayOfMonth(%v) == %v, want %v", c.t, got, c.expected)
		}
	}
}

func TestFirstDayOfNextMonth(t *testing.T) {
	cases := []struct {
		t, expected time.Time
	}{
		{
			time.Date(2015, time.July, 27, 15, 41, 0, 0, local),
			time.Date(2015, time.August, 1, 15, 41, 0, 0, local),
		},
		{
			time.Date(2017, time.November, 30, 9, 15, 56, 0, utc),
			time.Date(2017, time.December, 1, 9, 15, 56, 0, utc),
		},
		// make sure it wraps
		{
			time.Date(2012, time.December, 19, 0, 1, 34, 0, nyc),
			time.Date(2013, time.January, 1, 0, 1, 34, 0, nyc),
		},
	}

	for _, c := range cases {
		got := FirstDayOfNextMonth(c.t)
		if got != c.expected {
			t.Errorf("FirstDayOfNextMonth(%v) == %v, want %v", c.t, got, c.expected)
		}
	}
}

func TestFirstDayOfYear(t *testing.T) {
	cases := []struct {
		t, expected time.Time
	}{
		{
			time.Date(2015, time.July, 27, 15, 41, 0, 0, local),
			time.Date(2015, time.January, 1, 15, 41, 0, 0, local),
		},
		{
			time.Date(2017, time.November, 30, 9, 15, 56, 0, utc),
			time.Date(2017, time.January, 1, 9, 15, 56, 0, utc),
		},
		{
			time.Date(2012, time.December, 19, 0, 1, 34, 0, nyc),
			time.Date(2012, time.January, 1, 0, 1, 34, 0, nyc),
		},
	}

	for _, c := range cases {
		got := FirstDayOfYear(c.t)
		if got != c.expected {
			t.Errorf("FirstDayOfYear(%v) == %v, want %v", c.t, got, c.expected)
		}
	}
}

func TestFirstInMonth(t *testing.T) {
	cases := []struct {
		t        time.Time
		w        time.Weekday
		expected time.Time
	}{
		{
			time.Date(2015, time.July, 27, 15, 41, 0, 0, local),
			time.Wednesday,
			time.Date(2015, time.July, 1, 15, 41, 0, 0, local),
		},
		{
			time.Date(2015, time.July, 27, 1, 11, 56, 0, local),
			time.Monday,
			time.Date(2015, time.July, 6, 1, 11, 56, 0, local),
		},
	}

	for _, c := range cases {
		got := FirstInMonth(c.t, c.w)
		if got != c.expected {
			t.Errorf("FirstInMonth(%v, %s) == %v, want %v", c.t, c.w, got, c.expected)
		}
	}
}

func TestLastDayOfMonth(t *testing.T) {
	cases := []struct {
		t, expected time.Time
	}{
		// leap year cases first
		{
			time.Date(2015, time.February, 2, 15, 41, 0, 0, local),
			time.Date(2015, time.February, 28, 15, 41, 0, 0, local),
		},
		{
			time.Date(2016, time.February, 3, 9, 15, 56, 0, utc),
			time.Date(2016, time.February, 29, 9, 15, 56, 0, utc),
		},
		{
			time.Date(2012, time.December, 19, 0, 1, 34, 0, nyc),
			time.Date(2012, time.December, 31, 0, 1, 34, 0, nyc),
		},
	}

	for _, c := range cases {
		got := LastDayOfMonth(c.t)
		if got != c.expected {
			t.Errorf("LastDayOfMonth(%v) == %v, want %v", c.t, got, c.expected)
		}
	}
}

func TestLastDayOfYear(t *testing.T) {
	cases := []struct {
		t, expected time.Time
	}{
		{
			time.Date(2015, time.July, 27, 15, 41, 0, 0, local),
			time.Date(2015, time.December, 31, 15, 41, 0, 0, local),
		},
		{
			time.Date(2017, time.November, 30, 9, 15, 56, 0, utc),
			time.Date(2017, time.December, 31, 9, 15, 56, 0, utc),
		},
		{
			time.Date(2012, time.December, 19, 0, 1, 34, 0, nyc),
			time.Date(2012, time.December, 31, 0, 1, 34, 0, nyc),
		},
	}

	for _, c := range cases {
		got := LastDayOfYear(c.t)
		if got != c.expected {
			t.Errorf("LastDayOfYear(%v) == %v, want %v", c.t, got, c.expected)
		}
	}
}

func TestLastInMonth(t *testing.T) {
	cases := []struct {
		t        time.Time
		w        time.Weekday
		expected time.Time
	}{
		{
			time.Date(2015, time.March, 27, 15, 41, 0, 0, nyc),
			time.Tuesday,
			time.Date(2015, time.March, 31, 15, 41, 0, 0, nyc),
		},
		{
			time.Date(2015, time.May, 7, 1, 11, 56, 0, utc),
			time.Monday,
			time.Date(2015, time.May, 25, 1, 11, 56, 0, utc),
		},
	}

	for _, c := range cases {
		got := LastInMonth(c.t, c.w)
		if got != c.expected {
			t.Errorf("LastInMonth(%v, %s) == %v, want %v", c.t, c.w, got, c.expected)
		}
	}
}

func TestNextDayOfWeek(t *testing.T) {

	// this first set of tests will just move from Sunday to the
	// next specified day. Wrapping will be tested below
	sunday := time.Date(2015, time.February, 14, 12, 0, 0, 0, local)
	monday := sunday.AddDate(0, 0, 1)
	tuesday := monday.AddDate(0, 0, 1)
	wednesday := tuesday.AddDate(0, 0, 1)
	thursday := wednesday.AddDate(0, 0, 1)
	friday := thursday.AddDate(0, 0, 1)
	saturday := friday.AddDate(0, 0, 1)
	nextSunday := saturday.AddDate(0, 0, 1)

	runTest := func(toTest time.Time, wrap bool, expected time.Time) {
		d := expected.Weekday()
		got := NextDayOfWeek(toTest, d, wrap)
		if got != expected {
			t.Errorf("NextDayOfWeek(%v, %s, %t) == %v, want %v", toTest, d, wrap, got, expected)
		}
	}

	runTest(sunday, noWrap, monday)
	runTest(sunday, wrap, monday)
	runTest(sunday, noWrap, tuesday)
	runTest(sunday, wrap, tuesday)
	runTest(sunday, noWrap, wednesday)
	runTest(sunday, wrap, wednesday)
	runTest(sunday, noWrap, thursday)
	runTest(sunday, wrap, thursday)
	runTest(sunday, noWrap, friday)
	runTest(sunday, wrap, friday)
	runTest(sunday, noWrap, saturday)
	runTest(sunday, wrap, saturday)
	runTest(sunday, noWrap, sunday)   // without wrapping should be same as what we pass in
	runTest(sunday, wrap, nextSunday) // with wrapping should be a week later

	// now we test if we wrap (i.e. Wednesday to next Tuesday)
	tuesday = wednesday.AddDate(0, 0, 6)
	runTest(wednesday, noWrap, tuesday)
	runTest(wednesday, wrap, tuesday)
}

func TestNthDayOfWeek(t *testing.T) {
	// testing this function with a smattering of USA holidays

	cases := []struct {
		start    time.Time
		expected time.Time
		weekday  time.Weekday
		n        int
	}{
		// Martin Luther King day - 3rd Monday in January
		{
			time.Date(2015, time.January, 1, 0, 0, 0, 0, local),
			time.Date(2015, time.January, 19, 0, 0, 0, 0, local),
			time.Monday,
			3,
		},
		// Memorial Day - last Monday in May so we use -1
		{
			time.Date(2015, time.May, 1, 0, 0, 0, 0, nyc),
			time.Date(2015, time.May, 25, 0, 0, 0, 0, nyc),
			time.Monday,
			-1,
		},
		// Columbus Day - 2nd Monday in October
		{
			time.Date(2015, time.October, 1, 0, 0, 0, 0, nyc),
			time.Date(2015, time.October, 12, 0, 0, 0, 0, nyc),
			time.Monday,
			2,
		},
		// Thanksgiving Day - 4th Thursday in November
		// Using 2017 as there are 5 Thursdays in it
		{
			time.Date(2017, time.November, 1, 0, 0, 0, 0, utc),
			time.Date(2017, time.November, 23, 0, 0, 0, 0, utc),
			time.Thursday,
			4,
		},
		// random July Wednesday...
		{
			time.Date(2015, time.July, 23, 12, 3, 34, 0, local),
			time.Date(2015, time.July, 8, 12, 3, 34, 0, local),
			time.Wednesday,
			2,
		},
		// another random July Friday
		{
			time.Date(2015, time.July, 2, 2, 13, 4, 0, utc),
			time.Date(2015, time.July, 24, 2, 13, 4, 0, utc),
			time.Friday,
			-2,
		},
		// ensure that 0 returns the same date
		{
			time.Date(2017, time.November, 7, 23, 1, 59, 0, nyc),
			time.Date(2017, time.November, 7, 23, 1, 59, 0, nyc),
			time.Thursday,
			0,
		},
	}

	for _, c := range cases {
		got := NthDayOfWeek(c.start, c.weekday, c.n)
		if got != c.expected {
			t.Errorf("NthDayOfWeek(%s, %s, %d) == %s, want %s", c.start, c.weekday, c.n, got, c.expected)
		}
	}
}

func TestPrevDayOfWeek(t *testing.T) {

	// this first set of tests will just move from Sunday to the
	// previous specified day. Wrapping will be tested below
	sunday := time.Date(2015, time.February, 14, 12, 0, 0, 0, local)
	saturday := sunday.AddDate(0, 0, -1)
	friday := saturday.AddDate(0, 0, -1)
	thursday := friday.AddDate(0, 0, -1)
	wednesday := thursday.AddDate(0, 0, -1)
	tuesday := wednesday.AddDate(0, 0, -1)
	monday := tuesday.AddDate(0, 0, -1)
	prevSunday := monday.AddDate(0, 0, -1)

	runTest := func(toTest time.Time, wrap bool, expected time.Time) {
		d := expected.Weekday()
		got := PrevDayOfWeek(toTest, d, wrap)
		if got != expected {
			t.Errorf("PrevDayOfWeek(%v, %s, %t) == %v, want %v", toTest, d, wrap, got, expected)
		}
	}

	runTest(sunday, noWrap, monday)
	runTest(sunday, wrap, monday)
	runTest(sunday, noWrap, tuesday)
	runTest(sunday, wrap, tuesday)
	runTest(sunday, noWrap, wednesday)
	runTest(sunday, wrap, wednesday)
	runTest(sunday, noWrap, thursday)
	runTest(sunday, wrap, thursday)
	runTest(sunday, noWrap, friday)
	runTest(sunday, wrap, friday)
	runTest(sunday, noWrap, saturday)
	runTest(sunday, wrap, saturday)
	runTest(sunday, noWrap, sunday)   // without wrapping should be same as what we pass in
	runTest(sunday, wrap, prevSunday) // with wrapping should be a week later

	// now we test if we wrap (i.e. Tuesday to prev Thursday)
	thursday = tuesday.AddDate(0, 0, -5)
	runTest(tuesday, noWrap, thursday)
	runTest(tuesday, wrap, thursday)
}
