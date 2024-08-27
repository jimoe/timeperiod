package timeperiod

import (
	"unicode/utf8"
)

func splitValueAndUnit(period string) (valueAsStr string, unit string) {
	strLength := len(period)
	unitStrLength := 0

	lastChar, size := utf8.DecodeLastRuneInString(period)
	unitStrLength += size
	if isNumber(string(lastChar)) {
		valueAsStr = period
		return
	}

	nextToLastChar, size := utf8.DecodeLastRuneInString(period[:strLength-size])
	if !isNumber(string(nextToLastChar)) {
		unitStrLength += size
	}

	position := strLength - unitStrLength
	unit = period[position:]
	valueAsStr = period[0:position]

	return
}
