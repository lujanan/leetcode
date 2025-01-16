// 3097. 或值至少 K 的最短子数组 II

package algorithm_3000

import "testing"

func Test_minimumSubarrayLength2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2, 3}, 2}, 1},
		{"t2", args{[]int{2, 1, 8}, 10}, 3},
		{"t3", args{[]int{1, 2}, 0}, 1},
		{"t4", args{[]int{1, 2, 12, 1, 16, 10}, 20}, 2},
		{"t5", args{[]int{67,24,1,32,2}, 110}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSubarrayLength2(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumSubarrayLength2() = %v, want %v", got, tt.want)
			}
		})
	}
}
