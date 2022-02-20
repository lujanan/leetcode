package algorithm_100

import "testing"

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, 11},
		{"t2", args{[][]int{{-10}}}, -10},
		{"t3", args{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}, {12, 34, 5, 7, 34}, {34, 56, 23, 45, 67, 344}}}, 39},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
