package timex

import "time"

// BeginningOfDay returns a new time.Time with the current date but
// at the beginning of the day. The timezone is not modified.
func BeginningOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns a new time.Time with the current date but at the
// end of the day with 1 second remaining in the day. The timezone
// is not modified.
func EndOfDay(t time.Time) time.Time {
	nt := t.AddDate(0, 0, 1)
	nt = BeginningOfDay(nt)
	nt = nt.Add(-time.Second)

	return nt
}
