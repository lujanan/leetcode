//给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；整数 fee 代表了交易股票的手续费用。
//
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
//
// 返回获得利润的最大值。
//
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
//
//
//
// 示例 1：
//
//
//输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
//输出：8
//解释：能够达到的最大利润:
//在此处买入 prices[0] = 1
//在此处卖出 prices[3] = 8
//在此处买入 prices[4] = 4
//在此处卖出 prices[5] = 9
//总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
//
// 示例 2：
//
//
//输入：prices = [1,3,7,5,10,3], fee = 3
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 5 * 10⁴
// 1 <= prices[i] < 5 * 10⁴
// 0 <= fee < 5 * 10⁴
//
// Related Topics 贪心 数组 动态规划 👍 612 👎 0

package algorithm_700

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 3, 2, 8, 4, 9}, 2}, 8},
		{"t2", args{[]int{1, 3, 7, 5, 10, 3}, 3}, 6},
		{"t3", args{[]int{12, 42, 56, 34, 67, 12, 78, 45, 90, 1}, 5}, 168},
		{"t4", args{[]int{12, 42, 56, 34, 67, 12, 78, 45, 90, 1, 23, 12, 56, 78}, 5}, 246},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
