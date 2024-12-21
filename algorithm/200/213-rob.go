//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的
//房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,3,2]
//输出：3
//解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3,1]
//输出：4
//解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 3：
//
//
//输入：nums = [1,2,3]
//输出：3
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000
//
// Related Topics 数组 动态规划 👍 933 👎 0

package algorithm_200

import "math"

func rob(nums []int) int {
	var dp = [4]int{0, nums[0], 0, 0}
	for i := 1; i < len(nums); i++ {
		if i < len(nums)-1 {
			dp[1], dp[0] = int(math.Max(float64(dp[0]+nums[i]), float64(dp[1]))), dp[1]
		}
		dp[3], dp[2] = int(math.Max(float64(dp[2]+nums[i]), float64(dp[3]))), dp[3]
	}
	return int(math.Max(math.Max(float64(dp[0]), float64(dp[1])), math.Max(float64(dp[2]), float64(dp[3]))))
}

func robV2(nums []int) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n = len(nums)
	var dp = []int{0, nums[0], 0, 0}
	for i := 1; i < n; i++ {
		if i+1 < n {
			dp[0], dp[1] = dp[1], max(dp[0]+nums[i], dp[1])
		}
		dp[2], dp[3] = dp[3], max(dp[2]+nums[i], dp[3])
	}
	return max(dp[1], dp[3])
}
