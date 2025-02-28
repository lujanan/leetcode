// 42.接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5547 👎 0

package algorithm_0

// leetcode submit region begin(Prohibit modification and deletion)

// 双指针，从两侧向中间靠拢，每次移动较矮的位置
func trap(height []int) int {
	var num int
	var left, right = 0, len(height) - 1
	var leftMax, rightMax = 0, 0
	for left < right {
		leftMax, rightMax = max(leftMax, height[left]), max(rightMax, height[right])
		
		if height[left] <= height[right] {
			num += leftMax - height[left]
			left++
		} else {
			num += rightMax - height[right]
			right--
		}
	}
	return num
}

// 单调栈
func trapV2(height []int) int {
	var stack = make([]int, len(height))
	var num, top int
	for i := 1; i < len(height); i++ {
		for top > 0 && height[i] > height[stack[top]] {
			var (
				w = i - stack[top-1] - 1
				h = min(height[i], height[stack[top-1]]) - height[stack[top]]
			)
			num += w * h
			top--
		}

		for top >= 0 && height[stack[top]] <= height[i] {
			top--
		}
		top++
		stack[top] = i
	}
	return num
}

// 动态规划
// 需要预处理每个点左右两边的最高的点，才能拿到当前点能装多少水
func trapV3(height []int) int {
	var hl = len(height)
	var left, right = make([]int, hl), make([]int, hl)
	left[0] = height[0]
	right[hl-1] = height[hl-1]

	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
		right[hl-1-i] = max(right[hl-i], height[hl-1-i])
	}

	var num int
	for i := 0; i < hl; i++ {
		num += min(left[i], right[i]) - height[i]
	}
	return num
}

//leetcode submit region end(Prohibit modification and deletion)
