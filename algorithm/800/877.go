// 877.石子游戏
//Alice 和 Bob 用几堆石子在做游戏。一共有偶数堆石子，排成一行；每堆都有 正 整数颗石子，数目为 piles[i] 。
//
// 游戏以谁手中的石子最多来决出胜负。石子的 总数 是 奇数 ，所以没有平局。
//
// Alice 和 Bob 轮流进行，Alice 先开始 。 每回合，玩家从行的 开始 或 结束 处取走整堆石头。 这种情况一直持续到没有更多的石子堆为止，此
//时手中 石子最多 的玩家 获胜 。
//
// 假设 Alice 和 Bob 都发挥出最佳水平，当 Alice 赢得比赛时返回 true ，当 Bob 赢得比赛时返回 false 。
//
//
//
// 示例 1：
//
//
//输入：piles = [5,3,4,5]
//输出：true
//解释：
//Alice 先开始，只能拿前 5 颗或后 5 颗石子 。
//假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
//如果 Bob 拿走前 3 颗，那么剩下的是 [4,5]，Alice 拿走后 5 颗赢得 10 分。
//如果 Bob 拿走后 5 颗，那么剩下的是 [3,4]，Alice 拿走后 4 颗赢得 9 分。
//这表明，取前 5 颗石子对 Alice 来说是一个胜利的举动，所以返回 true 。
//
//
// 示例 2：
//
//
//输入：piles = [3,7,2,3]
//输出：true
//
//
//
//
// 提示：
//
//
// 2 <= piles.length <= 500
// piles.length 是 偶数
// 1 <= piles[i] <= 500
// sum(piles[i]) 是 奇数
//
//
// Related Topics 数组 数学 动态规划 博弈 👍 550 👎 0

package algorithm_800

// leetcode submit region begin(Prohibit modification and deletion)
func stoneGame(piles []int) bool {
	// 数学解法，先手必胜...，
	// 长度是偶数，可以分别直接计算偶数下标和奇数下标的数的和，先手的的玩家直接取大的那组下标就能赢
	// 直接 return true

	// dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])

	var ll = len(piles)
	var dp = make([]int, ll)
	for i := 0; i < ll; i++ {
		dp[i] = piles[i]
	}

	for i := ll - 2; i >= 0; i-- {
		for j := i + 1; j < ll; j++ {
			dp[j] = max(piles[i]-dp[j], piles[j]-dp[j-1])
		}
	}
	return dp[ll-1] > 0
}

func stoneGameV0(piles []int) bool {
	// dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])

	var ll = len(piles)
	var dp = make([][]int, ll)
	for i := 0; i < ll; i++ {
		dp[i] = make([]int, ll)
		dp[i][i] = piles[i]
	}

	for i := ll - 2; i >= 0; i-- {
		for j := i + 1; j < ll; j++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][ll-1] > 0
}

//leetcode submit region end(Prohibit modification and deletion)
