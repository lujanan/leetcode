// 300. 最长递增子序列

package algorithm_300

// dp方程：
// 1.如果 j < i && num[j] < num[i]时，那dp[i] = max(dp[j]) + 1
// 2.最长的子序列不一定包括数组的最后一个元素
func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums)+1)
	dp[0], dp[len(nums)] = 1, 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		dp[len(nums)] = max(dp[len(nums)], dp[i])
	}
	return dp[len(nums)]
}

func lengthOfLISV2(nums []int) int {
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
