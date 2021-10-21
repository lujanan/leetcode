//给定一个数组 A，将其划分为两个连续子数组 left 和 right， 使得：
//
//
// left 中的每个元素都小于或等于 right 中的每个元素。
// left 和 right 都是非空的。
// left 的长度要尽可能小。
//
//
// 在完成这样的分组后返回 left 的长度。可以保证存在这样的划分方法。
//
//
//
// 示例 1：
//
//
//输入：[5,0,3,8,6]
//输出：3
//解释：left = [5,0,3]，right = [8,6]
//
//
// 示例 2：
//
//
//输入：[1,1,1,0,6,12]
//输出：4
//解释：left = [1,1,1,0]，right = [6,12]
//
//
//
//
// 提示：
//
//
// 2 <= A.length <= 30000
// 0 <= A[i] <= 10^6
// 可以保证至少有一种方法能够按题目所描述的那样对 A 进行划分。
//
//
//
// Related Topics 数组 👍 86 👎 0

package algorithm_900

import "testing"

func Test_partitionDisjoint(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{[]int{5, 0, 3, 8, 6, 5}}, 3},
		{"t2", args{[]int{1, 1, 1, 0, 6, 12}}, 4},
		{"t3", args{[]int{5, 0, 3, 8, 6, 10, 11, 5, 4, 11, 12, 17, 14}}, 9},
		{"t4", args{[]int{0, 5, 3, 8, 6, 10, 11, 5, 4, 11, 12, 17, 14}}, 1},
		{"t5", args{[]int{90, 47, 69, 10, 43, 92, 31, 73, 61, 97}}, 9},
		{"t6", args{[]int{6, 0, 8, 30, 37, 6, 75, 98, 39, 90, 63, 74, 52, 92, 64}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := partitionDisjoint(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("partitionDisjoint() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
