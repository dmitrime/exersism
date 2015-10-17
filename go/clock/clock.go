package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	hour   int
	minute int
}

func Time(h int, m int) Clock {
	if m < 0 {
		h += -1 + m/60
		m = 60 + m%60
	}
	if h < 0 {
		h = 24 + h%24
	}
	return Clock{hour: (h + m/60) % 24, minute: m % 60}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(m int) Clock {
	return Time(c.hour, c.minute+m)
}
