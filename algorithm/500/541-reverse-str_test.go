package algorithm_500

import (
	"testing"
)

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"abcdefg", 2}, "bacdfeg"},
		{"t2", args{"abcd", 2}, "bacd"},
		{"t3", args{"abcdefg", 3}, "cbadefg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
