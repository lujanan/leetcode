package algorithm_2100

// 2100. 适合打劫银行的日子
// https://leetcode-cn.com/problems/find-good-days-to-rob-the-bank/

func goodDaysToRobBank(security []int, time int) []int {
	var dp = make([]int, len(security))
	for i := len(security) - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			dp[i] += dp[i+1] + 1
		}
	}

	var left int
	var res []int
	for i := range security {
		if i > 0 && security[i] <= security[i-1] {
			left += 1
		} else {
			left = 0
		}

		if left >= time && dp[i] >= time {
			res = append(res, i)
		}
	}
	return res
}
