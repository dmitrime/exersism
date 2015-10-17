package gigasecond

import "time"

const TestVersion int = 2

var Birthday time.Time = parse("1987-04-04T23:55:00", nil)

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(1000*1000*1000))
}
