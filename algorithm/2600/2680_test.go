// 2680.最大或值
//给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 k 。每一次操作中，你可以选择一个数并将它乘 2 。
//
// 你最多可以进行 k 次操作，请你返回 nums[0] | nums[1] | ... | nums[n - 1] 的最大值。
//
// a | b 表示两个整数 a 和 b 的 按位或 运算。
//
//
//
// 示例 1：
//
//
//输入：nums = [12,9], k = 1
//输出：30
//解释：如果我们对下标为 1 的元素进行操作，新的数组为 [12,18] 。此时得到最优答案为 12 和 18 的按位或运算的结果，也就是 30 。
//
//
// 示例 2：
//
//
//输入：nums = [8,1,2], k = 2
//输出：35
//解释：如果我们对下标 0 处的元素进行操作，得到新数组 [32,1,2] 。此时得到最优答案为 32|1|2 = 35 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
// 1 <= k <= 15
//
//
// Related Topics 贪心 位运算 数组 前缀和 👍 65 👎 0

package algorithm_2600

import "testing"

func Test_maximumOr(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"t1", args{[]int{12, 9}, 1}, 30},
		{"t2", args{[]int{8, 1, 2}, 2}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumOr(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
