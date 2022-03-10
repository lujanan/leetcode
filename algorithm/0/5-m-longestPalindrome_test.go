package algorithm_0

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"babad"}, "bab"},
		{"t2", args{"cbbd"}, "bb"},
		{"t3", args{"cbbdddd"}, "dddd"},
		{"t4", args{"cbbdddddddddddd"}, "dddddddddddd"},
		{"t5", args{"cbbddbbdd"}, "bbddbb"},
		{"t6", args{"ababababa"}, "ababababa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
