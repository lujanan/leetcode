package algorithm_700

import "sort"

// 740. 删除并获得点数
// https://leetcode-cn.com/problems/delete-and-earn/

func deleteAndEarn(nums []int) int {
	var nMap = make(map[int]int)
	var tmp []int
	for _, v := range nums {
		if _, ok := nMap[v]; ok {
			nMap[v] += v
		} else {
			nMap[v] = v
			tmp = append(tmp, v)
		}
	}
	sort.Ints(tmp)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dp = make([][2]int, len(tmp))
	dp[0] = [2]int{0, nMap[tmp[0]]}

	for i := 1; i < len(tmp); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		if tmp[i]-tmp[i-1] == 1 {
			dp[i][1] = max(dp[i-1][0]+nMap[tmp[i]], dp[i-1][1])
		} else {
			dp[i][1] = dp[i][0] + nMap[tmp[i]]
		}
	}
	return dp[len(tmp)-1][1]
}
