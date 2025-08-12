// 2787.将一个数字表示成幂的和的方案数

package algorithm_2700

import "math"

// 动态规划
// 0-1背包问题，针对1~n的数字取或不取进行遍历
func numberOfWays(n int, x int) int {
	var dp = make([]int, n+1)
	dp[0] = 1
	var numPowX int
	for i := 1; i <= n; i++ {
		numPowX = int(math.Pow(float64(i), float64(x)))
		if numPowX > n {
			break
		}
		for j := n; j >= 0; j-- {
			if j >= numPowX {
				dp[j] = (dp[j] + dp[j-numPowX]) % 1000000007
			}
		}
	}
	return dp[n]
}

// 回溯+剪枝
func numberOfWays1(n int, x int) int {
	var ansMap = make(map[int]int)
	var searchFn func(num, idx int) int
	searchFn = func(num, idx int) int {
		var tmp = 1
		for i := 0; i < x; i++ {
			tmp *= idx
			if tmp > num {
				return 0
			}
		}
		if tmp == num {
			return 1
		}

		if a, ok := ansMap[num|(idx<<9)]; ok {
			return a
		}

		ansMap[num|(idx<<9)] = (searchFn(num, idx+1) + searchFn(num-tmp, idx+1)) % 1000000007
		return ansMap[num|(idx<<9)]
	}

	return searchFn(n, 1)
}
