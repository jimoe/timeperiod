package timeperiod

import (
	"testing"
)

func Test_isNumber(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"number", args{"7"}, true},
		{"string", args{"syv"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.str); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
