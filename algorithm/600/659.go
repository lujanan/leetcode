// 659.分割数组为连续子序列
//给你一个按 非递减顺序 排列的整数数组 nums 。
//
// 请你判断是否能在将 nums 分割成 一个或多个子序列 的同时满足下述 两个 条件：
//
//
//
//
// 每个子序列都是一个 连续递增序列（即，每个整数 恰好 比前一个整数大 1 ）。
// 所有子序列的长度 至少 为 3 。
//
//
//
//
// 如果可以分割 nums 并满足上述条件，则返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,3,4,5]
//输出：true
//解释：nums 可以分割成以下子序列：
//[1,2,3,3,4,5] --> 1, 2, 3
//[1,2,3,3,4,5] --> 3, 4, 5
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3,3,4,4,5,5]
//输出：true
//解释：nums 可以分割成以下子序列：
//[1,2,3,3,4,4,5,5] --> 1, 2, 3, 4, 5
//[1,2,3,3,4,4,5,5] --> 3, 4, 5
//
//
// 示例 3：
//
//
//输入：nums = [1,2,3,4,4,5]
//输出：false
//解释：无法将 nums 分割成长度至少为 3 的连续递增子序列。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -1000 <= nums[i] <= 1000
// nums 按非递减顺序排列
//
//
// Related Topics 贪心 数组 哈希表 堆（优先队列） 👍 480 👎 0

package algorithm_600

// leetcode submit region begin(Prohibit modification and deletion)
func isPossible(nums []int) bool {
	var numMap = make(map[int][]int)
	for _, n := range nums {
		ns, ok := numMap[n-1]
		if !ok || len(ns) == 0 {
			numMap[n] = append(numMap[n], 1)
		} else {
			var idx = 0
			for i := 1; i < len(ns); i++ {
				if ns[i] < ns[idx] {
					idx = i
				}
			}

			numMap[n] = append(numMap[n], ns[idx]+1)
			if idx != len(ns)-1 {
				ns[idx], ns[len(ns)-1] = ns[len(ns)-1], ns[idx]
			}
			ns = ns[:len(ns)-1]
			if len(ns) > 0 {
				numMap[n-1] = ns
			} else {
				delete(numMap, n-1)
			}
		}
	}

	for _, ns := range numMap {
		for i := 0; i < len(ns); i++ {
			if ns[i] < 3 {
				return false
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
