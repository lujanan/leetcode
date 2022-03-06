package algorithm_0

// 41. 缺失的第一个正数
// https://leetcode-cn.com/problems/first-missing-positive/

// func firstMissingPositive(nums []int) int {
func firstMissingPositive(orign []int) int {
	nums := append([]int{}, orign...)
	var n = len(nums)
	var m = nums[0]
	for i := 1; i < n; i++ {
		if m < nums[i] {
			m = nums[i]
		} else if m < -nums[i] {
			m = -nums[i]
		}
	}
	if m < 1 {
		return 1
	}
	nums = append(nums, -1)
	m++

	for i := range nums {
		idx := nums[i] % m
		if idx >= 0 && idx <= n {
			if nums[idx] < 0 {
				nums[idx] -= m
			} else {
				nums[idx] += m
			}
		}
	}

	var res = 1
	for ; res <= n; res++ {
		if (nums[res] >= 0 && nums[res] < m) || (nums[res] < 0 && nums[res] >= -m) {
			return res
		}
	}
	return res
}
