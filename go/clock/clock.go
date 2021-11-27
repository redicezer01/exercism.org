package clock

import "fmt"

/*
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
	BenchmarkAddMinutes-8           355440529                3.302 ns/op
	BenchmarkSubtractMinutes-8      337415790                3.343 ns/op
	BenchmarkCreateClocks-8         177943062                6.667 ns/op
*/

// Clock type contains hours and minute.
type Clock int

const minInDay = 24 * 60
const minInHour = 60

// calcClock calculates minutes
func calcClock(m int) Clock {
	if m < 0 {
		m = minInDay + m%minInDay
	} else {
		m %= minInDay
	}
	return Clock(m)
}

// New creates Clock instance and sets with given time (hours, minutes).
func New(h, m int) Clock {
	return calcClock(h*minInHour + m)
}

// Add method adds minutes to current time.
func (c Clock) Add(m int) Clock {
	return calcClock(int(c) + m)
}

// Subtract method subtract minutes from current time.
func (c Clock) Subtract(m int) Clock {
	return calcClock(int(c) - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", int(c)/60, int(c)%60)
}
