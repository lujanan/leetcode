// 1760.袋子里最少数目的球
//给你一个整数数组 nums ，其中 nums[i] 表示第 i 个袋子里球的数目。同时给你一个整数 maxOperations 。
//
// 你可以进行如下操作至多 maxOperations 次：
//
//
// 选择任意一个袋子，并将袋子里的球分到 2 个新的袋子中，每个袋子里都有 正整数 个球。
//
//
//
// 比方说，一个袋子里有 5 个球，你可以把它们分到两个新袋子里，分别有 1 个和 4 个球，或者分别有 2 个和 3 个球。
//
//
//
//
// 你的开销是单个袋子里球数目的 最大值 ，你想要 最小化 开销。
//
// 请你返回进行上述操作后的最小开销。
//
//
//
// 示例 1：
//
//
//输入：nums = [9], maxOperations = 2
//输出：3
//解释：
//- 将装有 9 个球的袋子分成装有 6 个和 3 个球的袋子。[9] -> [6,3] 。
//- 将装有 6 个球的袋子分成装有 3 个和 3 个球的袋子。[6,3] -> [3,3,3] 。
//装有最多球的袋子里装有 3 个球，所以开销为 3 并返回 3 。
//
//
// 示例 2：
//
//
//输入：nums = [2,4,8,2], maxOperations = 4
//输出：2
//解释：
//- 将装有 8 个球的袋子分成装有 4 个和 4 个球的袋子。[2,4,8,2] -> [2,4,4,4,2] 。
//- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,4,4,4,2] -> [2,2,2,4,4,2] 。
//- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,4,4,2] -> [2,2,2,2,2,4,2] 。
//- 将装有 4 个球的袋子分成装有 2 个和 2 个球的袋子。[2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2] 。
//装有最多球的袋子里装有 2 个球，所以开销为 2 并返回 2 。
//
//
// 示例 3：
//
//
//输入：nums = [7,17], maxOperations = 2
//输出：7
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= maxOperations, nums[i] <= 10⁹
//
//
// Related Topics 数组 二分查找 👍 283 👎 0

package algorithm_1700

// leetcode submit region begin(Prohibit modification and deletion)
func minimumSize(nums []int, maxOperations int) int {
	// 题目转化为
	// 给定一个 m，使得每个袋子里最多有 m 个球，能否在 maxOperations 次操作内完成
	// 袋子中最大球数为 max ，有 k 个袋子，每个袋子至多有 m 个球
	// 有 k * m >= max，则 k = max / m，分成 k 个袋子则需要 k - 1 次操作
	// 当 max % m != 0 时，需向上取整(即剩余 m < max < 2*m 个球时，也需求进行一次操作)
	// 则分配 max 个球的操作次数 opt = (max - 1) / m，(类似 math.Floor(2.5 + 0.5) 实现四舍五入的实现)

	// 袋子中最大球数
	max := 0
	for _, x := range nums {
		if x > max {
			max = x
		}
	}

	// m 的取值范围为 (0, max) ，使用二分查询得到符合条件的最小的值
	var l, r = 0, max
	var res = max
	for l < r {
		if (l+r)>>1 == 0 {
			l = 1
			continue
		}

		var m = (l + r) >> 1
		var opt = 0
		for i := 0; i < len(nums); i++ {
			opt += (nums[i] - 1) / m
		}

		if opt > maxOperations {
			l = m + 1
		} else {
			res = m
			r = m
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
