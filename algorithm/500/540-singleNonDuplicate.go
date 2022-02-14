//给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
//
// 请你找出并返回只出现一次的那个数。
//
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,1,2,3,3,4,4,8,8]
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums =  [3,3,7,7,10,11,11]
//输出: 10
//
//
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁵
//
// Related Topics 数组 二分查找 👍 406 👎 0

package algorithm_500

// 优化
func singleNonDuplicate(nums []int) int {
	var l, r = 0, len(nums) - 1
	for l < r {
		var m = l + (r-l)>>1
		// 如果 target 不在左边，
		// m 为偶数时，num[m] == num[m+1]，位运算 m^1 = m+1
		// m 为基数时，num[m] == num[m-1]，位运算 m^1 = m-1
		if nums[m] == nums[m^1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}

func singleNonDuplicate1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var l, max = 0, len(nums)
	var r = max - 1
	for l <= r {
		var m = l + (r-l)>>1
		if (m-1 < 0 || nums[m-1] != nums[m]) && (m+1 >= max || nums[m+1] != nums[m]) {
			return nums[m]
		}

		if (m&1 == 0 && m-1 >= 0 && nums[m-1] == nums[m]) ||
			(m&1 == 1 && m+1 < max && nums[m+1] == nums[m]) {
			r = m - 1

		} else {
			l = m + 1
		}
	}

	return 0
}
