package timeperiod

import (
	"fmt"
)

func ExampleDuration() {
	var value Period = "300ms"
	d := Duration(value)
	fmt.Printf("duration is %d milliseconds", d.Milliseconds())
	// Output: duration is 300 milliseconds
}

func ExampleSeconds() {
	var value Period
	value = "1h"
	s := Seconds(value)
	fmt.Printf("period is %d seconds", s)
	// Output: period is 3600 seconds
}

func ExampleMilliseconds() {
	ms := Milliseconds("5s")
	fmt.Printf("period is %d milliseconds", ms)
	// Output: period is 5000 milliseconds
}

func ExampleMicroseconds() {
	us := Microseconds("3ms")
	fmt.Printf("period is %d microseconds", us)
	// Output: period is 3000 microseconds
}

func ExampleNanoseconds() {
	ns := Nanoseconds("4us")
	fmt.Printf("period is %d nanoseconds", ns)
	// Output: period is 4000 nanoseconds
}
