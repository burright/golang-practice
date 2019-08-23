package gigasecond

import "time"

var gigasecond = 1000000000

func AddGigasecond(t time.Time) time.Time {

	// Add a gigasecond to the time provided and return
	return t.Add(time.Duration(gigasecond) * time.Second)
}
