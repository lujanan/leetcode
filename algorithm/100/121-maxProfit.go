package algorithm_100

func maxProfit(prices []int) int {
	var money, buy int
	for i := 1; i < len(prices); i++ {
		if money < prices[i]-prices[buy] {
			money = prices[i] - prices[buy]
		} else if prices[i] < prices[buy] {
			buy = i
		}
	}
	return money
}

// 动态规划
// 重点在保证至多 只买只卖 1 次
func maxProfitDp(prices []int) int {
	var dp0, dp1 = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, -prices[i]) // 保证只买一次
	}
	return dp0
}
