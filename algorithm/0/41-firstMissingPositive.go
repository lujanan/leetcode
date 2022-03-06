package algorithm_0

// 41. 缺失的第一个正数
// https://leetcode-cn.com/problems/first-missing-positive/

// func firstMissingPositive(nums []int) int {
func firstMissingPositive(orign []int) int {
	nums := append([]int{}, orign...)
	// 如果数组中包含 1~n, return n+1
	// 如果数组中缺少 1~n 中的某个数, return 该数
	var n = len(nums)
	// 所有 <=0 转化为 n+1
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 所有 <=n 的转为负数
	for i := range nums {
		if nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] > 0 {
			nums[nums[i]-1] = -nums[nums[i]-1]
		} else if nums[i] < 0 && -nums[i] <= n && nums[-nums[i]-1] > 0 {
			nums[-nums[i]-1] = -nums[-nums[i]-1]
		}
	}

	// 第一个为正数的，下标 +1 即为缺失的数
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func firstMissingPositive1(nums []int) int {
	// func firstMissingPositive(orign []int) int {
	// nums := append([]int{}, orign...)
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
