//给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
//输出：6
//解释：[1,1,1,0,0,1,1,1,1,1,1]
//粗体数字从 0 翻转到 1，最长的子数组长度为 6。
//
// 示例 2：
//
//
//输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
//输出：10
//解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
//粗体数字从 0 翻转到 1，最长的子数组长度为 10。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// nums[i] 不是 0 就是 1
// 0 <= k <= nums.length
//
// Related Topics 数组 二分查找 前缀和 滑动窗口 👍 391 👎 0

package algorithm_1000

import "sort"

// 假设 [i,j] 为最大长度，则 [i,j] 区间内 0 的数量 <= k
// 前缀和 数组 p，计算 0 的数量，则符合题目的解为 (p[j] - p[i] <= k) <=> (p[j] - k <= p[i])

// 前缀和 + 滑动窗口
func longestOnes(nums []int, k int) int {
	var lsum, rsum, res, l int
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for r, v := range nums {
		rsum += 1 - v
		for lsum < rsum-k {
			lsum += 1 - nums[l]
			l++
		}
		res = max(res, r-l+1)
	}

	return res
}

// 前缀和 + 二分查找
func longestOnes1(nums []int, k int) int {
	var p = make([]int, len(nums)+1)
	for i := range nums {
		p[i+1] = p[i] + 1 - nums[i]
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var res int
	for r, v := range p {
		l := sort.SearchInts(p, v-k)
		res = max(res, r-l)
	}
	return res
}
