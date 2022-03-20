package algorithm_competition

import (
	"testing"
)

func Test_countHillValley(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{2, 4, 1, 1, 6, 5}}, 3},
		{"t2", args{[]int{6, 6, 5, 5, 4, 1}}, 0},
		{"t3", args{[]int{1, 6, 5, 3, 4, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHillValley(tt.args.nums); got != tt.want {
				t.Errorf("countHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countCollisions(t *testing.T) {
	type args struct {
		directions string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"RLRSLL"}, 5},
		{"t2", args{"LLRR"}, 0},
		{"t3", args{"RSRLLRSLLRSRSRLRRRRLLRRLSRR"}, 20},
		{"t4", args{"RRLLRRL"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCollisions(tt.args.directions); got != tt.want {
				t.Errorf("countCollisions() = %v, want %v", got, tt.want)
			}
		})
	}
}
