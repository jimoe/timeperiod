package timeperiod

import "strconv"

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
