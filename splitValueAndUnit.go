package timeperiod

func splitValueAndUnit(period string) (valueAsStr string, unit string) {
	strLength := len(period)

	lastChar := period[strLength-1:]
	if isNumber(lastChar) {
		valueAsStr = period
		return
	}

	unitStrLength := 1
	nextToLastChar := period[strLength-2 : strLength-1]
	if !isNumber(nextToLastChar) {
		unitStrLength = 2
	}

	position := strLength - unitStrLength
	unit = period[position:]
	valueAsStr = period[0:position]

	return
}
