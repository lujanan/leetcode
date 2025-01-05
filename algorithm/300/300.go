// 300. 最长递增子序列

package algorithm_300

func lengthOfLIS(nums []int) int {
	var dp = make([][2]int, len(nums)+1)
	dp[0], dp[1] = [2]int{-10e4 - 1, 1}, [2]int{nums[0], 1}
	for i := 1; i < len(nums); i++ {
		for j := i; j >= 0; j-- {
			if dp[j][1] < 1 || dp[j][0] >= nums[i] {
				continue
			}
			if dp[j+1][1] < 1 || (dp[j+1][1] > 0 && dp[j+1][0] > nums[i]) {
				dp[j+1] = [2]int{nums[i], 1}
			}
		}
	}

	for i := len(dp) - 1; i > 0; i-- {
		if dp[i][1] > 0 {
			return i
		}
	}
	return 1
}
