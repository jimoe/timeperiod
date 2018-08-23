package timeperiod

import (
	"time"
)

func convert(period string) (dur time.Duration, ok bool) {
	value, unit, ok := splitValueAndUnit(period)
	if !ok {
		return
	}

	switch unit {
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
