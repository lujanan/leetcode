package algorithm_700

// 至少是其他数字两倍的最大数
// https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/
func dominantIndex(nums []int) int {
	var max, second = -1, -1
	for i, v := range nums {
		if max == -1 {
			max = i
		} else if v > nums[max] {
			second = max
			max = i
		} else if second == -1 || v > nums[second] {
			second = i
		}
	}
	if max > -1 && second > -1 && nums[max] < nums[second]*2 {
		max = -1
	}
	return max
}
