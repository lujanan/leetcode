package algorithm_200

// todo 空间优化，只保留一个左数组lnums,计算左前缀乘，
// 然后再计算右前缀乘(结果保留在nums中)，
// 返回数组不算空间，所以空间复杂度变为O(1)
func productExceptSelf(nums []int) []int {
	var n = len(nums)
	var lnums = make([]int, n)
	var rnums = make([]int, n)
	lnums[0], rnums[n-1] = 1, 1
	for i := 1; i < n; i++ {
		lnums[i] = nums[i-1] * lnums[i-1]
		rnums[n-i-1] = nums[n-i] * rnums[n-i]
	}

	for i := 0; i < n; i++ {
		nums[i] = lnums[i] * rnums[i]
	}
	return nums
}
