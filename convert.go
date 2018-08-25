package timeperiod

import (
	"strconv"
	"time"
)

func convert(period string) (dur time.Duration, ok bool) {
	valueAsStr, unit := splitValueAndUnit(period)

	value, err := strconv.Atoi(valueAsStr)
	if err != nil {
		return time.Duration(0), false
	}

	ok = true
	switch unit {
	case "":
		dur = time.Duration(value) * time.Second
	case "w":
		dur = time.Duration(value*24*7) * time.Hour
	case "d":
		dur = time.Duration(value*24) * time.Hour
	default:
		var err error
		dur, err = time.ParseDuration(period)
		if err != nil {
			ok = false
		}
	}

	return
}
