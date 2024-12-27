// 80. 删除有序数组中的重复项 II

package algorithm_0

import "testing"

func Test_removeDuplicatesV2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 1, 1, 2, 2, 3}}, 5},
		{"t2", args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesV2(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicatesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
