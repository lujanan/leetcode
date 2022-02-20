package algorithm_300

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	if amount < coins[0] {
		return 0
	}

	var dp = make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		if coins[i] >= amount {
			break
		}
		dp[coins[i]] = 1
	}

	for i := coins[0]; i <= amount; i++ {
		for i := 0; i < len(coins); i++ {
			if coins[i] >= amount {
				break
			}
			dp[coins[i]] = 1
		}
		dp[i] = min(dp[i-1], dp[i-2]) + 1
	}

	var fn func(money int) int
	fn = func(money int) int {
		if money < 0 {
			return -1
		} else if money == 0 {
			return 0
		}
		var mi = math.MaxInt64
		for _, v := range coins {
			tmp := fn(money - v)
			if tmp >= 0 && mi > tmp {
				mi = tmp
			}
		}
		return mi + 1
	}
	return fn(amount)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
