// 3287. 求出数组中最大序列值

package algorithm_3200

func maxValue(nums []int, k int) int {
	findORs := func(nums []int, k int) []map[int]bool {
		dp := make([]map[int]bool, 0)
		prev := make([]map[int]bool, k+1)
		for i := 0; i <= k; i++ {
			prev[i] = make(map[int]bool)
		}
		prev[0][0] = true

		for i := 0; i < len(nums); i++ {
			for j := min(k-1, i+1); j >= 0; j-- {
				for x := range prev[j] {
					prev[j+1][x|nums[i]] = true
				}
			}
			current := make(map[int]bool)
			for key := range prev[k] {
				current[key] = true
			}
			dp = append(dp, current)
		}
		return dp
	}

	A := findORs(nums, k)
	reverse(nums)
	B := findORs(nums, k)
	mx := 0

	for i := k - 1; i < len(nums)-k; i++ {
		for a := range A[i] {
			for b := range B[len(nums)-i-2] {
				if a^b > mx {
					mx = a ^ b
				}
			}
		}
	}
	return mx
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func maxValue2(nums []int, k int) int {
	// dp[i][0]
	var res int
	var dp = make([]map[[2]int]bool, 2*k)
	for i := 0; i < len(dp); i++ {
		dp[i] = make(map[[2]int]bool)
	}
	dp[0][[2]int{nums[0]}] = true

	for i := 1; i < len(nums); i++ {
		for j := 2*k - 1; j >= 0; j-- {
			if j < k {
				if j > 0 {
					for nk := range dp[j-1] {
						dp[j][[2]int{nk[0] | nums[i]}] = true
					}
				} else {
					dp[j][[2]int{nums[i]}] = true
				}
			} else {
				for nk := range dp[j-1] {
					dp[j][[2]int{nk[0], nk[1] | nums[i]}] = true
				}
			}
		}
	}

	for nk := range dp[2*k-1] {
		if res < nk[0]^nk[1] {
			res = nk[0] ^ nk[1]
		}
	}
	return res
}
