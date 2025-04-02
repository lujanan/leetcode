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
	ll := len(nums)
	dp := make(map[int][]int) // 存储已数字nums[i]结尾的各个维度[0,k]的最长子序列长度
	zd := make([]int, k+2) // 存储nums[i]之前的各个维度[0,k]的最长子序列长度

	for i := 0; i < ll; i++ {
		var num = nums[i]
		if _, ok := dp[num]; !ok {
			dp[num] = make([]int, k+2)
		}

		// nums[i] 结尾的子序列长度
		// dp[i][l] = dp[i][l] + 1, nums[i] == nums[j]
		// dp[i][l] = zd[l-1] + 1, nums[i] != nums[j]
		// 状态转移时，zd[l-1] <= dp[i][l]，因为在l相等的情况下，nums[i] == nums[j]时的长度肯定大于等于nums[i] != nums[j]的子序列长度
		for j := 1; j < k+2; j++ {
			dp[num][j] = max(dp[num][j]+1, zd[j-1]+1) 
		}

		for j := 1; j < k+2; j++ {
			zd[j] = max(zd[j], dp[num][j]) // 维护zd的状态为各维度下的最大值
		}
	}
	return zd[k+1]
}

func maximumLengthV3(nums []int, k int) int {
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
