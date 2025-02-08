// 63.不同路径II
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
//
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//
//
// 示例 1：
//
//
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//
//
// 示例 2：
//
//
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1
//
//
//
//
// 提示：
//
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] 为 0 或 1
//
// Related Topics 数组 动态规划 矩阵 👍 725 👎 0

package algorithm_0

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var dp = make([]int, len(obstacleGrid[0]))
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			// 障碍物不可达，赋0
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}

			if j == 0 {
				if i == 0 {
					dp[j] = 1 // 起点赋值
				}
				// 第1列只能从上面方向下来，值不用变
				continue
			}

			if i == 0 {
				dp[j] = dp[j-1] // 第1行只能从左面方向过来，直接取左边的值
			} else {
				dp[j] += dp[j-1] // 上面与左边的值相加
			}
		}
	}
	return dp[len(dp)-1]
}

func uniquePathsWithObstaclesV2(obstacleGrid [][]int) int {
	var dp = make([]int, len(obstacleGrid[0])+1)
	dp[0] = 0
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 1; j <= len(obstacleGrid[0]); j++ {
			if i == 0 {
				dp[j] = obstacleGrid[i][j-1] ^ 1
				if j > 1 {
					dp[j] &= dp[j-1]
				}
			} else {
				if obstacleGrid[i][j-1] == 1 {
					dp[j] = 0
				} else {
					dp[j] += dp[j-1]
				}
			}
		}

	}
	return dp[len(obstacleGrid[0])]
}
