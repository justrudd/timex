package timex_test

import (
	"testing"
	"time"

	. "github.com/justrudd/timex"
)

func TestBeginningOfDay(t *testing.T) {
	// we use different time zones to ensure that they don't get messed with
	tz1, _ := time.LoadLocation("UTC")
	tz2, _ := time.LoadLocation("Local")
	tz3, _ := time.LoadLocation("America/Chicago")

	t1 := time.Date(2015, time.January, 12, 12, 9, 34, 0, tz1)
	et1 := time.Date(2015, time.January, 12, 0, 0, 0, 0, tz1)
	g1 := BeginningOfDay(t1)
	if g1 != et1 {
		t.Errorf("BeginningOfDay(%v) == %v, wanted %v", t1, g1, et1)
	}

	t2 := time.Date(1999, time.July, 5, 22, 19, 3, 0, tz2)
	et2 := time.Date(1999, time.July, 5, 0, 0, 0, 0, tz2)
	g2 := BeginningOfDay(t2)
	if g2 != et2 {
		t.Errorf("BeginningOfDay(%v) == %v, wanted %v", t2, g2, et2)
	}

	t3 := time.Date(2034, time.December, 31, 23, 59, 59, 0, tz3)
	et3 := time.Date(2034, time.December, 31, 0, 0, 0, 0, tz3)
	g3 := BeginningOfDay(t3)
	if g3 != et3 {
		t.Errorf("BeginningOfDay(%v) == %v, wanted %v", t3, g3, et3)
	}
}

func TestEndOfDay(t *testing.T) {
	// we use different time zones to ensure that they don't get messed with
	tz1, _ := time.LoadLocation("UTC")
	tz2, _ := time.LoadLocation("Local")
	tz3, _ := time.LoadLocation("America/Los_Angeles")

	t1 := time.Date(2015, time.January, 12, 15, 9, 34, 0, tz1)
	et1 := time.Date(2015, time.January, 12, 23, 59, 59, 0, tz1)
	g1 := EndOfDay(t1)
	if g1 != et1 {
		t.Errorf("EndOfDay(%v) == %v, wanted %v", t1, g1, et1)
	}

	t2 := time.Date(1999, time.July, 5, 0, 19, 3, 0, tz2)
	et2 := time.Date(1999, time.July, 5, 23, 59, 59, 0, tz2)
	g2 := EndOfDay(t2)
	if g2 != et2 {
		t.Errorf("EndOfDay(%v) == %v, wanted %v", t2, g2, et2)
	}

	t3 := time.Date(2034, time.December, 31, 23, 59, 59, 0, tz3)
	et3 := time.Date(2034, time.December, 31, 23, 59, 59, 0, tz3)
	g3 := EndOfDay(t3)
	if g3 != et3 {
		t.Errorf("EndOfDay(%v) == %v, wanted %v", t3, g3, et3)
	}
}
