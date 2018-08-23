package timeperiod

import (
	"strconv"
	"time"
)

func convert(period string) (dur time.Duration, ok bool) {
	lastCharPosition := len(period) - 1
	lastChar := period[lastCharPosition:]
	numberAsText := period[0:lastCharPosition]

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
