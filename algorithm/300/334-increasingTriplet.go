//给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
//
// 如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回
//true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,4,5]
//输出：true
//解释：任何 i < j < k 的三元组都满足题意
//
//
// 示例 2：
//
//
//输入：nums = [5,4,3,2,1]
//输出：false
//解释：不存在满足题意的三元组
//
// 示例 3：
//
//
//输入：nums = [2,1,5,0,4,6]
//输出：true
//解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？
// Related Topics 贪心 数组 👍 464 👎 0

package algorithm_300

import "math"

// 贪心
// 时间 O(n)  空间 O(1)
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	var first, second = nums[0], int(math.MaxInt64)
	for i := 1; i < len(nums); i++ {
		if nums[i] > second {
			return true
		}
		if nums[i] > first {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}
	return false
}

// 双指针遍历
// 时间 O(n)  空间 O(n)
func increasingTriplet1(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	var left, right = make([]int, len(nums)), make([]int, len(nums))
	left[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if left[i-1] > nums[i] {
			left[i] = nums[i]
		} else {
			left[i] = left[i-1]
		}
	}

	right[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if right[i+1] < nums[i] {
			right[i] = nums[i]
		} else {
			right[i] = right[i+1]
		}
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > left[i] && nums[i] < right[i] {
			return true
		}
	}
	return false
}
