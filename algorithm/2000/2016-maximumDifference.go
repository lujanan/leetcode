package algorithm_2000

// 2016. 增量元素之间的最大差值
// simple
// https://leetcode-cn.com/problems/maximum-difference-between-increasing-elements/

func maximumDifference(nums []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dp1, dp0 = -nums[0], 0
	for i := 1; i < len(nums); i++ {
		dp1, dp0 = max(dp1, -nums[i]), max(dp0, dp1+nums[i])
	}
	if dp0 <= 0 {
		dp0 = -1
	}
	return dp0
}
