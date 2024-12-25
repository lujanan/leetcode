//给你一个整数数组 nums ，一个整数 k 和一个整数 multiplier 。 
//
// 你需要对 nums 执行 k 次操作，每次操作中： 
//
// 
// 找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。 
// 将 x 替换为 x * multiplier 。 
// 
//
// k 次操作以后，你需要将 nums 中每一个数值对 10⁹ + 7 取余。 
//
// 请你返回执行完 k 次乘运算以及取余运算之后，最终的 nums 数组。 
//
// 
//
// 示例 1： 
//
// 
// 输入：nums = [2,1,3,5,6], k = 5, multiplier = 2 
// 
//
// 输出：[8,4,6,5,6] 
//
// 解释： 
//
// 
// 
// 
// 操作 
// 结果 
// 
// 
// 1 次操作后 
// [2, 2, 3, 5, 6] 
// 
// 
// 2 次操作后 
// [4, 2, 3, 5, 6] 
// 
// 
// 3 次操作后 
// [4, 4, 3, 5, 6] 
// 
// 
// 4 次操作后 
// [4, 4, 6, 5, 6] 
// 
// 
// 5 次操作后 
// [8, 4, 6, 5, 6] 
// 
// 
// 取余操作后 
// [8, 4, 6, 5, 6] 
// 
// 
// 
//
// 示例 2： 
//
// 
// 输入：nums = [100000,2000], k = 2, multiplier = 1000000 
// 
//
// 输出：[999999307,999999993] 
//
// 解释： 
//
// 
// 
// 
// 操作 
// 结果 
// 
// 
// 1 次操作后 
// [100000, 2000000000] 
// 
// 
// 2 次操作后 
// [100000000000, 2000000000] 
// 
// 
// 取余操作后 
// [999999307, 999999993] 
// 
// 
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10⁴ 
// 1 <= nums[i] <= 10⁹ 
// 1 <= k <= 10⁹ 
// 1 <= multiplier <= 10⁶ 
// 
//
// Related Topics 数组 模拟 堆（优先队列） 👍 40 👎 0

package algorithm_3200

import (
	"reflect"
	"testing"
)

func Test_getFinalStateV2(t *testing.T) {
	type args struct {
		nums       []int
		k          int
		multiplier int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{2, 1, 3, 5, 6}, 5, 2}, []int{8, 4, 6, 5, 6}},
		{"t2", args{[]int{100000, 2000}, 2, 1000000}, []int{999999307, 999999993}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFinalStateV2(tt.args.nums, tt.args.k, tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFinalStateV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
