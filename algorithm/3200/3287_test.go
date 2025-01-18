// 3287. 求出数组中最大序列值

package algorithm_3200

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{2, 6, 7}, 1}, 5},
		{"t2", args{[]int{4, 2, 5, 6, 7}, 2}, 2},
		{"t3", args{[]int{2, 6, 17, 19, 23, 34, 45, 56, 67, 79, 111, 123}, 2}, 125},
		{"t4", args{[]int{2, 6, 17, 19, 23, 34}, 2}, 53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
