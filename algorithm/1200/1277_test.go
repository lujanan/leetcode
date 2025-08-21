// 1277.统计全为1的正方形子矩阵

package algorithm_1200

import "testing"

func Test_countSquares(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}}, 15},
		{"t2", args{[][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}}, 7},
		{"t3", args{[][]int{{1, 1}, {0, 0}, {1, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSquares(tt.args.matrix); got != tt.want {
				t.Errorf("countSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
