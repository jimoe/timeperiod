package timeperiod

import "strconv"

func splitValueAndUnit(period string) (value int, unit string, ok bool) {
	ok = true
	strLength := len(period)

	unitStrLength := 1
	nextToLastChar := period[strLength-2 : strLength-1]
	if !isNumber(nextToLastChar) {
		unitStrLength = 2
	}

	position := strLength - unitStrLength
	unit = period[position:]
	valueAsText := period[0:position]

	value, err := strconv.Atoi(valueAsText)
	if err != nil {
		ok = false
	}

	return
}
