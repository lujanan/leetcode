package algorithm_800

// 813. 最大平均值和的分组
// https://leetcode-cn.com/problems/largest-sum-of-averages/

func largestSumOfAverages(nums []int, k int) float64 {
	var max = func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}
	var sum = make([]float64, len(nums)+1)

	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + float64(nums[i-1])
	}

	var dp = make([][]float64, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]float64, k+1)
	}

	for i := 1; i <= len(nums); i++ {
		dp[i][1] = sum[i] / float64(i)
		for tk := 2; tk <= k && tk <= i; tk++ {
			for j := 1; j < i; j++ {
				dp[i][tk] = max(dp[i][tk], dp[j][tk-1]+(sum[i]-sum[j])/float64(i-j))
			}
		}
	}
	return dp[len(nums)][k]
}
