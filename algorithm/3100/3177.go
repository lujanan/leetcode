// 3177.
//给你一个整数数组 nums 和一个 非负 整数 k 。如果一个整数序列 seq 满足在范围下标范围 [0, seq.length - 2] 中存在 不超过
//k 个下标 i 满足 seq[i] != seq[i + 1] ，那么我们称这个整数序列为 好 序列。
//
// 请你返回 nums 中 好 子序列 的最长长度
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,1,1,3], k = 2
//
//
// 输出：4
//
// 解释：
//
// 最长好子序列为 [1,2,1,1,3] 。
//
// 示例 2：
//
//
// 输入：nums = [1,2,3,4,5,1], k = 0
//
//
// 输出：2
//
// 解释：
//
// 最长好子序列为 [1,2,3,4,5,1] 。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10³
// 1 <= nums[i] <= 10⁹
// 0 <= k <= min(50, nums.length)
//
//
// Related Topics 数组 哈希表 动态规划 👍 37 👎 0

package algorithm_3100

// leetcode submit region begin(Prohibit modification and deletion)
func maximumLengthV2(nums []int, k int) int {
	lenNums := len(nums)
	dp := make(map[int][]int)
	zd := make([]int, k+1)

	for i := 0; i < lenNums; i++ {
		v := nums[i]
		if _, ok := dp[v]; !ok {
			dp[v] = make([]int, k+1)
		}

		tmp := dp[v]
		for j := 0; j <= k; j++ {
			tmp[j]++
			if j > 0 {
				tmp[j] = max(tmp[j], zd[j-1]+1)
			}
		}

		for j := 0; j <= k; j++ {
			zd[j] = max(zd[j], tmp[j])
		}
	}
	return zd[k]
}

//leetcode submit region end(Prohibit modification and deletion)
