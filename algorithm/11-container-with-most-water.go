package algorithm

// 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var (
		water, tmp  int
		left, right = 0, len(height) - 1
	)
	for left < right {
		if height[left] < height[right] {
			tmp = height[left] * (right - left)
			left++
		} else {
			tmp = height[right] * (right - left)
			right--
		}

		if water < tmp {
			water = tmp
		}
	}
	return water
}
