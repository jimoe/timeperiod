package timeperiod

import (
	"time"
)

// Seconds converts a time string with unit to an integer number of seconds
// Valid time units are 'ns', 'us', 'ms', 's', 'm', 'h', `d` and `w`
func Seconds(period string) int {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	return int(duration.Seconds())
}

// Milliseconds converts a time string with unit to an integer number of milliseconds
// Valid time units are 'ns', 'us', 'ms', 's', 'm', 'h', `d` and `w`
func Milliseconds(period string) int {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano / (1000 * 1000))
}

// Microseconds converts a time string with unit to an integer number of microseconds
// Valid time units are 'ns', 'us', 'ms', 's', 'm', 'h', `d` and `w`
func Microseconds(period string) int {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano / 1000)
}

// Nanoseconds converts a time string with unit to an integer number of nanoseconds
// Valid time units are 'ns', 'us', 'ms', 's', 'm', 'h', `d` and `w`
func Nanoseconds(period string) int {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano)
}

// Duration converts a time string with unit to time.Duration
// Valid time units are 'ns', 'us', 'ms', 's', 'm', 'h', `d` and `w`
func Duration(period string) time.Duration {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	return duration
}
