package algorithm_0

// 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	var minFn = func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	var maxFn = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	var water int
	for i, j := 0, len(height)-1; i < j; {
		water = maxFn(water, (j-i)*minFn(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return water
}

func maxAreaV2(height []int) int {
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
