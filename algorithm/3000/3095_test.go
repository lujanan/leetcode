// 3095. 或值至少 K 的最短子数组 I

package algorithm_3000

import "testing"

func Test_minimumSubarrayLength(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSubarrayLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimumSubarrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
