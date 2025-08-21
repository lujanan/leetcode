// 1277.统计全为1的正方形子矩阵

package algorithm_1200

func countSquares(matrix [][]int) int {
	var m, n = len(matrix), len(matrix[0])
	var minL = min(m, n) + 1
	var ans int
	var dp [2][][]int
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, minL)
		}
	}
	for i := 0; i < n; i++ {
		dp[0][i][1] = matrix[0][i]
		ans += dp[0][i][1]
	}

	for i := 1; i < m; i++ {
		dp[1][0][1] = matrix[i][0]
		ans += dp[1][0][1]

		for j := 1; j < n; j++ {
			dp[1][j][1] = matrix[i][j]
			ans += dp[1][j][1]

			for k := 2; k < minL; k++ {
				if matrix[i][j] == 0 || i-k+1 < 0 || j-k+1 < 0 {
					dp[1][j][k] = 0

				} else {
					dp[1][j][k] = dp[0][j-1][k-1] & dp[0][j][k-1] & dp[1][j-1][k-1]
					ans += dp[1][j][k]
				}
			}
		}

		dp[0], dp[1] = dp[1], dp[0]
	}

	return ans
}
