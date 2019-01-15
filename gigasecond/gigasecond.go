// Package gigasecond adds a gigasecond (1,000,000,000) to the Date / Time value provided.
package gigasecond

import "time"

// AddGigasecond takes a date or date time value, adds 1 gigasecond
// and returns the result as a time.Time value
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1000000000) * time.Second)
}