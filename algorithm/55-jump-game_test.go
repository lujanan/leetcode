package algorithm

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{nums: []int{2, 3, 1, 1, 4}}, true},
		{"t2", args{nums: []int{3, 2, 1, 0, 4}}, false},
		{"t3", args{nums: []int{1, 1, 0, 4, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
