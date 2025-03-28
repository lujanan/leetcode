// 2218.从栈中取出k个硬币的最大面值和
//一张桌子上总共有 n 个硬币 栈 。每个栈有 正整数 个带面值的硬币。
//
// 每一次操作中，你可以从任意一个栈的 顶部 取出 1 个硬币，从栈中移除它，并放入你的钱包里。
//
// 给你一个列表 piles ，其中 piles[i] 是一个整数数组，分别表示第 i 个栈里 从顶到底 的硬币面值。同时给你一个正整数 k ，请你返回在 恰
//好 进行 k 次操作的前提下，你钱包里硬币面值之和 最大为多少 。
//
//
//
// 示例 1：
//
//
//
//
//输入：piles = [[1,100,3],[7,8,9]], k = 2
//输出：101
//解释：
//上图展示了几种选择 k 个硬币的不同方法。
//我们可以得到的最大面值为 101 。
//
//
// 示例 2：
//
//
//输入：piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
//输出：706
//解释：
//如果我们所有硬币都从最后一个栈中取，可以得到最大面值和。
//
//
//
//
// 提示：
//
//
// n == piles.length
// 1 <= n <= 1000
// 1 <= piles[i][j] <= 10⁵
// 1 <= k <= sum(piles[i].length) <= 2000
//
//
// Related Topics 数组 动态规划 前缀和 👍 130 👎 0

package algorithm_2200

// leetcode submit region begin(Prohibit modification and deletion)
func maxValueOfCoins(piles [][]int, k int) int {
	// dp[1] = max(piles[i][1])
	// dp[2] = max(piles[i][2]) or [max(piles[i][1]) + max(dp[1])]

	// var dp = make([]int, k+1)
	// for _, p := range piles {
	// 	if dp[1] < p[0] {
	// 		dp[1], dp[0] = p[0], dp[1]
	// 	} else if dp[0] < p[0] {
	// 		dp[0] = p[0]
	// 	}

	// 	for j := 1; j < len(p) && j < k; j++ {
	// 		p[j] += p[j-1]
	// 	}
	// }

	// var dp = make([]int, k+1)
	// for i := 1; i < count; i++ {

	// }
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
