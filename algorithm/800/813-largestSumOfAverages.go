package algorithm_800

// 813. 最大平均值和的分组
// https://leetcode-cn.com/problems/largest-sum-of-averages/

func largestSumOfAverages(nums []int, k int) float64 {
	var nl = len(nums)
	var sum = make([]float64, nl+1)
	for i := range nums {
		sum[i+1] = sum[i] + float64(nums[i])
	}

	var dp = make([]float64, nl+1)
	for i := range sum {
		dp[i] = (sum[nl] - sum[i]) / float64(nl-i)
	}
	dp[nl] = 0

	for tk := 2; tk <= k; tk++ {
		for i := 0; i <= nl; i++ {
			for j := i + 1; j <= nl; j++ {
				dp[i] = max(dp[i], (sum[j]-sum[i])/float64(j-i)+dp[j])
			}
		}
	}
	return dp[0]
}

func largestSumOfAverages2(nums []int, k int) float64 {
	var nl = len(nums)
	var sum = make([]float64, nl+1)
	var dp = make([][]float64, nl+1)
	dp[nl] = make([]float64, k+1)
	for i := range nums {
		sum[i+1] = sum[i] + float64(nums[i])
		dp[i] = make([]float64, k+1)
	}

	for i := 1; i <= nl; i++ {
		dp[i][1] = sum[i] / float64(i)
		for l := 2; l <= k && l <= i; l++ {
			for j := 1; j < i; j++ {
				dp[i][l] = max(dp[i][l], dp[j][l-1]+(sum[i]-sum[j])/float64(i-j))
			}
		}
	}
	return dp[nl][k]
}

func largestSumOfAverages1(nums []int, k int) float64 {
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
