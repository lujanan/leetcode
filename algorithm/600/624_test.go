// 624.数组列表中的最大距离

package algorithm_600

import "testing"

func Test_maxDistance(t *testing.T) {
	type args struct {
		arrays [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}}, 4},
		{"t2", args{[][]int{{1}, {1}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.arrays); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
