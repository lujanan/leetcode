//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
//
//
//
// 示例 1：
//
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
// 示例 2：
//
//
//输入：coins = [2], amount = 3
//输出：-1
//
// 示例 3：
//
//
//输入：coins = [1], amount = 0
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= coins.length <= 12
// 1 <= coins[i] <= 2³¹ - 1
// 0 <= amount <= 10⁴
//
// Related Topics 广度优先搜索 数组 动态规划 👍 1732 👎 0

package algorithm_300

func coinChange(coins []int, amount int) int {
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dp = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1 // 默认无法凑成

		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] >= 0 {
				if dp[i] < 0 {
					dp[i] = dp[i-coins[j]] + 1
				} else {
					// dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[j]]+1)))
					dp[i] = min(dp[i], dp[i-coins[j]]+1)
				}
			}
		}
	}

	return dp[amount]
}

func coinChangeV2(coins []int, amount int) int {
	// dp[i] = min(dp[i-k]...) +1

	var dp = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
