//给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1：
//
//
//输入：k = 2, prices = [2,4,1]
//输出：2
//解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
//
// 示例 2：
//
//
//输入：k = 2, prices = [3,2,6,5,0,3]
//输出：7
//解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3
//。
//
//
//
// 提示：
//
//
// 0 <= k <= 100
// 0 <= prices.length <= 1000
// 0 <= prices[i] <= 1000
//
// Related Topics 数组 动态规划 👍 669 👎 0

package algorithm_100

import "testing"

func Test_maxProfit_188(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t0", args{2, []int{}}, 0},
		{"t1", args{2, []int{2, 4, 1}}, 2},
		{"t2", args{2, []int{3, 2, 6, 5, 0, 3}}, 7},
		{"t3", args{5, []int{31, 3, 5, 20, 0, 3, 1, 4, 1, 2, 34, 5, 12, 3, 4, 45, 46, 12, 34, 23, 45, 67, 89, 123, 324, 450, 456}}, 549},
		{"t4", args{15, []int{31, 3, 5, 20, 0, 3, 1, 4, 1, 2, 34, 5, 12, 3, 4, 45, 46, 12, 34, 23, 45, 67, 89, 123, 324, 450, 456}}, 561},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_188(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit_188() = %v, want %v", got, tt.want)
			}
		})
	}
}
