// 3363.最多可收集的水果数目

package algorithm_3300

import "testing"

func Test_maxCollectedFruits(t *testing.T) {
	type args struct {
		fruits [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[][]int{{1, 2, 3, 4}, {5, 6, 8, 7}, {9, 10, 11, 12}, {13, 14, 15, 16}}}, 100},
		{"t2", args{[][]int{{1, 1}, {1, 1}}}, 4},
		{"t3", args{[][]int{
			{4,  18, 19, 9,  20, 14}, 
			{16, 4,  4,  16, 15, 16}, 
			{2,  11, 15, 6,  8,  9}, 
			{6,  7,  11, 17, 7,  6}, 
			{17, 17, 2,  13, 2,  14}, 
			{16, 9,  6,  14, 7,  16}}}, 182},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCollectedFruits(tt.args.fruits); got != tt.want {
				t.Errorf("maxCollectedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
