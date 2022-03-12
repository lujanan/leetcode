package algorithm_500

import "sort"

// 518. 零钱兑换 II
// https://leetcode-cn.com/problems/coin-change-2/

func change(amount int, coins []int) int {
	var dp = make([]int, amount+1)
	dp[0] = 1
	for i := range coins {
		if amount < coins[i] {
			continue
		}
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}

func change1(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	sort.Ints(coins)
	for i := range coins {
		if amount < coins[i] {
			coins = coins[:i]
			break
		}
	}
	if len(coins) < 1 {
		return 0
	}

	var cs = len(coins)
	var dp = make([][]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = make([]int, cs+1)
	}
	dp[amount][cs] = 1

	for i := amount; i >= 0; i-- {
		for j := cs - 1; j >= 0; j-- {
			dp[i][j] += dp[i][j+1]
			if i == 0 || dp[i][j] <= 0 {
				continue
			}

			if i-coins[j] >= 0 {
				dp[i-coins[j]][j] += dp[i][j]
			}
		}
	}
	return dp[0][0]
}
