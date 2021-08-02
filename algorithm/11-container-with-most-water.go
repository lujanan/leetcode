package algorithm

// 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	var water, tmp int
	var left, right = 0, len(height) - 1

	for left < right {
		var width = right - left
		if height[left] < height[right] {
			tmp = height[left] * width
			left++
		} else {
			tmp = height[right] * width
			right--
		}

		if water < tmp {
			water = tmp
		}
	}
	return water
}
