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

	for l := 1; l <= n; l++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if l == 1 {
					if mat[i][j] == 0 {
						dp[0][i][j] = 0
						} else {
						dp[0][i][j] = 0
						
					}
				}

	
	
				if mat[i][j] == 0 {
					dp1[i][j] = 0
				} else {
					dp1[i][j] = mat[i][j] + dp1[i][j-1]
				}
				ans += dp1[i][j]
			}
		}
		
	}


	return ans
}
