// 2537.统计好子数组的数目
//给你一个整数数组 nums 和一个整数 k ，请你返回 nums 中 好 子数组的数目。
//
// 一个子数组 arr 如果有 至少 k 对下标 (i, j) 满足 i < j 且 arr[i] == arr[j] ，那么称它是一个 好 子数组。
//
// 子数组 是原数组中一段连续 非空 的元素序列。
//
//
//
// 示例 1：
//
// 输入：nums = [1,1,1,1,1], k = 10
//输出：1
//解释：唯一的好子数组是这个数组本身。
//
//
// 示例 2：
//
// 输入：nums = [3,1,4,3,2,2,4], k = 2
//输出：4
//解释：总共有 4 个不同的好子数组：
//- [3,1,4,3,2,2] 有 2 对。
//- [3,1,4,3,2,2,4] 有 3 对。
//- [1,4,3,2,2,4] 有 2 对。
//- [4,3,2,2,4] 有 2 对。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i], k <= 10⁹
//
//
// Related Topics 数组 哈希表 滑动窗口 👍 95 👎 0

package algorithm_2500

// leetcode submit region begin(Prohibit modification and deletion)
func countGood(nums []int, k int) int64 {
	var numMap = make(map[int]int)
	var goodNum int
	var ll = len(nums)
	var res int64

	var updateNum = func(idx int, add bool) {
		if add {
			numMap[nums[idx]]++
			var num = numMap[nums[idx]]
			if num >= 2 {
				goodNum += num - 1
			}

		} else {
			numMap[nums[idx]]--
			var num = numMap[nums[idx]]
			if num >= 1 {
				goodNum -= num
			} else {
				delete(numMap, nums[idx])
			}
		}
	}

	for i, j := 0, -1; j < ll; {
		for goodNum < k && j < ll-1 {
			j++
			updateNum(j, true)
		}
		if j == ll-1 && goodNum < k {
			break
		}

		for goodNum >= k && i < j {
			res += int64(ll - j)
			updateNum(i, false)
			i++
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
