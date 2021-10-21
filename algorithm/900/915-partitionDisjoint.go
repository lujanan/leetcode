//给定一个数组 A，将其划分为两个连续子数组 left 和 right， 使得：
//
//
// left 中的每个元素都小于或等于 right 中的每个元素。
// left 和 right 都是非空的。
// left 的长度要尽可能小。
//
//
// 在完成这样的分组后返回 left 的长度。可以保证存在这样的划分方法。
//
//
//
// 示例 1：
//
//
//输入：[5,0,3,8,6]
//输出：3
//解释：left = [5,0,3]，right = [8,6]
//
//
// 示例 2：
//
//
//输入：[1,1,1,0,6,12]
//输出：4
//解释：left = [1,1,1,0]，right = [6,12]
//
//
//
//
// 提示：
//
//
// 2 <= A.length <= 30000
// 0 <= A[i] <= 10^6
// 可以保证至少有一种方法能够按题目所描述的那样对 A 进行划分。
//
//
//
// Related Topics 数组 👍 86 👎 0

package algorithm_900

import "sort"

func partitionDisjoint(nums []int) (res int) {
	var left = make([]int, len(nums)-1)
	left[0] = nums[0]
	for i := 1; i < len(nums)-1; i++ {
		left[i] = nums[i]
		if left[i-1] > left[i] {
			left[i] = left[i-1]
		}
	}
	var right = make([]int, len(nums)-1)
	right[len(nums)-2] = nums[len(nums)-1]
	for i := len(nums) - 3; i >= 0; i-- {
		right[i] = nums[i+1]
		if right[i] > right[i+1] {
			right[i] = right[i+1]
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if left[i] <= right[i] {
			return i + 1
		}
	}
	return
}

// 排序再搜索
func partitionDisjoint0(nums []int) (res int) {
	var idx = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		idx[i] = i
	}
	sort.SliceStable(idx, func(i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})

	var max = 0
	for i := 0; i < len(idx); i++ {
		if idx[i] > max {
			max = idx[i]
		}

		if i >= max {
			return i + 1
		}
	}
	return
}
