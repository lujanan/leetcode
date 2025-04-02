// 3176.求出最长好子序列
//给你一个整数数组 nums 和一个 非负 整数 k 。如果一个整数序列 seq 满足在下标范围 [0, seq.length - 2] 中 最多只有 k 个
//下标 i 满足 seq[i] != seq[i + 1] ，那么我们称这个整数序列为 好 序列。
//
// 请你返回 nums 中 好 子序列 的最长长度。
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
// 1 <= nums.length <= 500
// 1 <= nums[i] <= 10⁹
// 0 <= k <= min(nums.length, 25)
//
//
// Related Topics 数组 哈希表 动态规划 👍 46 👎 0

package algorithm_3100

// leetcode submit region begin(Prohibit modification and deletion)
// 时间复杂度 O(k*n^2)
func maximumLength(nums []int, k int) int {
	// nums[i] == nums[j]:   dp[k][i] = max(dp[k][j] + 1, dp[k][i])
	// nums[i] != nums[j]:   dp[k][i] = max(dp[k-1][j] + 1, dp[k][i])

	var num = len(nums)
	var dp = make([][]int, num)

	var res int
	for i := 0; i < num; i++ {
		dp[i] = make([]int, k+2)
		for l := 1; l < k+2; l++ {
			dp[i][l] = 1
			for j := 0; j < i; j++ {
				if nums[i] == nums[j] {
					dp[i][l] = max(dp[j][l]+1, dp[i][l])
				} else {
					dp[i][l] = max(dp[j][l-1]+1, dp[i][l])
				}
			}
		}
		res = max(res, dp[i][k+1])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
