package algorithm_300

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"leetcode"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
