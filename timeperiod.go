package timeperiod

import (
	"strconv"
	"time"
)

func Seconds(period string) int {
	duration, ok := work(period)
	if !ok {
		return 0
	}
	return int(duration.Seconds())
}

func Milliseconds(period string) int {
	duration, ok := work(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano / (1000 * 1000))
}

func Microseconds(period string) int {
	duration, ok := work(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano / 1000)
}

func Nanoseconds(period string) int {
	duration, ok := work(period)
	if !ok {
		return 0
	}
	nano := duration.Nanoseconds()
	return int(nano)
}

func work(period string) (dur time.Duration, ok bool) {
	lastChar := period[len(period)-1:]
	numberAsText := period[0 : len(period)-1]
	number, err := strconv.Atoi(numberAsText)
	if err != nil {
		return 0, false
	}

	ok = true
	switch lastChar {
	case "w":
		dur = time.Duration(number*24*7) * time.Hour
	case "d":
		dur = time.Duration(number*24) * time.Hour
	default:
		var err error
		dur, err = time.ParseDuration(period)
		if err != nil {
			ok = false
		}
	}

	return
}
