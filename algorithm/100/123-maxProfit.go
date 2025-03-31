//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
//
//输入：prices = [3,3,5,0,0,3,1,4]
//输出：6
//解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//
// 示例 2：
//
//
//输入：prices = [1,2,3,4,5]
//输出：4
//解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
//
// 示例 3：
//
//
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
//
// 示例 4：
//
//
//输入：prices = [1]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 10⁵
// 0 <= prices[i] <= 10⁵
//
// Related Topics 数组 动态规划 👍 1021 👎 0

package algorithm_100

func maxProfit_123(prices []int) int {
	// 状态机
	var buy1, sell1, buy2, sell2 = -prices[0], 0, -prices[0], 0 // 第1天不管买卖多少次，手上无股票时收益=0，有股票是收益=-prices[0]
	// 状态转移 -> buy1 -> sell1 -> buy2 -> sell2 -> 交易两次收益
	for i := 1; i < len(prices); i++ {
		sell2 = max(sell2, buy2+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy1 = max(buy1, -prices[i])
	}
	return max(0, buy1, sell1, buy2, sell2)
}

func maxProfit_1235(prices []int) int {
	// dp[i][k][0] = dp[i-1][k][0] || dp[i-1][k-1][1] + p[i]
	// dp[i][k][1] = dp[i-1][k][0] - p[i] || dp[i-1][k][1]

	var k = 2
	var dp = make([][][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][2]int, k+1)
	}

	for i := 0; i < len(prices); i++ {
		for j := 0; j <= k; j++ {
			if i < 1 {
				dp[i][j][1] = -prices[i] // 第1天不管买卖多少次，手上无股票时收益=0，有股票是收益=-prices[0]

			} else {
				// 同一天同时进行买和卖操作收益=0，只有买1次或卖1次才会产生实际收益
				// j表示交易次数，卖出时才+1
				if j < 1 {
					// j = 0 时，表示当天没有卖出行为，但可以今日买入 或 取前1日的股票
					dp[i][j][1] = max(-prices[i], dp[i-1][j][1])

				} else {
					// 有交易次数(卖出)时
					// 手上无股票: 取前1日j次交易的值 或 把前1日j-1次操作持有的股票卖出
					// 手上有股票: 取前1日j次交易的值后再次买入后的值 或 把前1日j次卖出股票后的值+进入买股票
					dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i])
					dp[i][j][1] = max(dp[i-1][j][0]-prices[i], dp[i-1][j][1])
				}
			}
		}
	}

	var res = dp[len(prices)-1][0][0]
	for i := 1; i <= k; i++ {
		if res < dp[len(prices)-1][k][0] {
			res = dp[len(prices)-1][k][0]
		}
	}
	return res
}

func maxProfit_1234(prices []int) int {
	// 答案可同时解答 121, 123, 188 题

	// 买卖不多于 k 次，在第 i 天的状态，dp[i][0] 表示当天没股票，dp[i][1] 表示当天有股票
	// 买卖 0 次， dp[0][0] 恒为 0
	// 买卖 1 次,  dp[1][0] = max(dp[1][0], dp[1][1] + prices[i])
	//            dp[1][1] = max(dp[1][1], dp[1-1][0]-prices[i]) 当前持有股票 或 买卖 0 次数时卖出股票的情况, 保证只买卖 1 次
	// 买卖 2 次,  dp[2][0] = max(dp[2][0], dp[2][1] + prices[i])
	//            dp[2][1] = max(dp[2][1], dp[2-1][0]-prices[i]) 当前持有股票 或 买卖 1 次数时卖出股票的情况, 保证最多只买卖 2 次
	// ....
	// 买卖 k 次,  dp[k][0] = max(dp[k][0], dp[k][1] + prices[i])
	//            dp[k][1] = max(dp[k][1], dp[k-1][0]-prices[i]) 当前持有股票 或 买卖 k-1 次数时卖出股票的情况, 保证最多只买卖 k 次

	var k = 2
	var dp = make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][1] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[j][0], dp[j][1] = max(dp[j][0], dp[j][1]+prices[i]), max(dp[j][1], dp[j-1][0]-prices[i])
		}
	}

	var res = dp[0][0]
	for i := 1; i < len(dp); i++ {
		if res < dp[i][0] {
			res = dp[i][0]
		}
	}
	return res
}
