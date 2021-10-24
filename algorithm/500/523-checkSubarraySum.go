package algorithm_500

// 前缀和
// n^2 会超时
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	var res = map[int]int{0: -1}
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			nums[i] += nums[i-1]
		}
		nums[i] = nums[i] % k
		if idx, ok := res[nums[i]]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			res[nums[i]] = i
		}
	}
	return false
}
