package timeperiod

import (
	"unicode/utf8"
)

func splitValueAndUnit(period Period) (valueAsStr string, unit string) {
	strLength := len(period)
	unitStrLength := 0

	lastChar, size := utf8.DecodeLastRuneInString(period.String())
	unitStrLength += size
	if isNumber(string(lastChar)) {
		valueAsStr = period.String()
		return
	}

	nextToLastChar, size := utf8.DecodeLastRuneInString(period.String()[:strLength-size])
	if !isNumber(string(nextToLastChar)) {
		unitStrLength += size
	}

	position := strLength - unitStrLength
	unit = period.String()[position:]
	valueAsStr = period.String()[0:position]

	return
}
