//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
//被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2：
//
//
//输入：[2,7,9,3,1]
//输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400
//
// Related Topics 数组 动态规划 👍 1849 👎 0

package algorithm_100

import "math"

func rob(nums []int) int {
	var dp = []int{0, nums[0]}
	for i := 1; i < len(nums); i++ {
		dp[1], dp[0] = dp[0]+nums[i], int(math.Max(float64(dp[0]), float64(dp[1])))
	}
	return int(math.Max(float64(dp[0]), float64(dp[1])))
}

func robV3(nums []int) int {
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1])
	// dp[i][1] = dp[i-1][0] + nums[i]
	var res = [2][2]int{{0, nums[0]}}
	for i := 1; i < len(nums); i++ {
		res[1][0] = int(math.Max(float64(res[0][0]), float64(res[0][1])))
		res[1][1] = nums[i] + res[0][0]
		res[0] = res[1]
	}
	return int(math.Max(float64(res[0][0]), float64(res[0][1])))
}

func robV2(nums []int) int {
	// dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dp = [2]int{0, nums[0]}
	for i := 1; i < len(nums); i++ {
		dp[1], dp[0] = max(dp[1], dp[0]+nums[i]), dp[1]
	}
	return max(dp[1], dp[0])
}
