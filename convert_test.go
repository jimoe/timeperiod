package timeperiod

import (
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    Period
		expected time.Duration
		ok       bool
	}{
		{"nanoseconds", "7ns", 7 * time.Nanosecond, true},
		{"microseconds u", "77us", 77 * time.Microsecond, true},
		{"microseconds µ", "66µs", 66 * time.Microsecond, true},
		{"milliseconds", "15ms", 15 * time.Millisecond, true},
		{"seconds", "1s", 1 * time.Second, true},
		{"no unit", "8", 8 * time.Second, true},
		{"minutes", "2m", 2 * time.Minute, true},
		{"hours", "3h", 3 * time.Hour, true},
		{"days", "4d", 4 * 24 * time.Hour, true},
		{"weeks", "1w", 7 * 24 * time.Hour, true},
		{"string", "invalid", 0, false},
		{"number and string", "12invalid", 0, false},
		{"string and number", "invalid99", 0, false},
		{"empty string", "", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur, ok := convert(tt.input)
			if dur != tt.expected || ok != tt.ok {
				t.Errorf("convert(%q) = %v, %v; want %v, %v",
					tt.input, dur, ok, tt.expected, tt.ok)
			}
		})
	}
}
