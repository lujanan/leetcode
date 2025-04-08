// 416.分割等和子集
//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 2294 👎 0

package algorithm_400

// leetcode submit region begin(Prohibit modification and deletion)
func canPartition(nums []int) bool {
	var total, max int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		if max < nums[i] {
			max = nums[i]
		}
	}
	var half = total >> 1
	if total&1 != 0 || max > half {
		return false
	}
	if max == half {
		return true
	}

	// 转换成(0,1)背包问题，total/2为背包容量，求是否存在1~n-1个数字能恰好填满背包容量
	// dp[j] = dp[j-nums[i]] + nums[i]
	var dp = make([]int, half+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == half {
			return true
		}

		for j := len(dp) - 1; j >= nums[i]; j-- {
			if dp[j] != 1 && dp[j-nums[i]] == 1 {
				dp[j] = 1
			}
		}
	}

	return dp[len(dp)-1] > 0
}

//leetcode submit region end(Prohibit modification and deletion)
