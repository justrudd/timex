package timex

import "time"

// FirstDayOfMonth returns a new time.Time in the same month set to the
// first day of the month. The clock of the time is not adjusted.
func FirstDayOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	h, mi, s := t.Clock()

	return time.Date(y, m, 1, h, mi, s, t.Nanosecond(), t.Location())
}

// FirstDayOfNextMonth returns a new time.Time for the first day of the
// next month. The clock of the time is not adjusted.
func FirstDayOfNextMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	h, mi, s := t.Clock()

	nm := NextMonth(m)
	ya := 0
	if nm == time.January {
		ya = 1
	}

	return time.Date(y+ya, nm, 1, h, mi, s, t.Nanosecond(), t.Location())
}

// FirstDayOfYear returns a new time.Time for first day in the current year.
// The clock of the time is not adjusted.
func FirstDayOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	h, mi, s := t.Clock()

	return time.Date(y, time.January, 1, h, mi, s, t.Nanosecond(), t.Location())
}

// FirstInMonth returns a new time.Time in the same month with the first
// matching time.Weekday. The clock of the time is not adjusted.
func FirstInMonth(t time.Time, w time.Weekday) time.Time {
	nt := FirstDayOfMonth(t)
	return NextDayOfWeek(nt, w, false)
}

// LastDayOfMonth returns a new time.Time in the same month set to the
// last day of the month. The clock of the time is not adjusted.
func LastDayOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	h, mi, s := t.Clock()

	d := DaysInMonth(y, m)
	return time.Date(y, m, d, h, mi, s, t.Nanosecond(), t.Location())
}

// LastDayOfYear returns a new time.Time for first day in the current year.
// The clock of the time is not adjusted.
func LastDayOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	h, mi, s := t.Clock()

	return time.Date(y, time.December, 31, h, mi, s, t.Nanosecond(), t.Location())
}

// LastInMonth returns a new time.Time in the same month with the last
// matching time.Weekday. The clock of the time is not adjusted.
func LastInMonth(t time.Time, w time.Weekday) time.Time {
	nt := LastDayOfMonth(t)
	return PrevDayOfWeek(nt, w, false)
}

// NextDayOfWeek returns a new time.Time for the next specified Weekday. It
// is possible for this method to move beyond current month and year.
// For example, if you are on the last Monday of December and ask for next
// Monday, this method will return the first Monday in January of the next
// year. The clock of the time is not adjusted.
//
// If `wrap` is false, asking for next Monday when time is already on Monday
// will return the same time.
func NextDayOfWeek(t time.Time, w time.Weekday, wrap bool) time.Time {
	d := DaysBetweenWeekdays(t.Weekday(), w)
	if d == 0 && wrap {
		d = 7
	}

	return t.AddDate(0, 0, d)
}

// NthDayOfWeek returns a new time.Time for the `n`th specified Weekday. It
// uses the current month and year as the starting point. It is possible to
// have a time returned outside of the current month and year if you pass in
// a big enough value for `n`.
//
// `n` can be positive to move forward in time from the beginning of the month.
// `n` can be negative to move backward in time from the end of the month.
// Passing in 0 will return `t` unchanged.
func NthDayOfWeek(t time.Time, w time.Weekday, n int) time.Time {
	if n == 0 {
		return t
	}

	var nt time.Time
	if n > 0 {
		nt = FirstDayOfMonth(t)
		if nt.Weekday() == w {
			n--
		}

		if n > 0 {
			nt = NextDayOfWeek(nt, w, true)
			n--
		}
	} else {
		nt = LastDayOfMonth(t)
		if nt.Weekday() == w {
			n++
		}

		if n < 0 {
			nt = PrevDayOfWeek(nt, w, true)
			n++
		}
	}

	nt = nt.AddDate(0, 0, 7*n)

	return nt
}

// PrevDayOfWeek returns a new time.Time for the previous specified Weekday. It
// is possible for this method to move before current month and year.
// For example, if you are on the first Monday of January and ask for the previous
// Monday, this method will return the last Monday in December of the previous
// year. The clock of the time is not adjusted.
//
// If wrap is false, asking for previous Monday when time is already on Monday
// will return the same time.
func PrevDayOfWeek(t time.Time, w time.Weekday, wrap bool) time.Time {
	d := DaysBetweenWeekdays(w, t.Weekday())
	if d == 0 && wrap {
		d = 7
	}

	return t.AddDate(0, 0, -d)
}
