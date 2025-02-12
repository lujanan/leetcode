package algorithm_competition

func maximumTop(nums []int, k int) int {
	var nMap = make(map[int]int)
	for i := range nums {
		nMap[nums[i]]++
	}

	// dp[i][0] = max(num[i+1], dp[i-1][1])
	// dp[i][0] = max(num[i+1], dp[i-1][1])
	// dp[i][1] = max(num[i], dp[i-1][1])

	var dp = make([][2]int, k+1)
	dp[0][0], dp[0][1] = nums[0], -1

	for i := 1; i <= k; i++ {
		dp[i][1] = nums[i]

	}

	return dp[k][0]
}
