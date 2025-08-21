// 1504.统计全1子矩形

package algorithm_1500

func numSubmat(mat [][]int) int {
	var ans int
	var m, n = len(mat), len(mat[0])
	var dp = make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	// dp[0][i][j] 记录以 mat[i][j] 为右下角 且 矩阵高度=1 的矩阵数量
	for i := 0; i < m; i++ {
		dp[0][i][0], dp[1][i][0] = mat[i][0], mat[i][0]
		ans += dp[1][i][0]
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}

			dp[0][i][j] = 1 + dp[0][i][j-1]
			dp[1][i][j] = dp[0][i][j]
			ans += dp[1][i][j]
		}
	}

	// 枚举以 mat[i][j] 为右下角 且 矩阵高度=h+1 的矩阵数量
	// 右下角固定时，矩阵高度h数量 = min(高度1, 高度h-1) 
	// 所以只保留高度1(dp[0])的矩阵数量, 高度h-1(dp[1])动态保存
	// 防止数据被覆盖，需要从 mat[m-1][n-1]开始枚举
	for h := 1; h < m; h++ {
		for i := m - 1; i >= h; i-- {
			for j := n - 1; j >= 0; j-- {
				if mat[i][j] == 0 {
					continue
				}

				dp[1][i][j] = min(dp[0][i][j], dp[1][i-1][j])
				ans += dp[1][i][j]
			}
		}
	}

	return ans
}
