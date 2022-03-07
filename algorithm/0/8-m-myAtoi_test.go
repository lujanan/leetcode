package algorithm_0

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"42"}, 42},
		{"t2", args{"   -42"}, -42},
		{"t3", args{"4193 with words"}, 4193},
		{"t4", args{"  +042u12iewyrn"}, 42},
		{"t5", args{"  -042u12iewyrn"}, -42},
		{"t6", args{"  0+042u12iewyrn"}, 0},
		{"t7", args{"9223372036854775808"}, 2147483647},
		{"t8", args{"-5-"}, -5},
		{"t9", args{" - 5"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
