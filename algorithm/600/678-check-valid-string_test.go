package algorithm_600

import "testing"

func Test_checkValidString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{"()"}, true},
		{"t2", args{"(*)"}, true},
		{"t3", args{"(*))"}, true},
		{"t4", args{"(*)))"}, false},
		{"t5", args{"(*)()((())"}, false},
		{"t6", args{"(*)()(((*)**))"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.args.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
