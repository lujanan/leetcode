package algorithm_700

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{[]int{1, 0, 0}}, true},
		{"t2", args{[]int{1, 1, 1, 0}}, false},
		{"t3", args{[]int{1, 1, 1, 1, 0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneBitCharacter(tt.args.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
