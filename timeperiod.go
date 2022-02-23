package timeperiod

import (
	"time"
)

func Seconds(period string) int {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	return int(duration.Seconds())
}

func Milliseconds(period string) int {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano / (1000 * 1000))
}

func Microseconds(period string) int {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano / 1000)
}

func Nanoseconds(period string) int {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano)
}

func Duration(period string) time.Duration {
	duration, ok := convert(period)
	if !ok {
		return 0
	}
	return duration
}
