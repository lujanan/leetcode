package algorithm_700

import "sort"

// 740. 删除并获得点数
// https://leetcode-cn.com/problems/delete-and-earn/

// 转化成打家劫舍问题
func deleteAndEarn(nums []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 耗费空间
	var mVal int
	for _, v := range nums {
		mVal = max(mVal, v)
	}
	var tmp = make([]int, mVal+1)
	for _, v := range nums {
		tmp[v] += v
	}

	var dp = [2]int{0, tmp[0]}
	for i := 1; i < len(tmp); i++ {
		dp[0], dp[1] = dp[1], max(dp[1], dp[0]+tmp[i])
	}
	return max(dp[0], dp[1])
}

func deleteAndEarn1(nums []int) int {
	var nMap = make(map[int]int)
	var tmp []int
	for _, v := range nums {
		if _, ok := nMap[v]; !ok {
			tmp = append(tmp, v)
		}
		nMap[v] += v
	}
	sort.Ints(tmp)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dp = [2][2]int{}
	dp[0] = [2]int{0, nMap[tmp[0]]}

	for i := 1; i < len(tmp); i++ {
		dp[1][0] = max(dp[0][0], dp[0][1])
		if tmp[i]-tmp[i-1] == 1 {
			dp[1][1] = max(dp[0][0]+nMap[tmp[i]], dp[0][1])
		} else {
			dp[1][1] = dp[1][0] + nMap[tmp[i]]
		}
		dp[0] = dp[1]
	}
	return dp[0][1]
}
