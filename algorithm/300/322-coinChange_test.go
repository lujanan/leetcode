package algorithm_300

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2, 5}, 11}, 3},
		{"t2", args{[]int{2}, 3}, -1},
		{"t3", args{[]int{1}, 0}, 0},
		{"t4", args{[]int{1, 2, 5}, 145}, 29},
		{"t5", args{[]int{1, 10, 20, 55, 57, 63}, 5999}, 96},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
