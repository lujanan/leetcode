// 2275. 按位与结果大于零的最长组合

package algorithm_2200

import "testing"

func Test_largestCombination(t *testing.T) {
	type args struct {
		candidates []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{16, 17, 71, 62, 12, 24, 14}}, 4},
		{"t2", args{[]int{8, 8}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestCombination(tt.args.candidates); got != tt.want {
				t.Errorf("largestCombination() = %v, want %v", got, tt.want)
			}
		})
	}
}
