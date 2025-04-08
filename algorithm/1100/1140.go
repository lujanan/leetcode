// 1140.石子游戏II
//Alice 和 Bob 继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。
//
// Alice 和 Bob 轮流进行，Alice 先开始。最初，M = 1。
//
// 在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。
//
// 游戏一直持续到所有石子都被拿走。
//
// 假设 Alice 和 Bob 都发挥出最佳水平，返回 Alice 可以得到的最大数量的石头。
//
//
//
// 示例 1：
//
//
//输入：piles = [2,7,9,4,4]
//输出：10
//解释：如果一开始 Alice 取了一堆，Bob 取了两堆，然后 Alice 再取两堆。Alice 可以得到 2 + 4 + 4 = 10 堆。
//如果 Alice 一开始拿走了两堆，那么 Bob 可以拿走剩下的三堆。在这种情况下，Alice 得到 2 + 7 = 9 堆。返回 10，因为它更大。
//
//
// 示例 2:
//
//
//输入：piles = [1,2,3,4,5,100]
//输出：104
//
//
//
//
// 提示：
//
//
// 1 <= piles.length <= 100
//
// 1 <= piles[i] <= 10⁴
//
//
// Related Topics 数组 数学 动态规划 博弈 前缀和 👍 315 👎 0

package algorithm_1100

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func stoneGameII(piles []int) int {
	// dp[0][1][1] = p[0] - (dp[1][1][1], dp[1][1][2])
	// dp[0][1][2] = p[0] + p[1] - (dp[2][2][1], dp[2][2][2], dp[2][2][3], dp[2][2][4])

	// dp[i][m][x] = sum(p[i]...p[i+x-1]) - max(dp[i+x][max(m,x)][1 ~ 2*max(m,x)])

	var ll = len(piles)
	var dp = make([][]int, ll)
	dp[0] = make([]int, ll)
	for i := 1; i < len(piles); i++ {
		piles[i] += piles[i-1]
		dp[i] = make([]int, ll)
	}
	dp[ll-1][0] = piles[0]
	if ll > 1 {
		dp[ll-1][0] = piles[ll-1] - piles[ll-2]
	}

	for i := ll - 2; i >= 0; i-- {
		var pre = 0
		if i > 0 {
			pre = piles[i] - piles[i-1]
		}

		for m := 0; m < ll/2; m++ {
			for j := 0; j < 2*(m+1) && i+j+1 < ll; j++ {
				dp[i][m] = max(dp[i][m], piles[i+j]-pre-dp[i+j+1][max(m, j+1)])
			}
		}
	}

	
	return 0
}

func stoneGameIIV1(piles []int) int {
	// dp[0][1][1] = p[0] - (dp[1][1][1], dp[1][1][2])
	// dp[0][1][2] = p[0] + p[1] - (dp[2][2][1], dp[2][2][2], dp[2][2][3], dp[2][2][4])

	// dp[i][m][x] = sum(p[i]...p[i+x-1]) - max(dp[i+x][max(m,x)][1 ~ 2*max(m,x)])

	// 记忆化搜索
	var sMap = make([]map[int]int, len(piles))
	sMap[0] = make(map[int]int)
	for i := 1; i < len(piles); i++ {
		piles[i] += piles[i-1]
		sMap[i] = make(map[int]int)
	}

	var stoneFn func(idx, m int) int
	stoneFn = func(idx, m int) int {
		if idx >= len(piles) {
			return 0
		}
		if idx == len(piles)-1 {
			if idx > 0 {
				return piles[idx] - piles[idx-1]
			}
			return piles[idx]
		}
		if _, ok := sMap[idx][m]; ok {
			return sMap[idx][m]
		}

		var num = math.MinInt64
		var pre int
		if idx > 0 {
			pre = piles[idx-1]
		}
		for i := 0; i+idx < len(piles) && i < 2*m; i++ {
			next := piles[i+idx] - pre - stoneFn(i+idx+1, max(m, i+1))
			num = max(num, next)
		}
		sMap[idx][m] = num
		return num
	}

	winNum := stoneFn(0, 1)
	return winNum + (piles[len(piles)-1]-winNum)/2
}

//leetcode submit region end(Prohibit modification and deletion)
