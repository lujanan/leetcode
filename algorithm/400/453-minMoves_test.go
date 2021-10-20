//给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：3
//解释：
//只需要3次操作（注意每次操作会增加两个元素的值）：
//[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1]
//输出：0
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// 答案保证符合 32-bit 整数
//
// Related Topics 数组 数学 👍 309 👎 0

package algorithm_400

import "testing"

func Test_minMoves(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{[]int{1, 2, 3}}, 3},
		{"t2", args{[]int{1, 1, 1}}, 0},
		{"t3", args{[]int{1, 2, 3, 4, 5}}, 10},
		{"t4", args{[]int{1, 1, 1000000000}}, 999999999},
		{"t5", args{[]int{5, 6, 8, 8, 5}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := minMoves(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("minMoves() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
